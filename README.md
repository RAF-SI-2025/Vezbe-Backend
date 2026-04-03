# Vezbe Backend

## Running the server

```bash
go run main.go
```

Server starts on `http://localhost:8080`.

## API Endpoints

### Get all users

```bash
curl http://localhost:8080/users
```

### Get user by ID

```bash
curl http://localhost:8080/users/YOUR-UUID-HERE
```

### Create user

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"first_name":"John","last_name":"Doe","email":"john.doe@example.com","password":"secret123"}'
```

### Update user

```bash
curl -X PUT http://localhost:8080/users/YOUR-UUID-HERE \
  -H "Content-Type: application/json" \
  -d '{"first_name":"Jane","last_name":"Doe","email":"jane.doe@example.com"}'
```

### Delete user

```bash
curl -X DELETE http://localhost:8080/users/YOUR-UUID-HERE
```

### Change password

```bash
curl -X PUT http://localhost:8080/users/YOUR-UUID-HERE/password \
  -H "Content-Type: application/json" \
  -d '{"old_password":"secret123","new_password":"newSecret456"}'
```
