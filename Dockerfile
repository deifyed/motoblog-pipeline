FROM golang:1.16.5 AS build

WORKDIR /src

COPY go.mod .
COPY go.sum .

COPY main.go .
COPY Makefile .
COPY pkg/ ./pkg

RUN make release

FROM golang:1.16.5
WORKDIR /app
ENTRYPOINT ["./motoblog-pipeline"]

COPY --from=build /src/build/release/motoblog-pipeline ./
