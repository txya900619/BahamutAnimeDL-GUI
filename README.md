# BahamutAnimeDL-GUI
the Bahamut Anime Downloader GUI version, build by go package lorca and vue.js
## the Frontend is from [this repo](https://github.com/txya900619/BahamutAnimeDL-frontend)
## How to build this on my computer?
1. `go generate`
2. `go build`
## I don't want to build~
1. `go generate`
2. `go run main.go`
## I want to try my UI, what should I do
1. `use node to serve your frpntend`
2. change main_dev.go's 96 line 127.0.0.1:8080 to your port
3. you can `go run main_dev` or `go build -tags dev`
