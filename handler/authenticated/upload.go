package authenticated

import (
	"encoding/base64"
	"errors"
	"io"
	"net/http"

	"github.com/elotus_hackathon/model"
	"github.com/elotus_hackathon/pkg/httpserv"
)

const maximumBytes = 8000000

// UploadFile is handler func for upload file, just allow images
func (h Handler) UploadFile() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		uploadFile, header, err := r.FormFile("file")
		if err != nil {
			return err
		}
		defer uploadFile.Close()

		// valid and the content type of the uploaded file is an image
		contentType := header.Header.Get("Content-Type")
		if contentType != "image/png" && contentType != "image/jpeg" {
			return errors.New("unsupported content type")
		}

		// images larger than 8 megabytes should also be rejected
		if header.Size > maximumBytes {
			return errors.New("image too big")
		}

		data, err := io.ReadAll(uploadFile)
		if err != nil {
			return err
		}
		imgBase64Str := base64.StdEncoding.EncodeToString(data)

		_, err = h.fileCtrl.UploadFile(ctx, model.File{
			UserID: 1,
			Name:   header.Filename,
			Type:   contentType,
			Size:   header.Size,
			Data:   imgBase64Str,
		})
		if err != nil {
			return err
		}

		httpserv.RespondJSON(ctx, w, httpserv.Success{Message: "success"})
		return nil
	})
}
