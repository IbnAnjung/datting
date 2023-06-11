# Datting App


# Project Structure
    .
    ├── app                           # bootstrapping 
    ├   ├── bootstap.go               # build all config, driver, entity etc 
    ├   ├── config.go                 # Load apps configuration
    ├   ├── middleware.go             # http middleware
    ├   ├── router.go                 # http route
    ├── database                      # sql file for migration
    ├   ├── sql_dump.sql              # mysql dump file 
    ├── driver                        # contain driver / primary engine 
    ├   ├── http.go                   # build http service 
    ├   ├── mysql.go                  # mysql connectivity driver
    ├   ├── redis.go                  # redis connectivity driver 
    ├── entity                        # contain interface or apps core
    ├   ├── account_entity             
    ├   ├── auth_entity                
    ├   ├── user_entity                
    ├   ├── user_swap_entity           
    ├   ├── util_entity                
    ├── handler                       # http handler
    ├   ├── account_handler                
    ├   ├── auth_handler                
    ├   ├── user_handler                
    ├── repository                    # contain connectivity to data source 
    ├   ├── mysqlgorm                
    ├   ├── redis_repository                
    ├── usecase                       # implement of app core
    ├   ├── account                
    ├   ├── auth                
    ├   ├── user                
    ├   ├── user_swap                
    ├── utils                         # Tools and utilities
    ├   ├── bycrypt.go              
    ├   ├── errors.go              
    ├   ├── helper.go              
    ├   ├── jwt.go              
    ├   ├── orm.go              
    ├   ├── redis_cache.go              
    ├   ├── response.go              
    ├   ├── vlaidator.go              
    ├── .env.example                  
    ├── .gitignore                   
    ├── docker-compose.yaml
    ├── go.mod
    ├── main.go
    ├── mocker.yaml
    └── README.md

# How To run This Project
```bash
# clone repository
git clone git@github.com:IbnAnjung/datting.git

cd datting

cp .env.example .env

#if You'r host already install mysql and redis, just run database/sql_dump.sql on you'r sql client.

#or you can use docker for mysql and redis,
# run this command if you not use your mysql and redis on your host. 
docker-compose up -d 

#adjust you'r .env file

#now you can run this project
go run main.go

```