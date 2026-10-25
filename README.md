# Ghost-Snow ❄️

**A Redis-inspired in-memory data store built from scratch in Go.**

Ghost-Snow is a learning-focused project aimed at understanding how Redis works internally, from parsing the RESP (Redis Serialization Protocol) to building the foundations of a lightweight, high-performance data store.

## ✨ Features

- **RESP Protocol Decoder** — Parses Redis protocol data into Go values.
- **Simple Strings** — Supports RESP simple string parsing.
- **Error Responses** — Handles RESP error values.
- **Integer Parsing** — Decodes integer values.
- **Bulk Strings** — Extracts length-prefixed strings.
- **Array Parsing** — Decodes arrays, including support for nested values through recursive decoding.
- **Modular Codebase** — Separates protocol parsing logic into a dedicated `core` package.

> **Note:** This project is under active development. Features and protocol support may change as development progresses.

## 🛠️ Tech Stack

- **Language:** Go
- **Protocol:** RESP (Redis Serialization Protocol)
- **Standard Library:** Go standard library

## 📁 Project Structure

```text
Ghost-Snow/
├── core/
│   └── resp.go
├── go.mod
└── README.md
```

The repository structure may evolve as new components are added.

## 🚀 Getting Started

### Prerequisites

- Go installed on your system
- Git

Check your Go installation:

```bash
go version
```

### Clone the Repository

```bash
git clone https://github.com/krishna231122-del/Ghost-Snow.git
cd Ghost-Snow
```

### Run the Project

If the repository contains an executable entry point (`package main` with a `main()` function), run:

```bash
go run .
```

The RESP parser in `core/resp.go` is a library package and cannot be executed directly with `go run` unless it is changed into an executable program.

## 🧠 How RESP Decoding Works

Ghost-Snow uses a decoder to inspect the first byte of an incoming RESP value and dispatch it to the appropriate parsing function.

| Prefix | RESP Type | Example |
|---|---|---|
| `+` | Simple String | `+OK\r\n` |
| `-` | Error | `-ERR unknown command\r\n` |
| `:` | Integer | `:123\r\n` |
| `$` | Bulk String | `$5\r\nhello\r\n` |
| `*` | Array | `*2\r\n$3\r\ncat\r\n$3\r\ndog\r\n` |

The decoder returns the parsed value, the number of bytes consumed, and an error where applicable.

### Example

A RESP bulk string:

```text
$5\r\nhello\r\n
```

is decoded into:

```go
"hello"
```

The consumed-byte count allows the decoder to locate the next value in a stream of RESP data.

## 🎯 Project Goals

- Build a deeper understanding of Redis internals.
- Implement reliable RESP parsing.
- Develop an in-memory data storage engine.
- Explore command processing and client-server communication.
- Learn about efficient data structures, concurrency, and networking in Go.

## 🗺️ Roadmap

- [x] Implement foundational RESP decoding functions
- [x] Add decoding for strings, integers, bulk strings, and arrays
- [ ] Improve malformed-input validation and error handling
- [ ] Add comprehensive unit tests
- [ ] Implement an in-memory key-value store
- [ ] Add Redis-style commands such as `SET`, `GET`, and `DEL`
- [ ] Build a TCP server for client connections
- [ ] Explore expiration and TTL support
- [ ] Implement a sliding-window rate limiter
- [ ] Add benchmarks and performance comparisons

*Roadmap items should be updated as each feature
