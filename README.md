# Kurbankan 🐄

<div align="center">
<img src="./kurbankan.png" alt="app-icon" width="300" style="text-align:center" />
</div>

**Kurbankan** is a REST API platform that simplifies the Qurban process. Users can select a mosque based on their location, make payments via virtual account, and track the entire process transparently — from offering selection to distribution.

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.25 |
| Framework | Gin v1.12 |
| ORM | GORM v1.30 |
| Database | PostgreSQL |
| Authentication | JWT (golang-jwt/jwt v5) |
| Payment | Xendit Virtual Account |
| API Docs | Swagger (swaggo) |
| CORS | gin-contrib/cors |

---

## Project Structure

```
kurbankan/
├── config/         # Database connection
├── controllers/    # HTTP handler layer
├── database/       # SQL migrations & seed files
├── docs/           # Auto-generated Swagger docs
├── middlewares/    # JWT auth middleware
├── models/         # GORM models & request/response types
├── repository/     # Data access layer
├── routes/         # Route registration
└── utils/          # Helpers (JWT, hashing, response, validator, pagination)
```

---

## Getting Started

### Prerequisites

- Go 1.25+
- PostgreSQL
- [swag CLI](https://github.com/swaggo/swag) for regenerating API docs

### Environment Variables

Create a `.env` file in the project root:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=kurbankan
APP_PORT=8080
JWT_SECRET=your_jwt_secret
XENDIT_SECRET_KEY=your_xendit_secret
```

### Run

```bash
go mod tidy
go run main.go
```

### Regenerate Swagger Docs

```bash
swag init --generalInfo main.go --output docs
```

API documentation will be available at:
```
http://localhost:8080/swagger/index.html
```

---

## API Endpoints

All `/api/*` endpoints (except those listed as public) require a `Bearer` token in the `Authorization` header.

### Auth
| Method | Path | Description |
|---|---|---|
| POST | `/auth/register/participant` | Register as participant |
| POST | `/auth/register/mosque` | Register as mosque admin |
| POST | `/auth/login` | Login and get JWT token |

### Area (Public)
| Method | Path | Description |
|---|---|---|
| GET | `/api/area/provinces` | List provinces |
| GET | `/api/area/regencies` | List regencies |
| GET | `/api/area/districts` | List districts |
| GET | `/api/area/villages` | List villages |

### Constants (Public)
| Method | Path | Description |
|---|---|---|
| GET | `/api/constants/banks` | List supported banks |

### Qurban Periods (Auth required)
| Method | Path | Description |
|---|---|---|
| GET | `/api/qurban-periods` | List qurban periods (filterable by year) |
| POST | `/api/qurban-periods` | Create a qurban period |
| PATCH | `/api/qurban-periods/:id` | Update a qurban period |
| DELETE | `/api/qurban-periods/:id` | Delete a qurban period |

### Qurban Offerings (Auth required)
| Method | Path | Description |
|---|---|---|
| GET | `/api/qurban-offerings` | List qurban offerings |
| POST | `/api/qurban-offerings` | Create a qurban offering |
| PATCH | `/api/qurban-offerings/:id` | Update a qurban offering |
| DELETE | `/api/qurban-offerings/:id` | Delete a qurban offering |

### Mosques (Auth required)
| Method | Path | Description |
|---|---|---|
| GET | `/api/mosques` | List mosques |
| GET | `/api/mosques/:id` | Get mosque detail |
| POST | `/api/mosques` | Create a mosque |
| PATCH | `/api/mosques/:id` | Update a mosque |
| DELETE | `/api/mosques/:id` | Delete a mosque |
| POST | `/api/mosques/payment-methods` | Add mosque payment method |

### Participants (Auth required)
| Method | Path | Description |
|---|---|---|
| GET | `/api/participants` | List participants |
| GET | `/api/participants/:id` | Get participant detail |
| PATCH | `/api/participants/:id` | Update participant |
| DELETE | `/api/participants/:id` | Delete participant |

### Beneficiaries (Auth required)
| Method | Path | Description |
|---|---|---|
| GET | `/api/beneficiaries` | List beneficiaries |
| GET | `/api/beneficiaries/:id` | Get beneficiary detail |
| POST | `/api/beneficiaries` | Create a beneficiary |
| PATCH | `/api/beneficiaries/:id` | Update a beneficiary |
| DELETE | `/api/beneficiaries/:id` | Delete a beneficiary |

### Transactions (Auth required)
| Method | Path | Description |
|---|---|---|
| GET | `/api/transactions` | List transactions |
| GET | `/api/transactions/mosque` | List transactions by mosque |
| POST | `/api/transactions` | Create a transaction |
| PUT | `/api/transactions/:id/proof` | Upload payment proof |
| PUT | `/api/transactions/:id/verify` | Verify a transaction |
| POST | `/api/transactions/cancel-expired` | Cancel expired transactions (public) |

### Users (Auth required)
| Method | Path | Description |
|---|---|---|
| GET | `/api/users` | List users |
| PATCH | `/api/users/:id` | Update user |

### Xendit
| Method | Path | Description |
|---|---|---|
| POST | `/api/xendit/va-callback` | Virtual account payment webhook |

---

## API Response Format

**List (paginated)**
```json
{
  "data": [],
  "meta": { "page": 1, "limit": 10, "total": 100, "total_pages": 10 }
}
```

**Single item**
```json
{ "data": {} }
```

**Create / Update / Delete**
```json
{ "message": "Entity created successfully", "data": {} }
```

**Error**
```json
{ "error": { "code": "NOT_FOUND", "message": "Resource not found" } }
```

**Validation error**
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Validation failed",
    "details": { "field": "error message" }
  }
}
```

---

## Roadmap

| Feature | Status |
|---|---|
| Xendit Virtual Account integration | ⏳ In progress |
| Xendit Disbursement | ⏳ In progress |
| GitHub Actions CI/CD | 📎 Planned |
| Dockerfile | 📎 Planned |
| Docker Compose | 📎 Planned |
| Nginx | 📎 Planned |
| Cloudflare | 📎 Planned |
| Observability | 📎 Planned |