# AuthZEN PDP & Todo App

This repository contains two applications that demonstrate the functionality of the AuthZEN Policy Decision Point (PDP) alongside a simple Todo app. The Todo app utilises the AuthZEN PDP to enforce fine-grained authorisation rules.

## Overview

This repository is designed to be used alongside an AuthZEN Policy Enforcement Point (PEP) service, such as an API gateway, which integrates with the AuthZEN PDP to enforce medium-grained authorisation.

For more details, refer to [this document](https://hackmd.io/ecYxP6uxSCm5X0RexkAM2g?view), which explains how the AuthZEN PDP and PEP work with the Todo app and an API gateway.

## Setup

### Running the PDP Server
To start the PDP server, run:
```bash
go run cmd/pdp/main.go
```
The PDP server will start on port **9081**.

### Running the Todo App
To start the Todo app, run:
```bash
go run cmd/server/main.go
```
The Todo app will start on port **9080**.

## Example JWT Tokens

Use the following JWT tokens to simulate different users. These tokens can be used in API requests for authorisation.

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

## Example API Requests

### Create a Todo
```bash
curl -X POST http://localhost:9080/todos \
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
curl http://localhost:9080/todos \
-H "Authorization: Bearer <your-jwt-token>"
```

### Retrieve a User
```bash
curl http://localhost:9080/users/{userID} \
-H "Authorization: Bearer <your-jwt-token>"
```

### Update a Todo
```bash
curl -X PUT http://localhost:9080/todos/{id} \
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
curl -X DELETE http://localhost:9080/todos/{id} \
-H "Authorization: Bearer <your-jwt-token>"
```

Replace `<your-jwt-token>` with a valid JWT token, `{userID}` with the relevant user ID, and `{id}` with the corresponding Todo item ID.
