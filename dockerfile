# Use an official Golang runtime as the base image
FROM golang:1.22.1-alpine3.19 AS BuildStage

# Set the working directory inside the container
WORKDIR /go/src/csw

# Copy the local package files to the container's workspace
COPY . .

# Build the Go app
RUN go build -o main .

# Expose any ports your application listens on
EXPOSE 6996

# Deploy Stage
FROM alpine:latest

COPY env.json /

# Copy the built executable from BuildStage to the root directory
COPY --from=BuildStage /go/src/csw /
EXPOSE 6996
# Set permissions and execute the application

ENTRYPOINT ["/main"]