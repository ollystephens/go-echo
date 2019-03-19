FROM golang:1.12 AS build

# Expose HTTP port and set necessary environment variables
EXPOSE 8000

# copy source code into the $GOPATH and switch to that directory
WORKDIR /go/src/app
COPY main.go .

# compile source code
RUN go install -v .

# build final image distroless for size
FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /go-echo
CMD ["/go-echo"]
