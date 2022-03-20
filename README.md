Law site:
-
make shemas:
> migrate create -ext sql -dir ./shemas -seq init

make docker container:
> docker run --name=law_db -e POSTGRES_PASSWORD="postgres" -p 5436:5432 -d --rm postgres

make migrations:
> migrate -path ./shemas -database 'postgresql://postgres:postgres@localhost:5436/postgres?sslmode=disable' up

# rest_test
