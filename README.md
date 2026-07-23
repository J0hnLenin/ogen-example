# ogen-example
Репозиторий с примером создания простого сервиса с использованием генерации кода ogen

## Запуск

// TODO

## Использование ogen

Генерация файлов сервера
```bash
ogen --target=./server/internal/api/handler --package=handler --config ./api/players_api/server.ogen.yml ./api/players_api/openapi.yml
```

Генерация файлов клиента
```bash
ogen --target=./client/internal/api/players --package=players --config ./api/players_api/client.ogen.yml ./api/players_api/openapi.yml
```

## Внесённые изменения в ogen

// TODO


