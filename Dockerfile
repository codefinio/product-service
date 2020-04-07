FROM golang:alpine as builder
WORKDIR "/app"
RUN apk --no-cache add make git gcc libc-dev ca-certificates
ENV GO11MODULE=on
COPY . .
RUN make build

FROM scratch
COPY --from=builder /app/bin/server /app/bin/server
CMD ["/app/bin/server"]