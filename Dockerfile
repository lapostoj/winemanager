# build executable binary
FROM golang:1.14.3-alpine3.11 AS build_image

WORKDIR $GOPATH/src/github.com/lapostoj/winemanager
COPY . .

RUN go mod download
RUN go mod verify

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/winemanager main.go


# copy binary for runtime
FROM alpine:3.11

COPY --from=build_image /go/bin/winemanager /bin/winemanager

COPY frontend/ /frontend/

ENTRYPOINT ["bin/winemanager"]
