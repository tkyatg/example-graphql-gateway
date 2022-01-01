init:
	cp graphql/.env.app.example graphql/.env.app
	cp grpc/.env.app.example grpc/.env.app
	cp postgres/.env.db.example postgres/.env.db
	docker-compose build
serve:
	docker-compose up
gen-gql:
	docker compose exec graphql go get github.com/99designs/gqlgen/internal/imports@v0.14.0
	docker compose exec graphql go get github.com/99designs/gqlgen/internal/code@v0.14.0
	docker compose exec graphql go get github.com/99designs/gqlgen/cmd@v0.14.0
	docker compose exec graphql go run github.com/99designs/gqlgen generate