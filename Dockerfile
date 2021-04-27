FROM golang:alpine AS builder
RUN apk add --no-cache git
ADD src /src
RUN go version
RUN go env
RUN cd /src/interval_merger \
    && go mod init dummy \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o interval_merger . \
    && ls -al /src/interval_merger/interval_merger

FROM scratch
LABEL maintainer="Michael Oberdorf <michael.oberdorf@bridging-it.de>"
LABEL site.local.vendor="bridgingIT GmbH"
LABEL site.local.program.name="Interval merger"
LABEL site.local.program.version="1.0.0"
COPY --from=builder /src/interval_merger/interval_merger /interval_merger

USER 5200:5200

ENTRYPOINT ["/interval_merger"]
#CMD ["-i", "[[25,30],[2,19],[14,23],[4,8]]"]
CMD ["-h"]
