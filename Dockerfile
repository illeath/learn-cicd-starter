FROM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

ADD notely /usr/bin/notely

EXPOSE 8080

CMD ["/bin/sh", "-c", "PORT=8080 /usr/bin/notely --host 0.0.0.0"]
