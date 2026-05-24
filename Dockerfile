ARG BUILD_PLATFORM=linux/amd64
FROM --platform=$BUILD_PLATFORM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

ADD notely /usr/bin/notely

CMD ["notely"]
