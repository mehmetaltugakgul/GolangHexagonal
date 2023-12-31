# Golang Fiber REST API with Hexagonal Architecture

This is a REST API project developed in Golang, leveraging the Fiber framework for high-performance server-side operations. It's designed to be database agnostic, offering support for PostgreSQL, MySQL, and SQLite, and integrates GORM for object-relational mapping. A key feature is its use of JWT (JSON Web Token) for secure authentication.

## Features

- **RESTful API Design:** Streamlined endpoints for efficient data handling.
- **JWT-Based Authentication:** Ensures secure communication and user verification.
- **GORM ORM Integration:** Simplifies database operations with an object-relational mapper.
- **Database Compatibility:** Works seamlessly with PostgreSQL, MySQL, and SQLite.
- **CRUD Operations:** Complete Create, Read, Update, and Delete capabilities.

## Installation

### Cloning the Repository

To get started with this project, clone the repository and navigate to the project directory:

1. **Clone the repository:**

  
     ```bash
     git clone https://github.com/mehmetaltugakgul/GolangHexagonal.git
     cd your-repo-name
     go mod tidy
   

2. **Edit config file to connect DB**
    ```bash
    {
        "db_user": "yourusername",
        "db_pass": "yourpassword",
        "db_host": "localhost",
        "db_port": "3306",
        "db_name": "yourdbname"
    }

3. **Run the app**
     ```bash
     go run cmd/main.go

## API Usage
This section will provide detailed information about the API endpoints and how to use them effectively.

