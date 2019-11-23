FROM golang:1.13.4-alpine3.10
RUN apk update && apk add --no-cache git
ENV GO111MODULE=on
RUN mkdir app
WORKDIR app
COPY go.mod go.sum ./
COPY . ./
COPY resources ./app/
RUN ls -l ./app/
#RUN go mod init

RUN go mod download
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./pkg/main/server.go

# Expose port 8080 to the outside world
EXPOSE 4040

# Command to run the executable
CMD ["./server"]
