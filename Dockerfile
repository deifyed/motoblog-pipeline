FROM golang:1.16.5 AS build

WORKDIR /build

COPY go.mod .
COPY go.sum .

COPY main.go .
COPY Makefile .
COPY pkg/ ./pkg

RUN make build

FROM golang:1.16.5
WORKDIR /app
ENTRYPOINT ["./motoblog-pipeline"]

COPY --from=build /build/build/motoblog-pipeline ./
