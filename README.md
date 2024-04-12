docker build -t myapp .

docker network create my-network

docker run --name=banner-service-db -e POSTGRES_PASSWORD=admin -p 5432:5432 -d --rm --network=my-network postgres

docker run -d -p 8080:8080 --network=my-network myapp

postgres://postgres:admin@banner-service-db:5432/postgres?sslmode=disable

docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/api.yaml -g go-server -o /local/pkg/generated/open_api_server  && rm -f ${PWD}/pkg/generated/open_api_server/go.mod

Постарался отразить свою работу в MakeFile. Но опишу конкретные шаги. 

1) docker pull postgres - Качаем postgres
2) docker run --name=flood-controll-task -e POSTGRES_PASSWORD=admin -p 5432:5432 -d --rm postgres - Запускаем контейнер с БД
3) migrate create -ext sql -dir ./schema -seq init - Создали файлы, чтобы прописать миграцию. SQL Код для создания и удаления БД. (./schema)
4) migrate -path ./schema -database 'postgres://postgres:admin@localhost:5432/postgres?sslmode=disable' up - Подняли наши таблицы
5) migrate -path ./schema -database 'postgres://postgres:admin@localhost:5432/postgres?sslmode=disable' down - Убрали наши таблицы

