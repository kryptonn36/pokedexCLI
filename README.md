# Pokedex CLI

A small REPL-style Pokedex client that talks to the public PokeAPI and keeps a short-lived cache to reduce repeat requests.

## Features

- Browse location areas with `map` and `mapb`.
- Explore an area to list Pokemon found there.
- Catch Pokemon and keep a local Pokedex for the session.
- Inspect caught Pokemon stats and types.

## Requirements

- Go 1.24.4 (see `go.mod`)

## Install

```bash
go build -o pokedex
```

## Run

```bash
./pokedex
```

Or run directly:

```bash
go run .
```

## Commands

- `help` - Show help output.
- `exit` - Exit the REPL.
- `map` - List the next page of location areas.
- `mapb` - List the previous page of location areas.
- `explore <location_name>` - List Pokemon found in a location area.
- `catch <pokemon_name>` - Attempt to catch a Pokemon.
- `inspect <pokemon_name>` - Show details for a caught Pokemon.
- `pokedex` - List all caught Pokemon in this session.

## Notes

- Caught Pokemon are stored in memory only (no persistence).
- The client uses a small in-memory cache to reduce repeat API calls.

## Tests

```bash
go test ./...
```
