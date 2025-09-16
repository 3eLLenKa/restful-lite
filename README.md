# Book Library API

RESTful API для управления коллекцией книг. Реализовано на Go с использованием стандартной библиотеки.

## 🚀 Возможности

- 📖 Добавление новых книг
- 📚 Получение информации о книгах
- ✅ Отметка книг как прочитанных
- 🗑️ Удаление книг из коллекции
- 📋 Получение списка всех книг

## 📋 API Endpoints

### Получить все книги
```
GET /api/v1/books
```
Возвращает массив всех книг в коллекции.

**Ответ:**
```json
[
  {
    "id": 1,
    "title": "Book Title",
    "author": "Author Name",
    "pages": 300,
    "is_read": false,
    "added_at": "2024-01-15T10:30:00Z",
    "read_at": null
  }
]
```

### Добавить новую книгу
```
POST /api/v1/books
Content-Type: application/json
```
**Тело запроса:**
```json
{
  "title": "New Book",
  "author": "Author Name",
  "pages": 250
}
```

**Ответ:** (201 Created)
```json
{
  "id": 2,
  "title": "New Book",
  "author": "Author Name",
  "pages": 250,
  "is_read": false,
  "added_at": "2024-01-15T10:35:00Z",
  "read_at": null
}
```

### Получить информацию о книге
```
GET /api/v1/book?id={id}
```
**Параметры:**
- `id` - ID книги (обязательный)

**Ответ:**
```json
{
  "id": 1,
  "title": "Book Title",
  "author": "Author Name",
  "pages": 300,
  "is_read": true,
  "added_at": "2024-01-15T10:30:00Z",
  "read_at": "2024-01-16T14:20:00Z"
}
```

### Отметить книгу как прочитанную
```
POST /api/v1/book/read?id={id}
```
**Параметры:**
- `id` - ID книги (обязательный)

**Ответ:**
```json
{
  "id": 1,
  "title": "Book Title",
  "author": "Author Name",
  "pages": 300,
  "is_read": true,
  "added_at": "2024-01-15T10:30:00Z",
  "read_at": "2024-01-16T14:20:00Z"
}
```

### Удалить книгу
```
DELETE /api/v1/book?id={id}
```
**Параметры:**
- `id` - ID книги (обязательный)

**Ответ:** 204 No Content

## 🛠️ Установка и запуск

1. **Установите Go** (версия 1.16 или выше)
2. **Скачайте проект:**
   ```bash
   git clone <repository-url>
   cd book-library-api
   ```
3. **Запустите сервер:**
   ```bash
   go run main.go
   ```
4. **Сервер будет доступен по адресу:** `http://localhost:8080`

## 📦 Зависимости

Проект использует только стандартную библиотеку Go:
- `net/http` - HTTP сервер
- `encoding/json` - работа с JSON
- `sync` - синхронизация
- `time` - работа с временем
- `strconv` - конвертация строк

## 🧪 Примеры использования

### Добавление книги
```bash
curl -X POST http://localhost:8080/api/v1/books \
  -H "Content-Type: application/json" \
  -d '{"title":"1984", "author":"George Orwell", "pages":328}'
```

### Получение всех книг
```bash
curl http://localhost:8080/api/v1/books
```

### Отметка книги как прочитанной
```bash
curl -X POST http://localhost:8080/api/v1/book/read?id=1
```

### Удаление книги
```bash
curl -X DELETE http://localhost:8080/api/v1/book?id=1
```
---