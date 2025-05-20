

---

````md
# ☕ Coffee Shop API

A RESTful API for managing coffee cards. Built with Go, Chi router, SQLx, and PostgreSQL.

---

## 🚀 Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/your-name/coffee-shop.git
   cd coffee-shop
````

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Make sure PostgreSQL is running and accessible.

4. Create the table:

   ```sql
   CREATE TABLE coffee_card (
     id SERIAL PRIMARY KEY,
     img TEXT NOT NULL,
     coffee_name TEXT NOT NULL,
     country TEXT NOT NULL,
     price NUMERIC NOT NULL,
     recommended BOOLEAN NOT NULL,
     description TEXT NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

5. Start the server:

   ```bash
   go run main.go
   ```

---

## 🧱 Dependencies

* [`chi`](https://github.com/go-chi/chi) — lightweight HTTP router
* [`sqlx`](https://github.com/jmoiron/sqlx) — extensions to database/sql for struct support
* [`go-playground/validator`](https://github.com/go-playground/validator) — struct validation

---

## 📌 API Endpoints

| Method | URL         | Description             |
| ------ | ----------- | ----------------------- |
| GET    | /cards      | Get all coffee cards    |
| GET    | /cards/{id} | Get a specific card     |
| POST   | /cards      | Create a new card       |
| PATCH  | /cards/{id} | Update an existing card |
| DELETE | /cards/{id} | Delete a card           |

---

## 🧪 Sample Request

```http
POST /cards
Content-Type: application/json

{
  "img": "https://example.com/image.jpg",
  "coffeeName": "Arabica Premium",
  "country": "Colombia",
  "price": 12.50,
  "recommended": true,
  "description": "Rich taste with hints of chocolate"
}
```

---

## 🔧 TODO

* [ ] Add pagination
* [ ] Write tests
* [ ] Set up CI/CD




