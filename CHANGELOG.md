# Changelog

Все заметные изменения в этом проекте будут задокументированы в этом файле.

## [0.1.0] - 03-02-2025
### Добавлено
- Создан Makefile

## [0.2.0] - 04-02-2025
### Добавлено
- Написан скрипт init.sql создания таблицы containers в бд

## [0.3.0] - 04-02-2025
### Добавлено
- Добавлена в models.go модель Container
- Сервис backend разработан и ожидает тестирования

## [0.4.0] - 05-02-2025
### Добавлено
- Сервис Pinger разработан и ожидает тестирования
### Изменено
- Добавлено в таблицу бд ping_time

## [0.5.0] - 05-02-2025
### Изменено
- В backend/handlers теперь передаем еще и ping_duration
- Переработана структура backend
- Добавлены теги db в model Container
- Найден баг в handler AddContainer, не был вызван своевременно return

## [0.6.0] - 07-02-2025
### Добавлено
- Бд успешно поднимается и функционирует в docker, Pinger и Backend функционируют,
  но запрос от Pinger не доходит до сервиса Backend

## [0.7.0] - 07-02-2025
### Добавлено
- Pinger успешно отправляет POST-запросы в Backend, 
- Backend успешно сохраняет данные в db,
- GET-запрос также успешно протестирован

## [0.8.0] - 07-02-2025
### Добавлено
- Инициализирован базовый React (TypeScript) проект

## [0.9.0] - 07-02-2025
### Добавлено
- Фронтенд частично готов. Информация обновляется действительно каждые 10 секунд.
- Прописан CORS в backend
### Нужно сделать:
- Фронтенд не отображает Ping duration, status, last ping time
- Нужно сделать рефакторинг кода, подчистить заглушки, использовать os.Gotenv
- Добавить фронтенд в контейнер

## [0.9.1] - 07-02-2025
### Изменено
- Использованы параметры из enviroment (описано в docker-compose) контейнеров
  вместо захардкоженных переменных

## [0.9.2] - 08-02-2025
### Изменено
- Теперь frontend отображает last ping time и ping duration корректно

## [0.10.0] - 08-02-2025 - ветка containerize-frontend
### Добавлено
- ветка containerize-frontend
- контейнеризация frontend (пока frontend не видит backend)

## [0.10.1] - 08-02-2025 - ветка containerize-frontend
### Изменено
- убран volumes для сервиса db из docker-compose.yml
- **Запускай с очищением кэша и томов volume** (чтобы избежать ошибки с init.sql):
  ```bash
  docker-compose down --volumes
  docker-compose up --build
  ```

## [0.10.2] - 08-02-2025 - ветка containerize-frontend
### Изменено
- добавлен yarn.lock в репозиторий (и убран из .gitignore)

## [0.10.3] - 08-02-2025 - ветка containerize-frontend
### Изменено
- разрешен конфликт с политикой CORS как для frontend (поднимаемом в контейнере), так и на локальном frontend сервисе

## [0.10.4] - 08-02-2025 - ветка containerize-frontend -> (merge) main
### Изменено
- в postgres ping_duration теперь храним в наносекундах как BIGINT
- во frontend отображаем в ms (делим на 1e6)
- container_handler отправляет просто в наносекундах (ns)

**merge containerize-frontend в main!**

## [0.11.0] - 09-02-2025 - ветка rabbitmq
### Добавлено
- rabbitmq в docker-compose
- producer.go (rabbitMQ producer) в pinger
- rabbitMQ consumer в backend
### Изменено
- Переработана архитектура backend (MVC)
