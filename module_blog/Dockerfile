FROM golang:alpine as build_container
WORKDIR /app
COPY module_blog/go.mod .
COPY module_blog/go.sum .
RUN go mod download
COPY module_blog/. .
RUN go build -o blogs-module

FROM alpine
COPY --from=build_container /app/blogs-module /usr/bin
ENTRYPOINT [ "blogs-module" ]