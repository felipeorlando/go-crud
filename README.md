# go-crud
This is a simple crud developed usign Go and Mongo, with some clean archtecture concepts.

You can CRUD users on the following endpoints:

```
GET /users
GET /users/{id}
POST /users
PUT /users/{id}
DELETE /users/{id}
```

To run the project you can use `docker-compose` or the `Makefile` (don't forget to change the configs on `config/config.go`). And it will run on port 3000.

## tests
To run the tests you must be runnig MongoDB locally (don't forget to change the configs on `config/config.go`) and just run the following command:

```sh
make test
```

