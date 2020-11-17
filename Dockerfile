FROM golang:1.15.3-alpine AS builder
RUN apk -u add make git

WORKDIR /src
COPY --chown=nobody:nobody . .
RUN export GO111MODULE=on
RUN make

FROM alpine
WORKDIR /
COPY --from=builder ./src/cmd/api .
EXPOSE 3000
CMD ["./api"]
