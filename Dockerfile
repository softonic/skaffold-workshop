FROM caddy:2.3.0-alpine

ADD src /usr/share/caddy/

ADD empty/ /tmp
