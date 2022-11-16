FROM ubuntu:22.04
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV DEBIAN_FRONTEND=noninteractive
ENV PGVER 10
ENV POSTGRES_HOST /var/run/postgresql/
ENV POSTGRES_PORT 5432
ENV POSTGRES_DB db_finance
ENV POSTGRES_USER user
ENV POSTGRES_PASSWORD 123


RUN apt-get update && apt-get install -y postgresql-$PGVER

USER user

COPY build/createTableScript.sql createTableScript.sql

RUN service postgresql start &&\
    psql -U user -c "ALTER USER user PASSWORD '123';" &&\
    psql -U user -c 'CREATE DATABASE "db_finance";' &&\
    psql -U user -d db_finance -a -f createTableScript.sql &&\
    service postgresql stop

#...