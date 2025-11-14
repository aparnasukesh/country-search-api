# 🌍 Country Search API — Clean Architecture (Gin + Custom Cache + Tests)

A clean and modular **Country Search API** implemented using:

- **Gin** (HTTP Router)
- **Custom In-Memory Cache**
- **Clean Architecture Folder Structure**
- **Dependency Injection**
- **Timeout Handling**
- **Unit Tests for Handler, Service & Cache**

---

## 📁 Project Structure

country-search-api/
│
├── cmd/
│   └── main.go                 # Application entry point
│
├── internal/
│   ├── app/
│   │   └── country/            # Handler, Service, Models, Tests
│   │
│   ├── cache/                  # In-memory cache implementation + tests
│   │
│   ├── client/                 # External API client (if added later)
│   │
│   ├── boot/                   # App initialization (router, configs)
│   │
│   ├── di/                     # Dependency injection wiring
│   │
│   └── utils/                  # Helper functions
│
├── pkg/                        # Reusable packages (if needed later)
│
├── go.mod
└── go.sum

---

## 🚀 Features

✔ Search country by name  
✔ In-memory cache using **RWMutex**  
✔ Clean layered separation  
✔ Fast JSON responses  
✔ 5-second timeout using **context**  
✔ Fully unit-testable — mocks included  
✔ Real API-like behavior  

---

## ▶️ Run the Project

### **Install dependencies**

```bash
go mod tidy

Start the server
go run ./cmd


Server runs at:

http://localhost:8080

🌐 API Endpoint
GET /api/country/search?name=India
Sample Success Response (200)
{
  "name": "India",
  "capital": "New Delhi",
  "currency": "₹",
  "population": 1380004385
}

Error Responses
Code	Message
400	missing 'name' parameter
500	internal server error

Internal Code Overview
1️⃣ Handler — internal/app/country/handler.go


Reads name query param


Validates missing parameter


Creates 5-second timeout


Calls the service


Returns JSON



2️⃣ Service — internal/app/country/service.go


Checks in-memory cache


Calls the REST client


Extracts capital, currency, population


Saves result in cache



3️⃣ Cache — internal/cache/cache.go
Thread-safe implementation with:
sync.RWMutex

RWMutex allows:


Multiple readers (Get)


One writer (Set)



🧪 Running Tests
Run all tests
go test ./...

Run with coverage
go test ./... -cover


🧪 Tests Included
✔ Cache Tests
File: internal/cache/cache_test.go
✔ Service Tests
File: internal/app/country/service_test.go
✔ Handler Tests
Using Gin + httptest:
internal/app/country/handler_test.go
Test Cases Included
TestDescriptionMissing nameAPI returns 400Success responseAPI returns 200 + JSONService errorAPI returns 500

🧩 Why This Architecture?
| Folder        | Responsibility                               |
| ------------- | -------------------------------------------- |
| **internal/** | Prevents external imports (Go best practice) |
| **cmd/**      | Application startup logic                    |
| **app/**      | Feature modules – handlers, services, models |
| **boot/**     | Router, config, setup                        |
| **cache/**    | Shared utilities like thread-safe cache      |
| **di/**       | Dependency injection wiring                  |
| **pkg/**      | Exportable packages (if needed)              |

This structure mirrors production-grade Go microservices.