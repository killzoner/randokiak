#!/bin/sh
set -eu

envsubst '${RDKAPI_API_HOST} ${RDKAPI_API_PORT} ${NGINX_LISTEN_PORT} ${ELASTIC_HOST} ${ELASTIC_PORT}' \
 < /usr/local/openresty/nginx/conf/nginx.conf.template > /usr/local/openresty/nginx/conf/nginx.conf

exec "$@"