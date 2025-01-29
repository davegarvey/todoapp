# Todo App

## Start the Servers

### Run the PDP Server
```bash
go run cmd/pdp/main.go
```

The PDP server runs on port 8081

### Run the Todo App 
```bash
go run cmd/server/main.go
```

The Todo app runs on port 8080

## Example JWT Tokens

Use these JWT tokens to represent the users:

- **Rick Sanchez:**
  ```
  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDaVJtWkRBMk1UUmtNeTFqTXpsaExUUTNPREV0WWpkaVpDMDRZamsyWmpWaE5URXdNR1FTQld4dlkyRnMiLCJpYXQiOjE1MTYyMzkwMjJ9.mhVtH7VfV0drkVY_urluMjn2g3nddMk5AMpk-Sa_5q4
  ```

- **Beth Smith:**
  ```
  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDaVJtWkRNMk1UUmtNeTFqTXpsaExUUTNPREV0WWpkaVpDMDRZamsyWmpWaE5URXdNR1FTQld4dlkyRnMiLCJpYXQiOjE1MTYyMzkwMjJ9.dcMCblRLhfG2L0L_hYW4Gwu_arIe7upuyx0230ayHPs
  ```

- **Morty Smith:**
  ```
  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDaVJtWkRFMk1UUmtNeTFqTXpsaExUUTNPREV0WWpkaVpDMDRZamsyWmpWaE5URXdNR1FTQld4dlkyRnMiLCJpYXQiOjE1MTYyMzkwMjJ9._-FRr7dJEyAiNGkHsG6QuxoU6ZVwnh1Zi8rr_90lfWY
  ```

- **Summer Smith:**
  ```
  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDaVJtWkRJMk1UUmtNeTFqTXpsaExUUTNPREV0WWpkaVpDMDRZamsyWmpWaE5URXdNR1FTQld4dlkyRnMiLCJpYXQiOjE1MTYyMzkwMjJ9.ivOcmvV2AvWtJTqlMlwFbiKE2jHvFcB2tjAVSITN-Uc
  ```

- **Jerry Smith:**
  ```
  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDaVJtWkRRMk1UUmtNeTFqTXpsaExUUTNPREV0WWpkaVpDMDRZamsyWmpWaE5URXdNR1FTQld4dlkyRnMiLCJpYXQiOjE1MTYyMzkwMjJ9.O9O5hOwLn4iu41fyEbQKsSxKn8cpBaaOrRnIA_QMYlA
  ```

## Example Requests

### Create a Todo
```bash
curl -X POST http://localhost:8080/todos \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your-jwt-token>" \
-d '{
  "title": "Sample Todo",
  "completed": false,
  "ownerID": "user-123",
  "ownerName": "John Doe",
  "ownerEmail": "john.doe@example.com",
  "ownerPicture": "https://example.com/avatar.jpg"
}'
```

### List Todos
```bash
curl http://localhost:8080/todos \
-H "Authorization: Bearer <your-jwt-token>"
```

### Get a User
```bash
curl http://localhost:8080/users/{userID} \
-H "Authorization: Bearer <your-jwt-token>"
```

### Update a Todo
```bash
curl -X PUT http://localhost:8080/todos/{id} \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your-jwt-token>" \
-d '{
  "title": "Updated Todo",
  "completed": true,
  "ownerID": "user-123",
  "ownerName": "John Doe",
  "ownerEmail": "john.doe@example.com",
  "ownerPicture": "https://example.com/avatar.jpg"
}'
```

### Delete a Todo
```bash
curl -X DELETE http://localhost:8080/todos/{id} \
-H "Authorization: Bearer <your-jwt-token>"
```

Replace `<your-jwt-token>` with a valid JWT token and `{userID}` and `{id}` with the appropriate values.