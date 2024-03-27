FROM golang:1.22
LABEL authors="pptide"

WORKDIR /app

COPY go.mod go.sum ./
COPY *.go ./
COPY *.html ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /simple-wol-server

EXPOSE 8080

CMD ["/simple-wol-server"]