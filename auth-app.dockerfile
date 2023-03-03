FROM alpine:latest

RUN mkdir /app

COPY .build/authApp /app
COPY .env /

CMD ["/app/authApp"]