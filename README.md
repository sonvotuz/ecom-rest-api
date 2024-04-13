# E-Commerce REST API

This e-commerce REST API is built with Go and provides a robust backend for e-commerce applications. The API supports user authentication, product creation, product retrieve, and cart checkout functionality.

## Features

- User registration and login with JWT authentication.
- Product management (create, retrieve).
- Cart checkout process.

## Technology Stack
- Go - The primary programming language.
- PostgreSQL - Relational database management system.
- JWT (JSON Web Tokens) - For authorization and secure communication.
- Golang-Migrate - For database migrations.
- Validator - For data validation.
- Bcrypt - For password hashing.
- PQ - PostgreSQL driver for Go.
- Godotenv - For loading environment variables from .env files.

## How To Run
### Clone the Repository
```bash
git clone git@github.com:vnsonvo/ecom-rest-api.git
cd ecom-rest-api
```

### Set Up Environment Variables
Copy the sample environment configuration file and edit it to suit your needs:

```bash
cp .env.sample .env
```
Edit .env and set your database connection details and JWT secret key.

### Instal Migrate CLI
Refer to this page [migrate-cli](https://github.com/golang-migrate/migrate/tree/v4.17.0/cmd/migrate)


### Run Database Migrations
Ensure your PostgreSQL database is running, and then execute the following command to apply database migrations:

```bash
make migrate-up
```

Build and Run the Application
```bash
make run
```
### Run Tests
Run the command
```bash
make test
```

## Alternatively, use Docker Compose to build and run the application along with PostgreSQL:

```bash
docker-compose up --build
```

### Improvement
- Introduce a feature to manage user addresses. This will enable the storage of addresses for each user, which can then be dynamically used during the checkout process.
- Develop functionality to record and retrieve the history of orders placed by each user. This will allow orders to be stored and displayed within the user's profile area.
- Create an endpoint to cancel orders. This will provide users with the option to cancel their orders.

### Contributing
Contributions are welcome! Please feel free to submit a pull request.
