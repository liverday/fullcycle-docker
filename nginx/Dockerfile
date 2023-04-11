FROM nginx:latest

RUN apt-get update && \
  apt-get install vim -y

COPY html /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf