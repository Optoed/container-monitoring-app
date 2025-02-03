.PHONY: init
init:
	mkdir -p backend frontend database pinger
	touch README.md backend/main.go backend/Dockerfile pinger/main.go pinger/Dockerfile docker-compose.yml CHANGELOG.md .gitignore