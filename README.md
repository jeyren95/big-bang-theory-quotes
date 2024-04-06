# Big Bang Theory Quotes
API that returns quotes from the popular sitcom show - Big Bang Theory

![Alt Text](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExeXZtb25samd5OTR4dWJ4OHd5dHMxOGNrcW84djRuM3piOWNwYWlnMSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/VmM6OYEREouic/giphy.gif)

# Current features
Users may retrieve quotes from this API based on the following:
1. Character
2. Season
3. Episode

# Getting started
## Clone repository and download necessary dependencies
``` bash
git clone https://github.com/jeyren95/big-bang-theory-quotes
cd big-bang-theory-quotes
cp .example.env .env
go mod download
```

## Install goose
`pressly/goose` needed for database migration. You can install goose by running this following command:
``` bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Run the project
Run the project using a build
``` bash
make migrate.up
make run
```

Or, use go run
``` bash
make migate.up
go run main.go
```

# Deployment
This API is currently a WIP, and is only available on localhost

# Tech stack
- Golang
- Sqlite3

