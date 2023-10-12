# Sport Equipment Borrowing

## Project Structure

```plaintext
/sport-equipment-borrowing
    /cmd
        /sport-equipment-borrowing-api
            main.go
    /pkg
        /app
            app.go
            errors.go
            /user
                user.go
                user_repository.go
                user_service.go
            /equipment
                equipment.go
                equipment_repository.go
                equipment_service.go
            /borrowing
                borrowing.go
                borrowing_repository.go
                borrowing_service.go
        /db
            db.go
        /http
            /handler
                user_handler.go
                equipment_handler.go
                borrowing_handler.go
            server.go
        /middleware
            authentication_middleware.go
        /util
            hash.go
    /scripts
        migrate.go
    /configs
        config.yaml
    go.mod
    go.sum
```
## Hexagonal Architecture

### OverviewMake sure to run go mod tidy to update and tidy the project dependencies.

The Hexagonal Architecture, also known as Ports and Adapters, is a software architectural pattern that puts the application's business logic at the core and isolates it from external dependencies such as databases, frameworks, and UIs. This separation enables the application to be more maintainable, testable, and adaptable to changes.

### Components

1. **cmd**: Contains the main entry point for the application.

2. **pkg/app**: Core application logic including errors, user, equipment, borrowing models, repositories, and services.

3. **pkg/db**: Database connection and initialization.

4. **pkg/http**: Handles HTTP-related concerns, including handlers for user, equipment, borrowing operations.

5. **pkg/middleware**: Custom middleware for the application, such as authentication.

6. **pkg/util**: Utility functions, e.g., for hashing passwords.

7. **scripts**: Additional scripts, e.g., for database migration.

8. **configs**: Configuration files, e.g., `config.yaml` for application configuration.

### How to Run

1. Make sure you have Go installed on your machine.

2. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/sport-equipment-borrowing.git
   ```

### Configuration

Update the .env file in the configs directory with your specific configuration settings, such as database connection details.

### Database Migration

To perform database migration, run the migrate.go script in the scripts directory.

   ```bash
   go run scripts/migrate.go
   ```

### Dependencies

Make sure to run go mod tidy to update and tidy the project dependencies.
