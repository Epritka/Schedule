# ------------КОММАНДЫ ДЛЯ РАЗРАБОТКИ---------------

.PHONY: server
# Запуск тестового приложения
server:
	go run ./cmd/server/main.go

.PHONY: tgbot
# Запуск телеграм бота
tgbot:
	go run ./cmd/tgbot/main.go

.PHONY: sandbox
# Запуск тестового приложения
sandbox:
	go run ./cmd/sandbox/main.go

# --------------------------------------------------