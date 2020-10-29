module github.com/loopcontext/deliver-api-go

go 1.15

// For local dev
// replace github.com/loopcontext/go-graphql-orm v0.0.0-20201029123957-01f382e6092d => ../go-graphql-orm

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/akrylysov/algnhsa v0.12.1
	github.com/cloudevents/sdk-go v1.2.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/gorilla/mux v1.6.2
	github.com/graph-gophers/dataloader v5.0.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/loopcontext/cloudevents-aws-transport v1.0.9
	github.com/loopcontext/go-graphql-orm v0.0.0-20201029123957-01f382e6092d
	github.com/mitchellh/mapstructure v0.0.0-20180203102830-a4e142e9c047
	github.com/rs/cors v1.6.0
	github.com/rs/zerolog v1.20.0
	github.com/urfave/cli v1.22.4
	github.com/vektah/gqlparser/v2 v2.1.0
	gopkg.in/gormigrate.v1 v1.6.0
)
