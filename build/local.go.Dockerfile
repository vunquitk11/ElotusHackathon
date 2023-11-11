FROM golang:1.21

ENV GOPRIVATE=code.in.spdigital.sg

RUN apt-get update

COPY build/.netrc /root/.netrc
RUN chmod 600 /root/.netrc

RUN GO111MODULE=on go install golang.org/x/tools/cmd/goimports@v0.2.0

RUN GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4@v4.14.1 && \
    GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.14.1 \
