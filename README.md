# basic-backend-app-in-go

basic-backend-app-in-go is a sample implementation of basic back-end application with RDBMS ( I use postgresql at this time ).

## How to run application

### Prepare Docker Image
```
$ docker pull postgres
$ docker run --rm -d \
     --name postgres \
     -e POSTGRES_PASSWORD=password \
     -e PGDATA=/var/lib/postgresql/data/pgdata \
     -e POSTGRES_DB=basic_backend_app_in_go \
     -v ${PWD}/data:/var/lib/postgresql/data \
     -p 5432:5432 \
     postgres

$ docker exec -it postgres /bin/bash
# psql -h localhost -p 5432 -U postgres -d basic_backend_app_in_go
```

### Migrate Database

```
$ MYAPP_DBUSER=postgres MYAPP_DBPASSWORD=password MYAPP_DBNAME=basic_backend_app_in_go go run migrate/migrate.go
```

### Run
```
$ MYAPP_DBUSER=postgres MYAPP_DBPASSWORD=password MYAPP_DBNAME=basic_backend_app_in_go go run main.go
```

## Documentation

### API Doc
after run the application, see swagger api document.
```
$ curl http://localhost:3000/swagger
```

### Sample Request

Get a telephone information

```
$ curl http://localhost:3000/telephones/090222233333
```

List all telephone informations

```
$ curl http://localhost:3000/telephones
```

Post a telephone information

```
$ curl -X POST -H "Content-Type: application/json" http://localhost:3000/telephones/08033332222 -d '{"owner_id" : "2", "icc_id": "111111111111111"}'
```

```
$ curl -X PUT -H "Content-Type: application/json" http://localhost:3000/telephones/08033332222  -d '{"owner_id" : "2", "icc_id": "222222222222222"}'
```

## Testing

```
$ make test
```
