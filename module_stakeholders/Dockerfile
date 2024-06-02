FROM golang:alpine as build_container
WORKDIR /app
COPY module_stakeholders/go.mod .
COPY module_stakeholders/go.sum .
RUN go mod download
COPY module_stakeholders/. .
RUN go build -o stakeholders-module

FROM alpine
COPY --from=build_container /app/stakeholders-module /usr/bin
ENTRYPOINT [ "stakeholders-module" ]
