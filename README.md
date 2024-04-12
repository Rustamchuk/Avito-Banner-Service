docker build -t myapp .

docker network create my-network

docker run --name=banner-service-db -e POSTGRES_PASSWORD=admin -p 5432:5432 -d --rm --network=my-network postgres

docker run -d -p 8080:8080 --network=my-network myapp

postgres://postgres:admin@banner-service-db:5432/postgres?sslmode=disable
