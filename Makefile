install-mockgen:
	go get github.com/golang/mock/mockgen@v1.6.0

generate:
	go generate ./...

build:
	docker-compose build

up:
	docker-compose up

down:
	docker-compose down

order-migrate-create:
	docker-compose exec database migrate create -ext sql -dir db/migrations/order -seq yyy

order-migrate-up:
	docker-compose exec database migrate -path=db/migrations/order -database="mysql://root:password@tcp(database:3306)/order" up

order-migrate-down:
	docker-compose exec database migrate -path=db/migrations/order -database="mysql://root:password@tcp(database:3306)/order" down

product-migrate-create:
	docker-compose exec database migrate create -ext sql -dir db/migrations/product -seq xxx

product-migrate-up:
	docker-compose exec database migrate -path=db/migrations/product -database="mysql://root:password@tcp(database:3306)/product" up

product-migrate-down:
	docker-compose exec database migrate -path=db/migrations/product -database="mysql://root:password@tcp(database:3306)/product" down

setting-migrate-create:
	docker-compose exec database migrate create -ext sql -dir db/migrations/setting -seq yyy

setting-migrate-up:
	docker-compose exec database migrate -path=db/migrations/setting -database="mysql://root:password@tcp(database:3306)/setting" up

setting-migrate-down:
	docker-compose exec database migrate -path=db/migrations/setting -database="mysql://root:password@tcp(database:3306)/setting" down
