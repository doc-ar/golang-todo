###################################
# Generate templates
prod/templ:
	templ generate

# 
prod/tailwind: 
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

prod/build:
	go build -o ./tmp/main ./cmd
###################################

# Run Tailwindcss watcher
dev/tailwind:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify --watch

# Run air (Go + templ) hot reload
dev/air:
	air
###################################


# Run in dev mode
dev:
	make -j2 dev/tailwind dev/air

# Generate production build
prod:
	make prod/tailwind prod/templ prod/build
