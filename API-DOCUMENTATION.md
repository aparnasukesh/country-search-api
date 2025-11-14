# Country Search API Documentation

## Base URL

```
http://localhost:8000/api
```

---

## **1. Search Countries by Name**

### **Endpoint**

```
GET /countries/search
```

### **Query Parameter**

| Name | Type   | Required | Description                                     |
| ---- | ------ | -------- | ----------------------------------------------- |
| name | string | Yes      | Name (or partial name) of the country to search |

### **Example Request**

```
curl "http://localhost:8000/api/countries/search?name=india"
```

### **Example Successful Response (200 OK)**

```json
[
  {
    "name": "India",
    "capital": "New Delhi",
    "region": "Asia",
    "population": 1380004385,
    "flag": "🇮🇳"
  }
]
```

### **Response Codes**

| Code | Meaning                             |
| ---- | ----------------------------------- |
| 200  | Successful search                   |
| 400  | Missing or invalid query parameters |
| 404  | No countries found                  |
| 500  | Internal server error               |

---

## **Setup Instructions**

### **1. Clone the Repository**

```
git clone <your-repo-url>
cd country-search-api
```

### **2. Install Dependencies**

```
go mod download
```

### **3. Start the Server**

```
go run cmd/main.go
```

### **4. Test the API**

Use the CURL command:

```
curl "http://localhost:8000/api/countries/search?name=india"
```

---

## **Project Structure Overview**

```
/country-search-api
├── cmd
│   └── main.go
├── internal
│   ├── app
│   │   └── country
│   ├── boot
│   ├── cache
│   ├── client
│   ├── di
│   └── utils
├── pkg
├── config
└── go.mod
