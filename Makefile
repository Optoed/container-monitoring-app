.PHONY: init
init:
	mkdir -p backend frontend database pinger
	touch backend/main.go backend/Dockerfile backend/go.mod backend/models.go backend/handlers.go \
	pinger/main.go pinger/Dockerfile pinger/service.go pinger/container.go docker-compose.yml database/init.sql \
 	CHANGELOG.md .gitignore README.md .env