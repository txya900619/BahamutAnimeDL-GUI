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
