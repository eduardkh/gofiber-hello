# based on https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j#dockerfile-for-the-fiber-app
FROM golang:alpine AS Builder
RUN mkdir /src
COPY /src /src
WORKDIR /src
RUN go mod download
RUN go build -o app .
# for raspberry pi version
# RUN env GOOS=linux GOARCH=arm GOARM=7 go build -o app .

FROM alpine:latest AS Final
COPY  /src/public /app/public
COPY --from=Builder /src/app /app/app
WORKDIR /app
CMD ["/app/app"]