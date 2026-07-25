# Bugsloom Identity

A lightweight Go identity library providing UUID v7 based identifiers with a type-safe abstraction.

## Overview

`bugsloom-identity` is a standalone identity package designed to provide consistent identifier handling across Go applications.

The goal is to abstract identity management away from the underlying UUID implementation, allowing applications to use a stable and simple API.

## Features

Current features:

* UUID v7 identifier generation
* Type-safe `ID` abstraction
* String representation
* Identifier validation
* Zero-value checking

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

* [ ] Identifier parsing
* [ ] JSON serialization support
* [ ] Database/sql integration
* [ ] Additional validation helpers
* [ ] Stable v1.0 API

## License

MIT License
