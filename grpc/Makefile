init:
	cp app.env.example app.env
	cp db.env.example db.env
	docker-compose down
	docker-compose up -d --build
	docker-compose stop
down:
	docker-compose down
gen-mock:
	.scripts/mock.sh
test:
	go test -race -covermode=atomic ./... -test.v -p=8  -count=1