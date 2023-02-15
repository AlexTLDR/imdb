# imdb
Simple GraphQL Golang Application to query movies and actors

For this project, I will be using gqlgen as I want to use a Schema first approach.

notes:

To solve these errors when running the "go run github.com/99designs/gqlgen generate" command, run first "go get github.com/99designs/gqlgen" .

../../go/pkg/mod/github.com/99designs/gqlgen@v0.17.22/main.go:15:2: missing go.sum entry for module providing package github.com/urfave/cli/v2 (imported by github.com/99designs/gqlgen); to add:
        go get github.com/99designs/gqlgen@v0.17.22
../../go/pkg/mod/github.com/99designs/gqlgen@v0.17.22/codegen/field.go:15:2: missing go.sum entry for module providing package golang.org/x/text/cases (imported by github.com/99designs/gqlgen/codegen); to add:
        go get github.com/99designs/gqlgen/codegen@v0.17.22
../../go/pkg/mod/github.com/99designs/gqlgen@v0.17.22/codegen/field.go:16:2: missing go.sum entry for module providing package golang.org/x/text/language (imported by github.com/99designs/gqlgen/codegen); to add:
        go get github.com/99designs/gqlgen/codegen@v0.17.22
../../go/pkg/mod/github.com/99designs/gqlgen@v0.17.22/internal/imports/prune.go:15:2: missing go.sum entry for module providing package golang.org/x/tools/go/ast/astutil (imported by github.com/99designs/gqlgen/internal/imports); to add:
        go get github.com/99designs/gqlgen/internal/imports@v0.17.22
../../go/pkg/mod/github.com/99designs/gqlgen@v0.17.22/internal/code/packages.go:11:2: missing go.sum entry for module providing package golang.org/x/tools/go/packages (imported by github.com/99designs/gqlgen/codegen/config); to add:
        go get github.com/99designs/gqlgen/codegen/config@v0.17.22
../../go/pkg/mod/github.com/99designs/gqlgen@v0.17.22/internal/imports/prune.go:16:2: missing go.sum entry for module providing package golang.org/x/tools/imports (imported by github.com/99designs/gqlgen/internal/imports); to add: