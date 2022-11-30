# Ручная установка  базы Postgres и первая миграция данных migrate

# Скачивание образа из DockerHub
pgPull:
	docker pull postgres

# Запуск контейнера
## --name - имя контейнера
## -d - запуск в фоновом режиме
## --rm - удаление контейнера при его остановке
pgRun:
	docker run --name=go-todo-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres

# Посмотреть список запущенных контейнеров
show:
	docker ps

# Подключиться к контейнеру в интерактивном режиме
# Пример использования: > make exec c=1234567890
#
# Посмотреть таблицу базы данных
# > psql -U postgres
# postgres> \d
#
# exit exit
exec:
	docker exec -it $(c) /bin/bash

# Установим утилиту migrate для выполнения процедуры миграции
# Соберем по документации через homebrew
## Установим cli
## brew install golang-migrate
mgGet:
	go get -u github.com/golang-migrate/migrate

# Создадим файлы миграции
# для файлов .sql
# в каталоге ./schema
# c именем init
#
# Система минграции,
# позволяет вносит новые изменения в структуру базы 
# или делать откат к предыдущей версии
# без нарушений целостности данных
mgInit:
	migrate create -ext sql -dir ./sсhema -seq init

# Выполним тестовую миграцию для файлов .up
mgUp:
	migrate -path ./sсhema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

# Выполним тестовую миграцию для файлов .down
# и удалим все созданные таблицы
mgDown:
	migrate -path ./sсhema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down

# Генерация swagger документации через swaggo/swag
swInit:
	swag init -g cmd/main.go

.PHONY: pgPull pgRun show \
		mgGet mgInit mgUp mgDown \
		swInit

SHELL = /bin/sh
RAND = $(shell echo $$RANDOM)