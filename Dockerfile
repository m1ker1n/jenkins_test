FROM golang:1.22.2-alpine3.19

# creates /app dir in container and cd's to it
WORKDIR /app

# not using ./app, because we're here already 
# copying dependency files
COPY go.mod go.sum ./

RUN go mod download

# copying source code
COPY *.go ./

# building application  
RUN go build -o /go-echo

# The EXPOSE instruction doesn't actually publish the port
# It functions as a type of documentation
EXPOSE 8080

# running application
CMD ["/go-echo"]