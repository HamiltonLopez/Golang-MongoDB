FROM alpine:latest

RUN apk add --no-cache mongodb-tools bash curl

COPY backup.sh /backup.sh
COPY crontab /etc/crontabs/root

RUN chmod +x /backup.sh

CMD ["crond", "-f", "-L", "/dev/stdout"]
