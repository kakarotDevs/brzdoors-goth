run: build
	@./bin/app

build:
	@go build -o bin/app .

dev:
	@echo "Starting dev environment..."
	npx concurrently \
		"npx tailwindcss -i views/css/app.css -o public/styles.css --watch" \
		"templ generate --watch --proxy=http://localhost:3000" \
		"air"