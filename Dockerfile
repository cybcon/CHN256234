FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN mkdir /app
# RUN git config --global http.sslVerify false 
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o interval_merger .

FROM scratch
LABEL maintainer="Michael Oberdorf <michael.oberdorf@bridging-it.de>"
LABEL site.local.vendor="bridgingIT GmbH"
LABEL site.local.program.name="Interval merger"
LABEL site.local.program.version="1.0.0"
COPY --from=builder /app/inteval_merger /app/interval_merger

USER 5200:5200

WORKDIR /app

ENTRYPOINT ["/app/interval_merger"]
CMD ["-i", "[25,30] [2,19] [14, 23] [4,8]"]
