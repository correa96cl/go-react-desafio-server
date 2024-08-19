# go-react-desafio

docker-compose down -v 


go mod tidy


sqlc generate -f ./internal/store/pgstore/sqlc.yaml

go generate ./...
