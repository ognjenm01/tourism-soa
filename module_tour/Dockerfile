FROM golang:alpine as build_container
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o tours-module

FROM alpine
COPY --from=build_container /app/tours-module /usr/bin
ENTRYPOINT [ "tours-module" ]