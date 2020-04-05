FROM golang:1.14.0 as build

WORKDIR /go/src/github.com/dtamura/samplecontroller
COPY . .
RUN go get -d -v  .
RUN cd $GOPATH/src/k8s.io/klog && git checkout v0.4.0
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app

FROM alpine:3.11.5
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /go/src/github.com/dtamura/samplecontroller/app .

CMD [ "./app" ]