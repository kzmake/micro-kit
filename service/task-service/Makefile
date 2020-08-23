DB_DSN ?= test:test@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True

.PHONY: db
db: db/version

.PHONY: db/version
db/version:
	migrate -path migration/ -database "mysql://$(DB_DSN)" version

.PHONY: db/up
db/up:
	migrate -path migration/ -database "mysql://$(DB_DSN)" up

.PHONY: db/down
db/down:
	migrate -path migration/ -database "mysql://$(DB_DSN)" down
