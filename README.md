### Technology use:
 - Golang
 - GoRM
 - Gin
 - MySQL
 - Docker

### Step run:
- Install Golang, GoRM, Gin and orther packages
- Run docker-compose up -d
- Change environment in env file (not require)
- Run code: go run main.go
- Setup hot reload code (not require): 
  + Install cosmtrek/air: go install github.com/cosmtrek/air@latest
  + Init .air.toml file: air init
  + Command run code: air
- Build code: go build main.go
