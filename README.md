docker build -t myapp .

docker network create my-network

docker run --name=banner-service-db -e POSTGRES_PASSWORD=admin -p 5432:5432 -d --rm --network=my-network postgres

docker run -d -p 8080:8080 --network=my-network myapp

postgres://postgres:admin@banner-service-db:5432/postgres?sslmode=disable

Postgres
Постарался отразить свою работу в MakeFile. Но опишу конкретные шаги.

docker pull postgres - Качаем postgres
docker run --name=flood-controll-task -e POSTGRES_PASSWORD=admin -p 5432:5432 -d --rm postgres - Запускаем контейнер с БД
migrate create -ext sql -dir ./schema -seq init - Создали файлы, чтобы прописать миграцию. SQL Код для создания и удаления БД. (./schema)
migrate -path ./schema -database 'postgres://postgres:admin@localhost:5432/postgres?sslmode=disable' up - Подняли наши таблицы
migrate -path ./schema -database 'postgres://postgres:admin@localhost:5432/postgres?sslmode=disable' down - Убрали наши таблицы
