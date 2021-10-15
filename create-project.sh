#!/bin/bash

## in create-project.sh file
dir=$1
## first mkdir
mkdir -p "$dir" && cd $dir

## init go mod
go mod init $dir

## init gogen
## tools.go will auto generate golang code for graphql
printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt >tools.go
go mod tidy

## init graphql project
go run github.com/99designs/gqlgen init
