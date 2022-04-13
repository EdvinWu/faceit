FROM golang:1.18.1-alpine3.14 as build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o faceit-test cmd/faceit-test/main.go

FROM scratch

COPY --from=build /build/faceit-test /opt/faceit-test
COPY --from=build /build/migration /opt/migration
COPY --from=build /build/app-config-docker.yaml /opt/app-config.yaml

WORKDIR /opt
CMD ["/opt/faceit-test"]
