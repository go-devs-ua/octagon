# Octagon

## Arch
```
octagon/
├── app/
│   ├── ent/
│   │   └── user.go
│   ├── repo/
│   │   └── pg/
│   │         └── user.sql
│   ├── trans/ 
│   │   └── http/ 
│   │       ├── abs.go 
│   │       ├── handler.go
│   │       ├── resp.go
│   │       ├── router.go
│   │       ├── server.go
│   │       └── user.go
│   └── usecase/
│       ├── abs.go
│       └── user.goo
├── cfg/
│   ├── config.go
│   └── config.yml
├── cmd/
│   ├── cli/
│   │    └── main.go 
│   └── rest/
│       └── main.go 
├── migration/
│   └── migration.sql
├── pkg/
│   └── dummy/ 
│      └── app.go 
├── README.md
└── .env 
```