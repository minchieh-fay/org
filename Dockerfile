FROM golang:alpine AS go
WORKDIR $GOPATH/src
ENV CGO_ENABLED 0
ENV GO111MODULE on
COPY . .
#RUN swag init -o www/docs && go build -o app -ldflags "-s -w"
RUN go build -o app -ldflags "-s -w"

FROM alpine
COPY --from=go /go/src/app .
#COPY --from=go /go/src/www ./www
EXPOSE 80
ENTRYPOINT ["./app"]