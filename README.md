# Calculator – Go

A simple calculator REST API built with Go (standard library only).

## Requirements

- Go 1.21+

## Start

```bash
go run .
```

The API is available at `http://localhost:8080`.

## Endpoints

### POST /api/calculate

Calculate a result.

**Request body:**
```json
{
  "a": 10,
  "b": 2,
  "operation": "+"
}
```

**Supported operations:** `+`, `-`, `*`, `/`

**Response:**
```json
{
  "result": 12,
  "expression": "10 + 2 = 12"
}
```

## Run tests

```bash
go test ./...
```
