FROM nginx:latest

WORKDIR /app

COPY ./nginx/startup.sh /app/.

RUN chmod +x ./startup.sh

CMD ["./startup.sh"]