FROM mcr.microsoft.com/devcontainers/go:1-1.20-bullseye

RUN go install github.com/99designs/gqlgen@v0.17.36
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@latest

ARG GO_PATH=go env GOPATH
ENV PATH=$GO_PATH/bin:$PATH
