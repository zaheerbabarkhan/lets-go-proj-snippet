# syntax=docker/dockerfile:1
FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./cmd/web/*.go  ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /project-binary


FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=builder /project-binary /project-binary

COPY ./ui/ ./ui

EXPOSE 4000

CMD ["/project-binary"]



