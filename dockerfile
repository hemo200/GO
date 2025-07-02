# Dockerfile
FROM golang:1.21-alpine AS builder

#defind the folder
WORKDIR /app 

# copy files into a folder
COPY go.mod go.sum ./

# keep a copy of pakcages in cache
RUN go mod download
# copy all files
COPY . . 
# Build the go application
RUN go build -o main ./web

# build the image
FROM  alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 5000
CMD [ "./main" ]
