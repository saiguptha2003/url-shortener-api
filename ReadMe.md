# URL Shortener API

This project is a simple URL Shortener API built using the Go programming language and the Gin framework. It allows users to shorten URLs and retrieve the original URL using the shortened link.

---

## Features

- **Shorten URL**: Convert long URLs into shorter, easily shareable links.
- **Redirect URL**: Redirect users to the original URL when accessing a shortened link.
- **Top Domains**: Get a list of top domains used in the shortened URLs.

---

## Tech Stack

- **Language**: Go (Golang)
- **Framework**: Gin
- **Database**: In-memory map (for demo purposes) Maps
- **Testing**: `httptest` and `stretchr/testify`

---

## Prerequisites

Before running this application, ensure you have:

1. Go (1.20 or later) installed.
2. Docker (for containerized deployment).
3. A text editor or IDE for development (e.g., VS Code, Goland).

---

##
## Installation and Setup

### **1. Clone the Repository**
```bash
git clone https://github.com/saiguptha2003/url-shortener-api.git
cd url-shortener
```
---
### 2. Installation

#### Method 1 Direct Execution 

```bash 
go mod init url-shortener-api
go mod tidy
go run main.go
```

#### Method 2 Docker Execution

```bash
docker build -t url-shortener-api .
docker run -p 5000:5000 url-shortener-api
```

#### Method 3 Docker-compose Execution
```bash
docker-compose up --build
```
---

## API ENDPOINTS

#### 1. Shorten URL
#### Method: POST
#### URL: /shorten

#### Request
```json
{
  "url": "https://example.com"
}
```

#### Response
```json
{
    "short_url": "hisXMP"
}
```

#### 2. Redirect URL
#### Method: GET
#### URL: /:shortURL
#### Behavior: Redirects to the original URL.



#### 3. Top Domains
#### Method: GET
#### URL: /top-domains

#### Response 

```json
{
    "top_domains": [
        {
            "Domain": "github.com",
            "Count": 2
        },
        {
            "Domain": "example.com",
            "Count": 1
        }
    ]
}

```

## Additional Information
### Applciation is Exposed at 5000 port in localhost
---
## Tests

```bash
go test -v
```
---
### Contact
##### Author: V. D. Panduranga Sai Guptha
##### Email: saiguptha_v@srmap.edu.in
##### GitHub: github.com/saiguptha2003
##### LinkedIn: linkedin.com/in/saiguptha2003



```javascript
This `README.md` provides an overview of the project, setup instructions, API details, and other essential information. You can customize it further as needed.
```
