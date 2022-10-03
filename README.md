# basic-backend-app-in-go

basic-backend-app-in-go is a sample implementation of basic back-end application with RDBMS ( I use postgresql at this time ).

## Run Application

### Prepare Docker Image
```
$ docker pull postgres
$ docker run --rm -d \
     --name postgres \
     -e POSTGRES_PASSWORD=password \
     -e PGDATA=/var/lib/postgresql/data/pgdata \
     -e POSTGRES_DB=basic_backend_app_in_go
     -v ${PWD}/data:/var/lib/postgresql/data \
     -p 5432:5432 \
     postgres

$ docker exec -it basic-backend-app-in-go-postgres /bin/bash
# psql -h localhost -p 5432 -U postgres
```

Create a database.

```
CREATE DATABASE basic_backend_app_in_go;
```

### Migrate Database

```
$ go run migrate/migrate.go
```

### Run
```
$ MYAPP_DBHOST=localhost MYAPP_DBUSER=postgres MYAPP_DBPASSWORD=password MYAPP_DBNAME=basic_backend_app_in_go go run main.go
```

### API Document

after run the application, see swagger api document.
```
$ curl http://localhost:3000/swagger
```

# get
$ curl http://localhost:3000/telephones/<number>

# list
# /telephones
$ curl http://localhost:3000/telephones

# post
# vessels/<naccs_code>
$ curl -X POST -H "Content-Type: application/json" http://localhost:3000/vessels/3EDD7  -d '{"name" : "A KOU", "owner_id": "13DF"}'

# put the vessel
# vessels/<naccs_code>
$ curl -X PUT -H "Content-Type: application/json" http://localhost:3000/vessels/3EDD7  -d '{"name" : "BOKU IKEMEN", "owner_id": "44DF"}'
```

## Testing

```
$ go test ./...
```
