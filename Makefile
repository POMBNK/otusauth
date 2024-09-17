generate:
	go generate ./...
	go run docs/merger/main.go
	statik -src=docs/dist -include=*.html,*.css,*.js,*.png,*.json

.PHONY: generate