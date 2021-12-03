FROM postgres:alpine
ADD 1_init.sql /docker-entrypoint-initdb.d
ADD init-user-db.sh /docker-entrypoint-initdb.d
RUN chmod a+r /docker-entrypoint-initdb.d/*
EXPOSE 5432