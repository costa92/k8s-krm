FROM golang:1.21.0 as builder
WORKDIR /code
ADD . .
RUN ls -la

RUN go mod download
RUN go build -gcflags="all=-N -l" -trimpath -o /app cmd/krm/main.go

FROM alpine:3
COPY --from=builder /app .
RUN ls -la
RUN ./app
CMD ["./app"]