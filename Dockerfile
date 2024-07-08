#base image
FROM golang:1.22-alpine AS build

# Set destination for COPY
RUN mkdir -p /go/src/app
ADD ./ /go/src/app
WORKDIR /go/src/app

# Copy the populate.sh script
#COPY ./init-scripts/populate.sh /go/src/app/init-scripts/

# Download Go modules
RUN apk add --update git ca-certificates build-base && rm -rf /var/cache/apk/*
RUN go build -a -o /main cmd/main.go
RUN go mod download


EXPOSE 1323

CMD ["/main"]