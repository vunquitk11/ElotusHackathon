package authenticated

import (
	"encoding/base64"
	model2 "github.com/petme/api/internal/model"
	httpserv2 "github.com/petme/api/pkg/httpserv"
	"io"
	"net/http"
)

const maximumBytes = 8000000

// UploadFile is handler func for upload file, just allow images
func (h Handler) UploadFile() http.HandlerFunc {
	return httpserv2.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		// pull username from context
		username := ctx.Value("userName").(string)
		if username == "" {
			return &httpserv2.Error{
				Status: http.StatusUnauthorized,
				Code:   "user_not_found",
				Desc:   model2.ErrUserNotFound.Error(),
			}
		}

		uploadFile, header, err := r.FormFile("file")
		if err != nil {
			return err
		}
		defer uploadFile.Close()

		// valid and the content type of the uploaded file is an image
		contentType := header.Header.Get("Content-Type")
		if contentType != "image/png" && contentType != "image/jpeg" {
			return &httpserv2.Error{
				Status: http.StatusBadRequest,
				Code:   "unsupported_content_type",
				Desc:   "unsupported content type",
			}
		}

		// images larger than 8 megabytes should also be rejected
		if header.Size > maximumBytes {
			return &httpserv2.Error{
				Status: http.StatusBadRequest,
				Code:   "image_too_big",
				Desc:   "image too big",
			}
		}

		data, err := io.ReadAll(uploadFile)
		if err != nil {
			return err
		}
		imgBase64Str := base64.StdEncoding.EncodeToString(data)

		_, err = h.fileCtrl.UploadFile(ctx, username, model2.File{
			UserID: 1,
			Name:   header.Filename,
			Type:   contentType,
			Size:   header.Size,
			Data:   imgBase64Str,
		})
		if err != nil {
			return err
		}

		httpserv2.RespondJSON(ctx, w, httpserv2.Success{Message: "success"})
		return nil
	})
}
