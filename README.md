##Поднять сервис

`docker-compose up throttle`

##Проверка работоспособности

Выполняем запрос
`curl -v -H "User-Id: 1" http://localhost:8080/v1/payments`

На 4 и следующию попытки получаем 429 ошибку, кулдайн - 1 час