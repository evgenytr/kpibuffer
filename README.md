# kpibuffer

# Как запускать (docker-compose)
Запустить Docker, в терминале в корневой директории проекта выполнить команду
```
docker-compose up 
```

Сервер поддерживает три ручки:
- GET http://localhost:8080/ выведет количество неотправленных записей
- POST http://localhost:8080/fact принимает JSON c одной записью вида
```json
{
"period_start":"2024-05-01",
"period_end":"2024-05-31",
"period_key":"month",
"indicator_to_mo_id":227373,
"indicator_to_mo_fact_id":0,
"value":1,
"fact_time":"2024-05-31",
"is_plan":0,
"auth_user_id":40,
"comment": "buffer"
}
```
- POST http://localhost:8080/facts принимает JSON c массивом записей