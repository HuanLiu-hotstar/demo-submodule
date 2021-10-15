# demo-submodule

use git submodule to use the central schema repo

## init project

- use the `create-project.sh` file
- delete the `graph` `schema` directory
- we should use schema from repo https://github.com/hotstar/hs-core-graphql-schema

## update schema

```sh

## add submodule, and download the schema
git submodule add https://github.com/hotstar/hs-core-graphql-schema 


## modify schema location in  gqlgen.yml

## in gqlgen.yml file
# Where are all the schema files located? globs are supported eg  src/**/*.graphqls
schema:
  - hs-core-graphql-schema/schema/pcdds/*.graphql

## regenerate the project using new schema
go run github.com/99designs/gqlgen generate
# or 
bash regenerate.sh

```

- at last the layout of project may like this

```sh
.
├── README.md
├── create-project.sh
├── go.mod
├── go.sum
├── gqlgen.yml
├── graph
│   ├── generated
│   │   └── generated.go
│   ├── model
│   │   └── models_gen.go
│   ├── queries.resolvers.go
│   ├── resolver.go
│   └── schema.resolvers.go
├── hs-core-graphql-schema # submodule
│   ├── README.md
│   ├── appoloStudio
│   │   └── schemaUploader.sh
│   ├── config.json
│   ├── hs-core-graphql-schema.iml
│   └── schema
│       ├── cdds
│       │   ├── queries.graphql
│       │   └── schema.graphql
│       └── pcdds  ## our schema location
│           ├── queries.graphql
│           └── schema.graphql
├── regenerate.sh
├── server.go
└── tools.go
```

## maintain the schema

- we can update schema in hs-core-schema repo, and use branch of that repo.
- we can update schema in local and regenerate the project.
- we should commit the  schema change first, then our GraphQL server code
