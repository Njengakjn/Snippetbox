start:
	go build cmd/web/main.go cmd/web/handlers.go cmd/web/routes.go cmd/web/templates.go cmd/web/middleware.go cmd/web/helpers.go
	./main.exe
run:
	go run ./cmd/web -addr=":8000"
clear:
	rm main.exe