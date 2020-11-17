# go-crud
This is a simple crud developed usign Go and Mongo, with some clean archtecture concepts.

You can CRUD users on the following endpoints:

```
GET /api/v1/users
GET /api/v1/users/{id}
POST /api/v1/users
PUT /api/v1/users/{id}
DELETE /api/v1/users/{id}
```

To run the project you can use `docker-compose` or the `Makefile`. And it will run on port 3000.
