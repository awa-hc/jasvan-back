FROM postgres:latest

ENV POSTGRES_USER=jasvan
ENV POSTGRES_PASSWORD=jasvan
ENV POSTGRES_DB=jasvan 

RUN localdef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

COPY init.sql /docker-entrypoint-initdb.d/