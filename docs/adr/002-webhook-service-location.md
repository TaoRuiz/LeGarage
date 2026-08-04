# ADR 002 — Webhook service location

## Options
- Dedicated microservice
- Integrated in the portal

## Pros/Cons

**Dedicated service**
+ Single responsibility
+ Independent scalability
- Overkill for current scope
- Adds operational complexity (5th container)

**Integrated in portal**
+ Simpler deployment
+ Sufficient for current scale
- Mixes concerns

## Decision
Integrated in the portal, in a clearly scoped `internal/webhook` package.
Extracting to a dedicated service remains straightforward if the need arises.