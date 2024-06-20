FROM golang:1.22.4-alpine3.20 AS builder
RUN go env -w GOPROXY=direct
RUN apk add  git
RUN apk add ca-certificates && update-ca-certificates
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /keycloak-golang

FROM golang:1.21 AS runner
COPY --from=builder /keycloak-golang /keycloak-golang
ENTRYPOINT [ "/keycloak-golang" ]
