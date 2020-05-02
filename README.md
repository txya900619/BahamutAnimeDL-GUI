# BahamutAnimeDL-GUI
the Bahamut Anime Downloader GUI version, build by go package lorca and vue.js
## the Frontend is from [this repo](https://github.com/txya900619/BahamutAnimeDL-frontend)
## How to build this on my computer?
1. `go get github.com/markbates/pkger/cmd/pkger`
2. `pkger`
3. `go build`
## I don't want to build~
1. `go run main.go`
## I want to try my UI, what should I do
1. `use node to serve your frontend`
2. change main_dev.go :8080 to your npm serve port
3. you can `go run main_dev.go` or `go build -tags dev`
## Pakages I use in this project
- [lorca](https://github.com/gocolly/colly), a pakage let you can build desktop application in go, base on chrome
- [colly](https://github.com/gocolly/colly), a pakage for crawler
- [pkger](https://github.com/markbates/pkger), a pakege can embed static files to binary 
- [go-sqlite3](https://github.com/mattn/go-sqlite3), a sqlite3 dirver for go
- [sql-migrate](https://github.com/rubenv/sql-migrate), database migrate tool in go
- [sqlboiler](https://github.com/volatiletech/sqlboiler), a datbase-first orm in go
