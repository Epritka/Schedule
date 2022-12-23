ifeq (migrate,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

define read_env_file
	$(eval ENV_FILE := ./.env)
	@echo " - setup env $(ENV_FILE)"
	$(eval include ./.env)
	$(eval export sed 's/=.*//' ./.env)
endef


# --------------КОММАНДЫ ДЛЯ СБОРКИ-----------------

.PHONY: build
# Сборка исполняемых файлов
build:
	go build -o ./deploy/bin/server ./cmd/server/main.go
	go build -o ./deploy/bin/migrate ./cmd/migrate/main.go

# ------------КОММАНДЫ ДЛЯ РАЗРАБОТКИ---------------

.PHONY: migrate
# Запуск миграций goose
# https://github.com/pressly/goose
migrate:
	go run ./cmd/migrate/main.go $(RUN_ARGS)

.PHONY: server
# Запуск локального сервера
server:
	go run ./cmd/server/main.go

.PHONY: install
# Развертывание проекта
install:
	cp -n ./.env.example ./.env
	go mod download

.PHONY: test
test:
	go test -cover ./...

# --------------------------------------------------
