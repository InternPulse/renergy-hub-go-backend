#!/bin/bash

git reset --hard
git pull origin main
go get
go build -o renergyhubgo
~/go/bin/swag init
pm2 restart renergyhubgo