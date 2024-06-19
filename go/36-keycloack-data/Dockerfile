FROM golang:1.22.3-alpine3.20 AS builder
RUN go env -w GOPROXY=direct
RUN apk add  git
RUN apk add ca-certificates && update-ca-certificates
WORKDIR /src
COPY ./src/go.mod /src
COPY ./src/go.sum /src
RUN go mod tidy
COPY ./src /src
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /keycloak-golang

FROM golang:1.21 AS runner
COPY --from=builder /keycloak-golang /keycloak-golang
EXPOSE 1323
ENTRYPOINT [ "/keycloak-golang" ]
