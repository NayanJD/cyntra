FROM debezium/postgres:12-alpine

COPY user/rpc/model/0_user.sql /docker-entrypoint-initdb.d/0_user.sql
COPY user/rpc/model/1_seed.sql /docker-entrypoint-initdb.d/1_seed.sql
COPY product/rpc/model/0_product.sql /docker-entrypoint-initdb.d/2_seed.sql