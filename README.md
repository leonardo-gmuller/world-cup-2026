# World Cup 2026 Betting Pool API

Backend API for the World Cup 2026 betting pool platform.

This project was built using Go following Clean Architecture principles, separating domain, use cases, gateways, and infrastructure layers.

---

# Goal

The system allows users to:

- create betting pool groups;
- authenticate users;
- submit match predictions;
- automatically calculate scores;
- view rankings by group;
- automatically import and update matches/results from an external football API.

---

# Stack

- Go
- PostgreSQL
- Redis
- SQLC
- Chi Router
- JWT
- Docker
- Football API

---

# Architecture

```txt
cmd/
  api/

internal/
  app/
    domain/
    usecase/
    gateway/
    delivery/
    library/

migrations/