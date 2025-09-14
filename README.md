
# Go API Gateway

A simple API Gateway in Go that routes requests to separate microservices for products and users. This project demonstrates basic reverse proxying and service composition using Go's standard library and the Gorilla Mux router.

## Project Structure

```
go-api-gateway/
├── cmd/
│   └── gateway/
│       └── main.go         # API Gateway implementation (port 8082)
├── products/
│   └── main.go             # Products service (port 8081)
├── users/
│   └── main.go             # Users service (port 8080)
├── go.mod
├── go.sum
└── README.md
```

## Services Overview

- **API Gateway** (`cmd/gateway/main.go`):
	- Listens on port **8082**
	- Proxies `/products/*` requests to the Products service (http://localhost:8081)
	- Proxies `/users/*` requests to the Users service (http://localhost:8080)
	- Root endpoint `/` returns a welcome message

- **Products Service** (`products/main.go`):
	- Listens on port **8081**
	- Exposes `/products/healthcheck` endpoint

- **Users Service** (`users/main.go`):
	- Listens on port **8080**
	- Exposes `/users/healthcheck` endpoint

## How to Run

Open three terminal windows/tabs and run each service separately:

```sh
# Terminal 1: Start the users service
cd users
go run main.go

# Terminal 2: Start the products service
cd products
go run main.go

# Terminal 3: Start the API gateway
cd cmd/gateway
go run main.go
```

## Example Endpoints

- Gateway root: [http://localhost:8082/](http://localhost:8082/)
- Products healthcheck: [http://localhost:8082/products/healthcheck](http://localhost:8082/products/healthcheck)
- Users healthcheck: [http://localhost:8082/users/healthcheck](http://localhost:8082/users/healthcheck)

## Dependencies

- [github.com/gorilla/mux](https://github.com/gorilla/mux)

## License

MIT License
