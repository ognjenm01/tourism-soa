FROM golang:alpine as build_container
WORKDIR /app
COPY api_gateway/go.mod .
COPY api_gateway/go.sum .
RUN go mod download
COPY api_gateway/. .
RUN go build -o api-gateway

FROM alpine
COPY --from=build_container /app/api-gateway /usr/bin
ENTRYPOINT [ "api-gateway" ]
