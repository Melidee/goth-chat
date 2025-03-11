build-app:
	@go build -o bin/app ./app/

css: 
	@npx @tailwindcss/cli -i ./assets/app.css -o ./public/assets/app.css --watch

run: build-app
	@templ generate
	@./bin/app

dev: 
	air

clean: 
	@rm -rf bin