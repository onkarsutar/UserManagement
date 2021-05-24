# Base Image.
FROM golang:1.16.2-alpine3.13

WORKDIR /go/src/gitlab.com/onkarsutar/UserManagement/
COPY go.mod ./  go.sum ./

RUN go mod download

WORKDIR /go/src/gitlab.com/onkarsutar/UserManagement/server
COPY ./ ./
# RUN go build 

# CMD ["./server"]

RUN go install ./

ENTRYPOINT ["/go/bin/UserManagement"]
EXPOSE 8000