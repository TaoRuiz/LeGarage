# ADR 004 — Frontend architecture

## Options
- JS vanilla + fetch (no framework)
- Vue via CDN (no build step)
- Vue + Vite (separate frontend project)

## Pros/Cons

**JS vanilla**
+ Zero dependencies
+ No build step
- Becomes unmaintainable quickly (no components, no reactivity)

**Vue via CDN**
+ No build step
+ Simple setup
- No SFC, templates as strings
- Manual dependency management
- No tree-shaking

**Vue + Vite**
+ Component-based, reusable across views
+ Scales cleanly as the portal grows
+ Proper build pipeline (path filters in CI/CD)
- Adds a build step

## Decision
Vue + Vite. The portal is a living tool, new views and components will be added
over time. A component-based architecture prevents duplication and keeps the codebase
maintainable without rewiring as requirements evolve.