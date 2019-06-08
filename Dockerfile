FROM golang:1.12.5 as builder

LABEL maintainer="jearzamendia@gmail.com"

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/jarzamendia/konger

COPY get.sh .

RUN bash get.sh

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/konger .

######## Start a new stage from scratch #######
FROM alpine:3.9  

WORKDIR /root

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/konger .

ENTRYPOINT ["./konger"]

ARG BUILD_DATE

# Labels.
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.build-date=$BUILD_DATE
LABEL org.label-schema.name="konger"
LABEL org.label-schema.description="Imagem com Alpine 3.9, binario com Golang 1.12.5"
LABEL org.label-schema.url="https://golang.org/"
LABEL org.label-schema.vendor="Jarza"
LABEL org.label-schema.version="1"
LABEL org.label-schema.docker.cmd="docker run -it --e KONG_ADMIN_ADDR:http://localhost/admin -e TLS_SKIP_VERIFY:true -e KONG_API_KEY:xxxxxxxxxxxxxxxxxxxxxx -p 8081:8081 jarzamendia/konger:latest"