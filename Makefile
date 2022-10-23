
.PHONY: rsue_init
rsue_init:
	python3 parsers/rsue/init.py

.PHONY: tgbot
tgbot:
	python3 tgbot/main.py