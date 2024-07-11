```
// pull postgres image
docker pull postgres

// start a postgres instance
docker run --name pg-container -e POSTGRESS_PASSWORD=secret -p 5433:5432 -d postgres

// confirm container is running
docker ps

// create a database
docker exec -ti pg-container createdb -U postgres gopgtest

// check database exists
docker exec -ti pg-container psql -U postgres
(postgres=#) \c gopgtest

```
