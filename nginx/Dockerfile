FROM nginx

COPY nginx.conf.template /tmp/nginx.conf.template

CMD /bin/bash -c "envsubst '\$API_HOST' < /tmp/nginx.conf.template > /etc/nginx/nginx.conf && nginx -g 'daemon off;'"
