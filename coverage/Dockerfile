FROM golang:1.11-alpine AS build_base
 
# Install some dependencies needed to build the project
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/github.com/masmovil
 
# Force the go compiler to use modules
ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .
RUN go mod download


FROM build_base AS app_build

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o app


FROM alpine
RUN apk add ca-certificates
# Finally we copy the statically compiled Go binary.
COPY --from=app_build /go/src/github.com/masmovil/app /app
ENTRYPOINT ["/app"]