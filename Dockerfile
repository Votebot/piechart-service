FROM --platform=$TARGETOS/$TARGETARCH alpine
ARG TARGETARCH
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --chmod=755 $TARGETARCH/main .

CMD ["/root/main"]
