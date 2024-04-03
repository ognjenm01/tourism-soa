FROM golang:alpine as build_container
WORKDIR /app
COPY module_tour/go.mod .
COPY module_tour/go.sum .
RUN go mod download
COPY module_tour/. .
RUN go build -o tours-module

FROM alpine
COPY --from=build_container /app/tours-module /usr/bin
ENTRYPOINT [ "tours-module" ]