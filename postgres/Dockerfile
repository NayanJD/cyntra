FROM postgres:12

COPY user/rpc/model/0_user.sql /docker-entrypoint-initdb.d/0_user.sql
COPY user/rpc/model/1_seed.sql /docker-entrypoint-initdb.d/1_seed.sql
