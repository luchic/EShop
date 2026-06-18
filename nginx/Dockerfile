FROM nginx:alpine

EXPOSE 80

USER root

RUN addgroup -S nginxgroup && adduser -S nginxuser -G nginxgroup \
    && mkdir -p /var/www/html \
    && chown -R nginxuser:nginxgroup /var/www/html \
    && chown -R nginxuser:nginxgroup /etc/nginx \
    && chown -R nginxuser:nginxgroup /var/cache/nginx \
    && chown -R nginxuser:nginxgroup /var/log/nginx \
    && chown -R nginxuser:nginxgroup /run

USER nginxuser:nginxgroup

CMD ["nginx", "-g", "daemon off;"]