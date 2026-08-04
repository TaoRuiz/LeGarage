# ADR 001 — HTTP framework

## Options
- Go stdlib `net/http`
- Gin

## Pros/Cons

**stdlib**
+ Zero dependencies
- More verbose routing and middleware
- Manual template handling

**Gin**
+ Cleaner routing and middleware
+ Built-in template rendering
+ Widely used in production Go projects
- Adds an external dependency

## Decision
Gin. The goal is to deliver a functional, readable project — not to demonstrate
framework internals. Gin reflects real-world Go project choices and keeps the
codebase concise.