FROM ubuntu:latest

RUN mkdir /app

COPY /node/run.sh /tmp/run.sh
COPY /node/build/main /app

RUN chmod +x /tmp/run.sh

ENTRYPOINT ["sh","/tmp/run.sh"]
