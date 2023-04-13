# ================================================================
# build application
# ================================================================
FROM golang:1.15-alpine AS build

# build outside of GOPATH (simpler when using Go modules)
WORKDIR /src

# download dependencies
COPY go.mod go.sum ./
RUN apk add git
RUN go mod download

# build application
COPY . .
RUN go install

# ================================================================
# create working image
# ================================================================
FROM alpine:latest
WORKDIR /root/
COPY --from=build /go/bin/leaderboard /usr/local/bin/
RUN mkdir ./config
COPY config/config.yaml ./config/
EXPOSE 8080 9800

CMD ["leaderboard", "start", "--type=api"]
