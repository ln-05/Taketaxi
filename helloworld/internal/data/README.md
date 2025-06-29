# Data Layer

This package contains the data access layer for the application, including database and Redis connections.

## Database Connection

The application uses GORM with MySQL driver for database operations.

### Configuration

Database configuration is defined in `configs/config.yaml`:

```yaml
data:
  database:
    driver: mysql
    source: "user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
```

### Usage

```go
// Get database connection
db := data.GetDB()

// Example: Create a user
user := &User{
    Name: "John Doe",
    Age:  30,
}
err := db.Create(user).Error

// Example: Query users
var users []User
err := db.Find(&users).Error
```

## Redis Connection

The application uses go-redis for Redis operations.

### Configuration

Redis configuration is defined in `configs/config.yaml`:

```yaml
data:
  redis:
    addr: "localhost:6379"
    password: "your_password"
    read_timeout: 0.2s
    write_timeout: 0.2s
```

### Usage

```go
// Get Redis client
rdb := data.GetRedis()

// Example: Set a key
err := rdb.Set(ctx, "key", "value", 0).Err()

// Example: Get a key
val, err := rdb.Get(ctx, "key").Result()
```

## Data Models

### User Model

```go
type User struct {
    ID   uint   `gorm:"primarykey"`
    Name string `gorm:"size:255;not null"`
    Age  int    `gorm:"not null"`
}
```

## Repository Pattern

The data layer implements the repository pattern to abstract data access:

```go
type GreeterRepo interface {
    Save(context.Context, *Greeter) (*Greeter, error)
    Update(context.Context, *Greeter) (*Greeter, error)
    FindByID(context.Context, int64) (*Greeter, error)
    ListByHello(context.Context, string) ([]*Greeter, error)
    ListAll(context.Context) ([]*Greeter, error)
}
```

## Testing

Run the database connection test:

```bash
go test ./internal/data -v
```

## Dependencies

- `gorm.io/gorm` - ORM for database operations
- `gorm.io/driver/mysql` - MySQL driver for GORM
- `github.com/redis/go-redis/v9` - Redis client
