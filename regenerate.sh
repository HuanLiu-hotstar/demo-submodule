#!/bin/bash
# in regenerate.sh
## regenerate the graphql code when schema changed
go run github.com/99designs/gqlgen generate
