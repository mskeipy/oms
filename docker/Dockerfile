ARG GOLANG_VERSION=1.22.12-alpine3.21

####################
### GOLANG BASE
####################
FROM public.ecr.aws/docker/library/golang:${GOLANG_VERSION} AS base

WORKDIR /app

ENV GOFLAGS="-buildvcs=false"
ENV GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOPATH=/go
ENV GOMODCACHE=/go/pkg/mod

RUN export GOPATH=$HOME/go && \
    export PATH=$PATH:/usr/local/go/bin

EXPOSE 8080

####################
### GOLANG DEVELOPMENT
####################
FROM base AS development

RUN go install github.com/cosmtrek/air@v1.41

RUN export PATH=$PATH:$(go env GOPATH)/bin

####################
### GOLANG BUILDER
####################
FROM development AS builder

ENV GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn
ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app
COPY . .

RUN go mod tidy \
    && go build -buildvcs=false -a -installsuffix cgo -o ./build/serve ./cmd/server/main.go \
    && go build -buildvcs=false -a -installsuffix cgo -o ./build/migrate ./internal/infrastructure/database/migration/migrate.go

####################
### GOLANG FINAL APPLICATION
####################
FROM public.ecr.aws/docker/library/golang:${GOLANG_VERSION} AS application

WORKDIR /app
COPY --from=builder /app/.env* /app/
COPY --from=builder /app/build/serve /app/build/migrate /app/
COPY --from=builder /app/internal/infrastructure/database/migration/migrations /app/internal/infrastructure/database/migration/migrations

RUN chmod +x /app/serve /app/migrate