FROM golang:1.17.2-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/hotel-rental /app/
COPY --from=builder /build/application.yaml /app/
WORKDIR /app
EXPOSE 9090
CMD ["./hotel-rental"]
