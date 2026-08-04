# ADR 003 — Webhook events storage

## Options
- PostgreSQL (dedicated instance)
- SQLite
- In-memory (lost on restart)

## Pros/Cons

**PostgreSQL**
+ Production-grade
+ Already used by Forgejo and Authentik
- Adds a third database instance for minimal data
- Overkill for append-only event logs

**SQLite**
+ Zero infrastructure overhead
+ Sufficient for sequential, append-only data
+ Single file, easy to inspect and backup
- Not suited for high write concurrency

**In-memory**
+ Trivial to implement
- Data lost on restart

## Decision
SQLite. Webhook events are append-only and low-volume.
A dedicated PostgreSQL instance would add operational complexity
with no practical benefit at this scale.