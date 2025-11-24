build:
	@echo "ビルド中..."
	@go build -o todo cmd/main.go
	@echo "✅ ビルド成功！"
	@echo ""
	@./todo

run: build
	@./todo list