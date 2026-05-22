# search-service

[![Golang CI/CD](https://github.com/ermasavior/parkirpintar-search/actions/workflows/cicd.yml/badge.svg)](https://github.com/ermasavior/parkirpintar-search/actions/workflows/cicd.yml)

Read-only service that queries real-time parking spot availability directly from PostgreSQL.

## Responsibilities

- `GetAvailability` — returns total available spots and per-floor breakdown, filtered by vehicle type
- `ListSpots` — returns all spots on a given floor with their current status (`AVAILABLE` / `LOCKED`)

## gRPC API

```
service SearchService {
  rpc GetAvailability (GetAvailabilityRequest) returns (GetAvailabilityResponse);
  rpc ListSpots       (ListSpotsRequest)       returns (ListSpotsResponse);
}
```

Proto: [`proto/search/v1/search.proto`](proto/search/v1/search.proto)

## Dependencies

| Dependency | Purpose |
|---|---|
| PostgreSQL | Read spot inventory |

## Configuration

```bash
cp .env.example .env
```

Key variables: `POSTGRES_DSN`

## Development

```bash
make run              # run locally
make build            # compile binary → bin/search
make test             # all tests
make test-unit        # unit tests only
make unit-test-coverage
make proto            # regenerate gRPC code from .proto
make mock             # regenerate mocks
```
