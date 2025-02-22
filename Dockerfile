# to Download the latest stable version of golang-- alpine is a Lightweight (khfifa) version of linux with a small size  5 MB
FROM golang:alpine

WORKDIR /app

# Install the Bash shell in our image
RUN apk add  bash

COPY . .

RUN go build -o main ./cmd

LABEL maintainer="ranniz | abaid | ayazizi"
LABEL version="1.0"
LABEL description="groupie-tracker"

EXPOSE 8080

CMD ["./main"]