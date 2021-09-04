FROM golang:1.16 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go env -w GO111MODULE=on
ENV GOOS=linux
ENV CGO_ENABLED=0
RUN go mod download
RUN go build -a -o currency-conversion .

FROM alpine:3.14
COPY --from=builder /build/currency-conversion .

ENTRYPOINT [ "./currency-conversion" ]
