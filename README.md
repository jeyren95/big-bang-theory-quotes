# Big Bang Theory Quotes
API that returns quotes from the popular sitcom show - Big Bang Theory

![Alt Text](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExeXZtb25samd5OTR4dWJ4OHd5dHMxOGNrcW84djRuM3piOWNwYWlnMSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/VmM6OYEREouic/giphy.gif)

# Current features
Users may retrieve quotes from this API based on the following:
1. Character
2. Season
3. Episode

# Getting started

### Clone the repository
You can clone this repository by this following command:
``` bash
git clone https://github.com/jeyren95/big-bang-theory-quotes
cd big-bang-theory-quotes
```

### Initialize module
Go project modules file is `go.mod`. You can get the Go project modules by running this following command:
``` bash
go mod tidy
```

### Install goose (if not installed)
Goose needed for database migration. You can install goose by running this following command:
``` bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Start build and migrating process
Build project and migrate needed table for the database by running this following command:
``` bash
make init
```

### Run!
Finally, run the project by running this following command:

``` bash
make run
```
This project runs in `0.0.0.0:8080` by default. You can change it by overriding `IPADDR` variable in `.env` file.

# Deployment
This API is currently a WIP, and is only available on localhost

# Tech stack
- Golang
- Sqlite3

