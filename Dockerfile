FROM golang:1.16-alpine AS builder

RUN apk add --no-cache git

WORKDIR /jwt-generate-server

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# RUN CGO_ENABLED=0 go test -cover ./api ./plugin

RUN go build -o /jwt-generate-server .

FROM alpine

COPY --from=builder /jwt-generate-server/jwt-generate-server ./

RUN mkdir conf
ADD ./conf/app.dev.ini ./conf

ENTRYPOINT ["./jwt-generate-server"]

# CMD [ "-h"]