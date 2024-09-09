FROM --platform=$TARGETOS/$TARGETARCH alpine

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY main .

CMD ["/root/main"]
