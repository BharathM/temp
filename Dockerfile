FROM golang:1.16-alpine
RUN apk update
RUN apk add net-tools
RUN apk add curl

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY 1MB.pdf ./
COPY 1MB.docx ./

RUN go build -o /serve-file

EXPOSE 8092

CMD [ "/serve-file" ]
