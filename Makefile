# Makefile

# 🟢 Запуск приложения
run:
	docker-compose up --build

# 🔴 Остановка и удаление контейнеров
down:
	docker-compose down -v

# 🔁 Перезапуск (обновить контейнеры)
restart:
	docker-compose down -v
	docker-compose up --build

# ✅ Юнит-тесты
test:
	go test ./...

# 📦 Тесты по отдельным папкам
test-handler:
	go test ./internal/handler

test-task:
	go test ./internal/task

test-storage:
	go test ./internal/storage

# 🧹 Очистка всего (контейнеры, кэш, тома)
clean:
	docker-compose down -v
	docker system prune -f
