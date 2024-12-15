# go-fullstack

# Golang FullStack Project

This project is a simple golang web application that allows users to add users. The project uses the following technologies:

- Golang
- SQLite
- HTML
- CSS
- JavaScript


## Getting Started

To run the project, follow these steps:

1. Clone the repository using the following command:
```bash
git clone https://github.com/LamichhaneBibek/golang-fullstack.git

```
2. Navigate to the project directory:
```bash
cd golang-fullstack
```

3. Install the project dependencies:
```bash
	go mod tidy
```
4. Create a new SQLite database:
```bash
	goose -dir=assets/migrations create users sql
	goose -dir=assets/migrations sqlite3 database.db up
```

5. Run the project:
```bash
go run cmd/web/main.go
```