
###https://github.com/go-gorm/gen/blob/master/README.ZH_CN.md
目录结构
   demo
   ├── cmd
   │   └── generate
   │       └── generate.go # execute it will generate codes
   ├── dal
   │   ├── dal.go # create connections with database server here
   │   ├── model
   │   │   ├── method.go # DIY method interfaces
   │   │   └── model.go  # store struct which corresponding to the database table
   │   └── query  # generated code's directory
   |       ├── user.gen.go # generated code for user
   │       └── gen.go # generated code
   |       └── user.gen_test.go # generated unit test
   ├── biz
   │   └── query.go # call function in dal/gorm_generated.go and query databases
   ├── config
   │   └── config.go # DSN for database server
   ├── generate.sh # a shell to execute cmd/generate
   ├── go.mod
   ├── go.sum
   └── main.go