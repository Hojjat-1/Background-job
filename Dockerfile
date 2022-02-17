FROM golang:1.16.4 AS go-build

ENV GO111MODULE=on

WORKDIR /go/src/gitlab.com/utopiops-water/test-image
ENV GONOSUMDB=gitlab.com/utopiops-water
ENV GOPRIVATE=gitlab.com/utopiops-water/*

COPY go.mod .
COPY go.sum .

RUN go mod tidy
RUN go mod download

COPY . .

# RUN go run main.go
RUN go build -o main . 

# EXPOSE 8080

CMD ["/go/src/gitlab.com/utopiops-water/test-image/main"]
