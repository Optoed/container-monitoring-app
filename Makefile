.PHONY: init
init:
	mkdir -p backend frontend database pinger
	touch backend/main.go backend/Dockerfile pinger/main.go pinger/Dockerfile docker-compose.yml database/init.sql \
 	CHANGELOG.md .gitignore README.md .env