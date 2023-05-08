FROM --platform=$BUILDPLATFORM golang:latest AS builder
WORKDIR  /go/src/github.com/GoTalkie/go-talkie-auth/
COPY ./ ./
RUN go mod vendor
ARG TARGETOS TARGETARCH
RUN env GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 go build -buildvcs=false -o authApp ./cmd/auth-app/


FROM alpine:latest
RUN mkdir /app
COPY --from=builder /go/src/github.com/GoTalkie/go-talkie-auth/authApp /app
COPY .env /

CMD ["/app/authApp"]