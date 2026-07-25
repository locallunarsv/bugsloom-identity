[![Go Test](https://github.com/locallunarsv/bugsloom-identity/actions/workflows/test.yml/badge.svg)](...)

# Bugsloom Identity

A lightweight Go identity library providing UUID v7 based identifiers with a type-safe abstraction.

## Overview

`bugsloom-identity` is a standalone identity package designed to provide consistent identifier handling across Go applications.

The goal is to abstract identity management away from the underlying UUID implementation, allowing applications to use a stable and simple API.

## Features

Current features:

- UUID v7 identifier generation
- Type-safe `ID` abstraction
- String representation
- Identifier validation
- Zero-value checking

## Installation

```bash
go get github.com/locallunarsv/bugsloom-identity
```

## Usage

Create a new identity:

```go
package main

import (
	"fmt"

	"github.com/locallunarsv/bugsloom-identity"
)

func main() {
	id := identity.New()

	fmt.Println(id.String())
}
```

Example output:

```text
0192f5e7-8d32-7a91-b5e1-2a3c4d5e6f70
```

## Design Philosophy

`bugsloom-identity` intentionally does not expose the underlying UUID implementation.

Instead of coupling applications directly to a UUID library:

```go
type User struct {
	ID uuid.UUID
}
```

applications can depend on:

```go
type User struct {
	ID identity.ID
}
```

This keeps identity handling consistent and allows the underlying implementation to evolve without affecting consumers.

## Roadmap

Planned features:

- [x] Identifier parsing
- [x] JSON serialization support
- [ ] Database/sql integration
- [ ] Additional validation helpers
- [ ] Stable v1.0 API

## Parsing

Existing identifiers can be parsed back into an `identity.ID`.

Example:

```go
id, err := identity.Parse(
    "0192f5e7-8d32-7a91-b5e1-2a3c4d5e6f70",
)

if err != nil {
    panic(err)
}
```

## JSON Support

`identity.ID` implements `json.Marshaler` and `json.Unmarshaler`.

Example:

```go
type Program struct {
    ID identity.ID `json:"id"`
}
```

JSON Ouput:

```json
{
  "id": "0192f5e7-8d32-7a91-b5e1-2a3c4d5e6f70"
}
```

## Database Support

`identity.ID` implements:

- `database/sql/driver.Valuer`
- `sql.Scanner`

Example:

```go
type Program struct {
    ID identity.ID
}
```

Database:

```sql
CREATE TABLE programs (
    id UUID PRIMARY KEY
);
```

## License

MIT License
