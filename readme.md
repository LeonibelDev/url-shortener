# URL Shortener API

![Home Page](./img/image.png)

A high-performance URL shortener service built with Go. This API allows users to generate shortened links and efficiently redirect to the original URL using unique identifiers.

---

## Overview

This project demonstrates the design and implementation of a lightweight URL shortening service, focusing on performance, simplicity, and clean architecture using Go.

---

## Features

- URL shortening with unique identifiers  
- Fast HTTP redirection  
- Clean and minimal API design  
- Scalable structure for future enhancements  

---

## Tech Stack

- Go (Golang)  
- net/http (or your router if applicable)  
- JSON API  

---

## API Endpoints

| Method | Endpoint   | Description |
|--------|-----------|-------------|
| `POST` | `/shorten` | Generate a shortened URL |
| `GET`  | `/{id}`    | Redirect to the original URL |

---

## Installation & Setup

Clone the repository and run the project locally:

```bash
git clone https://github.com/leonibeldev/url-shortener.git
cd url-shortener
go mod tidy
go run main.go
````

The server will start at:

```bash
http://localhost:8080
```


## Usage

### Shorten a URL

Send a `POST` request to `/shorten`:

```json
{
  "link": "https://example.com"
}
```

Response:

```json
{
  "id": "a1b2c3",
  "link": "https://example.com"
}
```

### Redirect

Access the shortened URL:

```http
GET /a1b2c3
```

This will redirect to the original URL.


## Project Structure

```
.
├── main.go
├── handlers/
├── models/
├── storage/
└── utils/
```

## Future Improvements

* Persistent storage (PostgreSQL / Redis)
* Custom aliases for URLs
* Expiration time for links
* Rate limiting and security enhancements
* Analytics (click tracking)


## Author

Leonibel Segura Etx

* GitHub: [https://github.com/LeonibelDev](https://github.com/LeonibelDev)
* Twitter: [https://twitter.com/@SrLeonibel](https://twitter.com/@SrLeonibel)
* Portfolio: [https://mytechblog.dev](https://mytechblog.dev)

