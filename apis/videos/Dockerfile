FROM golang:alpine as builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

COPY . .

RUN go get .

RUN go build -o app .

# deployment image
FROM scratch

WORKDIR /bin/

# copy app from builder
COPY --from=builder /build/app .
COPY --from=builder /build/.test-env .

# change in .test-env, make file too
EXPOSE 5000

CMD [ "./app" ]