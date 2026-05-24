# Local database for development

```shell
docker compose up -d
```

## Access

Replace `${POSTGRES_USER}` and `${POSTGRES_PASSWORD}` with the values from `docker-compose.yml`.

```shell
POSTGRES_USER=$(yq e ".services.db.environment.POSTGRES_USER" docker-compose.yml)
POSTGRES_PASSWORD=$(yq e ".services.db.environment.POSTGRES_PASSWORD" docker-compose.yml)
docker compose exec db postgres -u ${POSTGRES_USER} -p${POSTGRES_PASSWORD}
```

### Environment Variables

```shell
export DB_HOST=localhost
export DB_USERNAME=sbcntrapp
export DB_PASSWORD=password
export DB_NAME=sbcntrapp
export DB_CONN=1
```
