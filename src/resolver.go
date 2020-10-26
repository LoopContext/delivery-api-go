package src

import (
	"context"

	"github.com/loopcontext/deliver-api-go/gen"
	"github.com/loopcontext/go-graphql-orm/events"
	//"github.com/loopcontext/checkmail"
)

// New ...
func New(db *gen.DB, ec *gen.EventController) *Resolver {
	resolver := NewResolver(db, ec)

	// resolver.Handlers.CreateUser = func(ctx context.Context, r *gen.GeneratedResolver, input map[string]interface{}) (item *gen.User, err error) {
	// Before save

	// Verify email, for example.
	// err = checkmail.ValidateFormat(item.Email)
	// if err != nil {
	// 	return nil, err
	// }
	// return gen.CreateUserHandler(ctx, r, input)
	// }
	resolver.Handlers.OnEvent = func(ctx context.Context, r *gen.GeneratedResolver, e *events.Event) (err error) {
		// After save
		// if e.Entity == "User" && (e.Type == events.EventTypeCreated || e.Type == events.EventTypeUpdated) {
		// do something...
		// }
		return nil
	}

	return resolver
}

// You can extend QueryResolver for adding custom fields in schema
// // Hello world
// func (r *QueryResolver) Hello(ctx context.Context) (string, error) {
// 	return "world", nil
// }
