#FROM golang:1.22.2 AS build-stage
FROM golang:alpine as build-stage

# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git curl && rm -rf /var/lib/apt/lists/*

#Set working directory
WORKDIR /app

#Copy gomod and gosum files and download dependencies
COPY go.mod go.sum ./
#COPY go.mod ./
RUN go mod download

#Copy files from all directories
COPY . .

#Compile the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./Domain-Orchestrator-Connectors

#ToDo add certificates
# Start a new stage from scratch
#FROM alpine:3.17
#RUN apk --no-cache add ca-certificates
#RUN addgroup -S HorseContext && adduser -S HorseContext -G HorseContext -u 1001
#USER HorseContext

#Expose port to serve API REST, ToDO should be variable
EXPOSE 8080

#Run binary file
CMD ["./Domain-Orchestrator-Connectors"]