FROM alpine

COPY ./news_aggregator .
COPY ./templates ./templates
RUN apk update && apk add --no-cache ca-certificates nano

ENTRYPOINT [ "./news_aggregator" ]