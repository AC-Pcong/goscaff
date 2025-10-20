#!/bin/bash

echo "Creating Go Web Scaffold project structure..."

mkdir -p cmd/server
mkdir -p internal/{config,handler,middleware,model,repository,service,router}
mkdir -p pkg/{database,logger,utils,validator}
mkdir -p api/docs
mkdir -p web/{static/{css,js,images},templates/auth}
mkdir -p migrations
mkdir -p scripts
mkdir -p configs
mkdir -p tests/{integration,unit}

touch cmd/server/main.go
touch internal/config/config.go
touch internal/handler/{user_handler.go,auth_handler.go,common_handler.go}
touch internal/middleware/{auth.go,logger.go,recovery.go}
touch internal/model/{user.go,common.go}
touch internal/repository/{user_repository.go,interface.go}
touch internal/service/{user_service.go,auth_service.go}
touch internal/router/router.go
touch pkg/database/database.go
touch pkg/logger/logger.go
touch pkg/utils/{jwt.go,password.go,response.go}
touch pkg/validator/validator.go
touch migrations/001_create_users.sql
touch scripts/{build.sh,deploy.sh,migrate.sh,init.sh}
touch configs/{config.yaml,config.dev.yaml,config.prod.yaml}
touch Makefile
touch README.md
touch Dockerfile
touch docker-compose.yml
touch .gitignore
touch go.mod

echo "Project structure created successfully!"
echo "You can now run: go mod init github.com/yourname/goscaff"
