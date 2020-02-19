FROM node:12.6-alpine as builder

WORKDIR /build

# copies all files, except those marked by .dockerignore
COPY . .
RUN yarn install --frozen-lockfile
RUN yarn run build:prod

RUN yarn install --frozen-lockfile

FROM openresty/openresty:1.15.8.2-alpine-nosse42

ENV TERM=xterm
WORKDIR /usr/local/openresty/nginx/conf

COPY --from=builder /build/dist /var/www/html
ADD ./nginx/nginx.conf /usr/local/openresty/nginx/conf/nginx.conf.template
ADD ./nginx/docker-entrypoint.sh /
RUN chmod a+x /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["/usr/local/openresty/bin/openresty", "-g", "daemon off;", "-g", "daemon off;"]

# Update apk
RUN apk update && apk add curl

# Install envsubst command for replacing config files in system startup
RUN set -x && \
    apk add libintl && \
    apk add --virtual build_deps gettext &&  \
    cp /usr/bin/envsubst /usr/local/bin/envsubst && \
    apk del build_deps

# Clean apk cache
RUN rm -rf /var/cache/apk/*

HEALTHCHECK --interval=30s --timeout=3s CMD curl -f http://localhost/health-check || exit 1