Run the PDP server (port 8081):
```bash
go run cmd/pdp/main.go
```

Run the Todo app (port 8080):
```bash
go run cmd/server/main.go
```

Send request to the Todo app:
```bash
curl localhost:8080/todos/1 \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJDaVJtWkRBMk1UUmtNeTFqTXpsaExUUTNPREV0WWpkaVpDMDRZamsyWmpWaE5URXdNR1FTQld4dlkyRnMiLCJpYXQiOjE1MTYyMzkwMjJ9.mhVtH7VfV0drkVY_urluMjn2g3nddMk5AMpk-Sa_5q4'
```
