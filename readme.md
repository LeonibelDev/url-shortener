# URL Shortener 🚀  
![Home Page](./img/image.png)

A simple URL shortener service built with Go. It allows users to shorten links and redirect to the original URL using a unique ID.  

## ✨ Features  
✅ Easily shorten URLs  
✅ Fast and efficient redirection  
✅ Built with Go  

## 📌 Endpoints  

| Method | Route     | Description |
|--------|----------|-------------|
| `POST` | `/shorten` | Shortens a given URL |
| `GET`  | `/{id}`   | Redirects to the original URL using the shortened ID |


## ⚡ Installation  

Follow these steps to install and run the project:  

```bash
git clone https://github.com/leonibeldev/url-shortener.git
cd url-shortener
go mod tidy
go run main.go
```


## 🚀 Usage  

### 🔗 Shorten a URL  
Send a `POST` request to `/shorten` with a JSON payload containing the URL to shorten:  

```json
{
  "link": "https://example.com"
}
```

You will receive a response with a unique ID:

```json
{
  "id": "a1b2c3",
  "link": "https://example.com"
}
```

### 🔄 Redirect to the Original URL  
Access the root of the server with the shortened ID to be redirected:  

```http
GET /a1b2c3
```

## 🧑‍💻 Developer  

Developed by **LeonibelDev**  

📌 [Twitter](https://twitter.com/@SrLeonibel)  
📌 [Portfolio](https://leonibel.herokuapp.com)  
