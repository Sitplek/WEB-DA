# Employee Management API
Описание
API для управления сотрудниками, их иерархией, департаментами и фильтрацией данных. Реализован на языке Go с использованием PostgreSQL и Redis.

Содержание

1. Технологии
2. Установка и запуск
3. Архитектура проекта
4. API-эндпоинты
5. Работа с Redis
6. Пример конфигурации
7. Пример использования

## Технологии

- Язык программирования: Go
- База данных: PostgreSQL
- Кэш: Redis
- Веб-фреймворк: Gin
- Swagger-документация: Swagger UI
- Миграции базы данных: Custom Migrator

## Установка и запуск
### Шаг 1: Установка зависимостей
- Установите Go.
- Установите Docker и Docker Compose.
- Клонируйте репозиторий:
```
git clone https://github.com/Sitplek/WEB-DA.git
```

### Шаг 2: Настройка окружения
Создайте файл .env в корневой папке и заполните его следующими данными:

```
POSTGRES_USER=<user>
POSTGRES_PASSWORD=<password>
POSTGRES_DB=<mame_db>
POSTGRES_HOST=<host>
POSTGRES_PORT=<port>

REDIS_HOST=<redis_host>
REDIS_PORT=<redis_port>
```

### Шаг 3: Запуск
Запустите проект с помощью Docker Compose:


```
docker-compose up --build
```

- Frontend проекта будет доступен по адресу: http://localhost:8080
- Swagger будет доступен по адресу: http://localhost:8081

- Структура каталогов проекта:

```
├── cmd/
│   └── main.go          # Точка входа
├── internal/
│   ├── handler/         # Обработчики HTTP-запросов
│   ├── middleware/      # Миддлвары
│   ├── models/          # Определения структур данных
│   ├── service/         # Бизнес-логика
│   ├── storage/         # Работа с базой данных и Redis
│   └── migrations/      # SQL-файлы миграций
├── pkg/
│   └── migrator/        # Логика работы с миграциями
├── .env                 # Файл окружения
├── docker-compose.yml   # Конфигурация Docker Compose
├── Dockerfile           # Dockerfile для сборки API
└── docs/                # Swagger-документация
```

4. API-эндпоинты
Получение всех сотрудников

```GET /employees```

Параметры запроса:

```
Параметр		Тип			Описание
first_name		string		Имя сотрудника
last_name		string		Фамилия сотрудника
position		string		Должность
role			string		Роль
phone			string		Телефон
email			string		Электронная почта
manager_id		int			ID руководителя
department_id	int			ID департамента
```

Пример ответа:
```
[
    {
        "id": 1,
        "first_name": "John",
        "last_name": "Doe",
        "position": "Developer",
        "role": "Backend",
        "phone": "123456789",
        "email": "john.doe@example.com",
        "manager_id": null,
        "department_id": 2
    }
]
```

Получение иерархии сотрудников

```GET /employees/hierarchy```

Параметры запроса:
```
Параметр		Тип		Описание
manager_id		int		ID руководителя
department_id	int		ID департамента
```
Пример ответа:
```
[
    {
        "id": 1,
        "first_name": "Jane",
        "last_name": "Smith",
        "position": "Manager",
        "department_id": 1,
        "manager_id": null
    },
    {
        "id": 2,
        "first_name": "John",
        "last_name": "Doe",
        "position": "Developer",
        "manager_id": 1
    }
]
```

Получение доступных фильтров

```GET /filters```

Пример ответа:
```
[
    {"key": "first_name", "name": "Имя сотрудника"},
    {"key": "last_name", "name": "Фамилия сотрудника"}
]
```

5. Работа с Redis

В проекте используется Redis для кэширования данных, чтобы снизить нагрузку на базу данных.

### Примеры использования Redis:
- Кэширование всех сотрудников:
```
Ключ: employees:all
Время жизни (TTL): 60 секунд.
```
Кэширование данных с фильтрацией:
```
Формат ключа: employees:filters:<условия фильтрации>
Пример: employees:filters:first_name=John&last_name=Doe
```

Кэширование иерархии сотрудников:
```
Формат ключа: employees:hierarchy:manager_id=<id>&department_id=<id>
Пример: employees:hierarchy:manager_id=1&department_id=2
```

6. Пример конфигурации Docker Compose:
```
version: '3.9'
services:
  postgres:
    image: postgres:15
    environment:
      POSTGRES_USER: admin
      POSTGRES_PASSWORD: KqhfT89J
      POSTGRES_DB: organization_db
    ports:
      - "5432:5432"

  redis:
    image: redis:7
    ports:
      - "6379:6379"

  api_service:
    build: .
    environment:
      POSTGRES_USER: admin
      POSTGRES_PASSWORD: KqhfT89J
      POSTGRES_DB: organization_db
      POSTGRES_HOST: postgres
      REDIS_HOST: redis
      REDIS_PORT: 6379
    ports:
      - "8080:8080"

```

7. Пример использования
- Запуск проекта

Выполните:

- ```docker-compose up --build```
- Откройте Frontend проекта по адресу: http://localhost:8080
- Откройте Swagger по адресу: http://localhost:8081


Воспользуйтесь API через Swagger или Postman.
