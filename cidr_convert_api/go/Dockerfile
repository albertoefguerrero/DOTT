FROM golang:1.10-alpine

ENV GOPATH=/go:/app
RUN apk add --update git

WORKDIR /app

COPY Goopfile Goopfile
RUN go get github.com/karmakaze/goop \
    && go get github.com/gorilla/mux \
    && go get github.com/stretchr/testify/assert \
    && goop install
COPY . .
CMD ["ash", "-c", "goop go run api.go convert.go"]
# docker build -t gom .
# docker run -ti -p 8000:8000 gom
EXPOSE 8000
