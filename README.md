# Планировщик задач

Итоговый проект курса. Веб-сервер на Go с REST API и SQLite, фронтенд лежит в `web`.

## Задания со звёздочкой

- Правила повторения `w` и `m` (не только `d`/`y`)
- Поиск: `GET /api/tasks?search=...`
- Аутентификация: `TODO_PASSWORD`, `POST /api/signin`, JWT
- Docker

## Запуск

```
go run main.go
```

Переменные окружения (`.env` в корне, необязательные):

```
TODO_PORT=7540
TODO_DBFILE=scheduler.db
TODO_PASSWORD=password
JWT_SECRET=supersecretsecret
```

Если `TODO_PASSWORD` не задан — авторизация выключена. Адрес: `http://localhost:7540`.

## Тесты

```
go test ./tests
```

В `tests/settings.go`:

- `FullNextDate = true`
- `Search = true`
- `Token` — токен из `/api/signin`, если задан `TODO_PASSWORD`

## Docker

```
make docker-build
make docker-run
make docker-stop
```

Или руками:

```
docker build -t scheduler .

docker run -d \
  --name scheduler \
  -p 7540:7540 \
  --env-file .env \
  -e TODO_DBFILE=/app/data/scheduler.db \
  -v "$(pwd)/data:/app/data" \
  scheduler
```

База пишется в `./data` на хосте.
