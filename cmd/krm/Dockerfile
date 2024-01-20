FROM golang:1.21.0 as builder
WORKDIR /code
COPY cmd/krm cmd/krm
COPY pkg pkg
COPY go.mod .

ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -trimpath -o /app cmd/krm/main.go


FROM alpine:3
COPY --from=builder /app .
CMD ["./app"]