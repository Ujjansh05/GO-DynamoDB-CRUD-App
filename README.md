# GO DynamoDB CRUD App - Postman Testing Guide

This README is focused on testing the API with Postman.

## Status

The project is under active development. This guide documents the intended API test flow so you can validate behavior quickly.

## Prerequisites

- Go installed
- AWS credentials configured (`~/.aws/credentials`) or environment-based AWS auth
- DynamoDB access (AWS or local setup)
- Postman desktop app

## Run The API

1. Open project folder:

```bash
cd GO_DynamoDB_CRUD_App
```

2. If `go.mod` is missing, initialize it:

```bash
go mod init github.com/<your-username>/GO_DynamoDB_CRUD_App
go mod tidy
```

3. Set environment variables (PowerShell example):

```powershell
$env:PORT="8080"
$env:TIMEOUT="30"
$env:AWS_REGION="ap-south-1"
```

4. Start server:

```bash
go run ./cmd
```

Base URL example: `http://localhost:8080`

## Postman Setup

1. Create a Postman Environment named `local`.
2. Add these variables:
   - `baseUrl` = `http://localhost:8080`
   - `productId` = (leave blank initially)
3. Create a collection named `GO DynamoDB CRUD`.
4. Set request body type to `raw` + `JSON` where needed.

## Endpoint List For Testing

- `GET {{baseUrl}}/health`
- `GET {{baseUrl}}/product`
- `POST {{baseUrl}}/product`
- `GET {{baseUrl}}/product/{{productId}}`
- `PUT {{baseUrl}}/product/{{productId}}`
- `DELETE {{baseUrl}}/product/{{productId}}`

## Recommended Postman Test Order

1. Health Check
   - Request: `GET {{baseUrl}}/health`
   - Expected: `200 OK`

2. Create Product
   - Request: `POST {{baseUrl}}/product`
   - Body:

```json
{
  "name": "Postman Test Product"
}
```

   - Expected: `200 OK` with created `id`
   - Add this in the **Tests** tab to save ID:

```javascript
pm.test("Status code is 200", function () {
  pm.response.to.have.status(200);
});

const res = pm.response.json();
if (res.result && res.result.id) {
  pm.environment.set("productId", res.result.id);
}
```

3. Get All Products
   - Request: `GET {{baseUrl}}/product`
   - Expected: `200 OK`

4. Get Product By ID
   - Request: `GET {{baseUrl}}/product/{{productId}}`
   - Expected: `200 OK`

5. Update Product
   - Request: `PUT {{baseUrl}}/product/{{productId}}`
   - Body:

```json
{
  "name": "Postman Updated Product"
}
```

   - Expected: `204 No Content` (or project-specific success status)

6. Delete Product
   - Request: `DELETE {{baseUrl}}/product/{{productId}}`
   - Expected: `204 No Content`

7. Negative Test (Invalid UUID)
   - Request: `GET {{baseUrl}}/product/invalid-id`
   - Expected: `400 Bad Request`

## Optional Collection-Level Tests

Use this in collection scripts to validate JSON response shape:

```javascript
pm.test("Response is JSON", function () {
  pm.response.to.be.json;
});
```

## Common Troubleshooting

- `500 Internal Server Error`:
  - Check AWS credentials and region.
  - Verify DynamoDB table exists and is accessible.
- Timeout issues:
  - Increase `TIMEOUT` env var.
- `productId` not set in Postman:
  - Confirm Create Product response includes `result.id`.
  - Ensure environment is selected in Postman.

## Current Utility Tests

```bash
go test ./...
```

Test folders currently include:

- `utils/env`
- `utils/logger`
