[![Build Status](https://travis-ci.com/trigun117/news_aggregator.svg?branch=master)](https://travis-ci.com/trigun117/news_aggregator) [![codecov](https://codecov.io/gh/trigun117/news_aggregator/branch/master/graph/badge.svg)](https://codecov.io/gh/trigun117/news_aggregator) [![Go Report Card](https://goreportcard.com/badge/github.com/trigun117/news_aggregator)](https://goreportcard.com/report/github.com/trigun117/news_aggregator)

# news_aggregator

News Aggregator 

![site screenshot](https://github.com/trigun117/news_aggregator/blob/master/image.JPG)
![site screenshot](https://github.com/trigun117/news_aggregator/blob/master/image1.JPG)

On http://ua-news.ml you can find live example

# Get Started

For start, build docker image from Dockerfile and run with this command
```
docker run -d \
-p 80:80 \
-e LINK=link_to_news_api \
-e EF=email_from \
-e EFL=email_from_login \
-e EFP=email_from_password \
-e ET=email_to \
--restart always \
image_name
```
and then open http://localhost or http://your_server_ip

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
