.PHONY: dev
dev:
		air

.PHONY: tailwind-generate
tailwind-generate:
		tailwindcss -i ./index.css -o ./cmd/static/css/styles.min.css --minify

.PHONY: watch
watch:
		templ generate && $(MAKE) tailwind-generate && go build -o ./tmp/main.exe cmd/main.go 


