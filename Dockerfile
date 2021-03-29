FROM golang:alpine AS Builder
RUN mkdir /src
COPY /src /src
WORKDIR /src
RUN go mod download
RUN go build -o app .

FROM alpine:latest AS Final
COPY --from=Builder /src/app /app/app
WORKDIR /app
CMD ["/app/app"]