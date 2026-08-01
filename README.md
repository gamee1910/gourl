# Go URL Shortener

A high-performance URL shortener service built with Go (Golang), using the Gin Web Framework and Redis for in-memory caching.

## Features

- Generate short URLs from long original URLs.
- Uses SHA-256 hashing combined with Base58 encoding for concise short links.
- Stores URL mappings in Redis with Time-To-Live (TTL) support.
- Simple and lightweight RESTful API.

## Tech Stack

- Go
- Gin Web Framework
- Redis
- Base58 / SHA-256

