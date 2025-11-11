FROM golang:1.25-alpine AS builder
WORKDIR /src
#COPY ./vendor ./vendor
COPY ./go.mod ./go.mod
#COPY ./go.sum ./go.sum
COPY ./main.go ./main.go
COPY ./index.html ./index.html
RUN CGO_ENABLED=0 GOOS=linux \
    go build -o server ./main.go


FROM alpine
WORKDIR /app
COPY --from=builder ./src/server ./
ENV PATH=$PATH:/app
CMD ./server


