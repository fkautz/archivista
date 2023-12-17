#!/bin/sh

DB_SECRET_VALUE=$(aws secretsmanager --region=us-east-1 get-secret-value --secret-id "$ARCHIVISTA_DB_SECRET_ARN")

DATABASE_USER="$(echo $DB_SECRET_VALUE | jq -r '.SecretString | fromjson | .username')"
DATABASE_PASSWORD="$(echo $DB_SECRET_VALUE | jq -r '.SecretString | fromjson | .password')"
DATABASE_HOST="$(echo $DB_SECRET_VALUE | jq -r '.SecretString | fromjson | .host')"
DATABASE_DBNAME="$(echo $DB_SECRET_VALUE | jq -r '.SecretString | fromjson | .dbname')"

export ARCHIVISTA_LISTEN_ON="tcp://0.0.0.0:8082"
export ARCHIVISTA_SQL_STORE_CONNECTION_STRING="$DATABASE_USER:$DATABASE_PASSWORD@tcp($DATABASE_HOST)/$DATABASE_DBNAME"
export ARCHIVISTA_SQL_STORE_BACKEND='MYSQL'
export ARCHIVISTA_STORAGE_BACKEND='BLOB'
export ARCHIVISTA_BLOB_STORE_ENDPOINT='s3.amazonaws.com'

/usr/bin/archivista