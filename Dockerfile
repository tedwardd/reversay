# syntax=docker/dockerfile:1

FROM golang:1.26 AS build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 go build -o /out/echo-reverse .

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/echo-reverse /echo-reverse
EXPOSE 8080
ENTRYPOINT ["/echo-reverse"]
CMD ["-host", "0.0.0.0", "-port", "8080"]
