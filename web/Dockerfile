FROM nginx:latest

COPY static /usr/share/nginx/html

RUN rm /etc/nginx/conf.d/default.conf
COPY docker/nginx.conf /etc/nginx/conf.d/default.conf