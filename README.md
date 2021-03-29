# Gofiber Hello World (Dockerized)
## Write the Application
>initiate a go project

`go mod init github.com/eduardkh/gofiber-hello`
> get the framework

`go get github.com/gofiber/fiber/v2`
> test the app locally
> 
`go run main.go`
## Dockerization
> Build and Upload the Application

`docker login`

`docker build -t eduardkh/gofiber:v0.1 .`

`docker build -t eduardkh/gofiber:latest .`

>for raspberry pi version (better build on raspberry pi)

**`docker build --platform linux/arm/v7 -t eduardkh/gofiber-arm:v0.1 .`**

`docker push eduardkh/gofiber:v0.1`

`docker push eduardkh/gofiber:latest`
> Run the Application

`docker pull eduardkh/gofiber`

`docker run -p 80:80 eduardkh/gofiber:latest`