FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN mkdir /src
# RUN git config --global http.sslVerify false 
ADD src /src 
#RUN ls -al; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o interval_merger ./src/interval_merger
RUN GOOS=linux go build -a -installsuffix cgo -o interval_merger ./src/interval_merger

FROM scratch
LABEL maintainer="Michael Oberdorf <michael.oberdorf@bridging-it.de>"
LABEL site.local.vendor="bridgingIT GmbH"
LABEL site.local.program.name="Interval merger"
LABEL site.local.program.version="1.0.0"
COPY --from=builder inteval_merger interval_merger

USER 5200:5200

ENTRYPOINT ["interval_merger"]
#CMD ["-i", "[[25,30],[2,19],[14,23],[4,8]]"]
CMD ["-h"]
