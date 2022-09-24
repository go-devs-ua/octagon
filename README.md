# Octagon

## Arch
```
octagon/
├── api/
│   └── openapi.yml
├── cfg/
│   └── config.go
├── cmd/
│   └── srv/
│       └── main.go 
├── core/
│   ├── error.go 
│   ├── user.go 
│   └── validation.go
├── services/
│   ├─ agent/
│   │   ├── agent.go
│   │   └── user.go
│   ├── mov/
│   │   ├── mover.go
│   │   ├── response.go
│   │   └── user.go
│   ├── repo/
│   │   ├── sql/
│   │   │   └── migrations.sql
│   │   ├── repo.go
│   │   ├── migrate.go
│   │   └── user.go
├── README.md
└── .env 
```