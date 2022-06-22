FROM postgres:12

ENV POSTGRES_USER docker
ENV POSTGRES_PASSWORD docker
ENV POSTGRES_DB docker

EXPOSE 5432

COPY db/postgres/* /docker-entrypoint-initdb.d/
