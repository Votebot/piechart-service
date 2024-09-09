FROM --platform=$TARGETOS/$TARGETARCH alpine

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY $TARGETARCH/main .

CMD ["/root/main"]
