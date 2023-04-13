ADDRESS=$(if $(DB_ADDRESS),$(DB_ADDRESS),'localhost:9200')
DATABASE=$(if $(DB_DATABASE),$(DB_DATABASE),'postgres')
USER=$(if $(DB_USER),$(DB_USER),'user')
PASSWORD=$(if $(DB_PASSWORD),$(DB_PASSWORD),'')

run-with-deps:
	@make kill-deps
	@make deps
	@go run .

deps:
	@./scripts/deps/main.sh
	@make migrations-init migrations-up ADDRESS=$(ADDRESS)

kill-deps:
	@echo "Removing containers"
	@for DIR in docker/*; do \
		docker-compose -f $$DIR/docker-compose.yaml down --remove-orphans -v; \
	done

deps-test:
	@echo "Setting test up environment"
	@ENV=test ./scripts/deps/main.sh
	@make migrations-init migrations-up ADDRESS=localhost:9050

kill-deps-test:
	@cd docker/test && docker-compose down

migrations-init:
	@cd migrations && go run *.go -address $(ADDRESS) -user $(USER) -pass $(PASSWORD) -database $(DATABASE) init

migrations-up:
	@echo "Running migration on database: $(ADDRESS)"
	@cd migrations && go run *.go -address $(ADDRESS) -user $(USER) -pass $(PASSWORD) -database $(DATABASE) up

migrate-down:
	@echo "Running rollback migration on database: $(ADDRESS)"
	@cd migrations && go run *.go -address $(ADDRESS) -user $(USER) -pass $(PASSWORD) -database $(DATABASE) down

## migration-version: get the current version of migrations on database; use ADDRESS="" and PASSWORD="" to specify another database
migration-version:
	@cd migrations && go run *.go -address $(ADDRESS) version
