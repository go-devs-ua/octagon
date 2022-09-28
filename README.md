# Octagon

## Arch
```
octagon/
├── app/
│   ├── entities/
│   │   └── user.go
│   ├── repository/
│   │   └── pg/
│   │         └── user.sql
│   ├── transport/ 
│   │   └── http/ 
│   │       ├── contracts.go 
│   │       ├── handler.go
│   │       ├── resp.go
│   │       ├── router.go
│   │       ├── server.go
│   │       └── user.go
│   └── usecase/
│       ├── contracts.go
│       └── user.goo
├── cfg/
│   ├── config.go
│   └── config.yml
├── cmd/
│   └── rest/
│       └── main.go 
├── migration/
│   └── migration.sql
├── README.md
└── .env 
```