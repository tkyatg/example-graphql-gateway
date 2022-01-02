init:
	cp graphql/.env.app.example graphql/.env.app
	cp grpc/.env.app.example grpc/.env.app
	cp postgres/.env.db.example postgres/.env.db
	docker-compose build
serve:
	docker-compose up