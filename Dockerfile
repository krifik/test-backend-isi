
FROM golang:1.20
WORKDIR /application
COPY . .

# COPY go.mod go.sum ./

# download depedency
RUN go mod tidy
# Build the binary files
RUN go build -o binary 
RUN ls
# ENTRYPOINT ["/binary"]
CMD ["ls"]
CMD ["./binary"]