#!/bin/sh

docker build --rm -t trigun117/news_aggregator .
echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
docker push trigun117/news_aggregator