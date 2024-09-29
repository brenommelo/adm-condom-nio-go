
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/golang-jwt/jwt/v4
go get -u github.com/joho/godotenv
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon

go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
swag init

go install github.com/pressly/goose/v3/cmd/goose@latest

ESTRUTURA
    config
        logger (goportunit)
    docs
        swagger (goportunit)

    initializers ((jwt-auth-go))
    middleware
        valid auth jwt (jwt-auth-go)

    database
        database confi 
        migrations (rssagg)


    models
        1 para todas as entidades

    .env (jwt-auth-go)

*criar objeto para request e um outro diferente para models como no gooportunities

    project/
├── cmd/
│   └── app/
│       ├── main.go
├── docs/
│   └── swagger.go
├── pkg/
│   ├── utils/
│   │   └── util.go (write json, parse json)
│   └── auth/
│       └── auth.go (gerar e validar toke)
├── internal/
│   ├── config/
│   │   └── config.go (load envs)
│   ├── handlers/
│   │   └── handler.go (para coisas gerais como health check, response e request do gooportuniti)
│   ├── services/
│   │    └── service.go (logica de negocio, database)
│   ├── models/
│   │   └── user.go
│   ├── routes/
│   │   └── router.go
│   ├── controllers/
│   │   └── user_controller.go (login hadler, create user, get user)
│   └── middlewares/
│       └── auth_middleware.go
├── migrations/
│   └── 0001_init.up.sql
├── go.mod
├── Makefile
├── .env
└── README.md