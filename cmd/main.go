package main

import (
	"context"
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/app/usecase"
	"github.com/jpmoraess/service-scheduling/internal/infra/handlers"
	"github.com/jpmoraess/service-scheduling/internal/infra/middleware"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8080", "the listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(configs.DATABASE_URI))
	if err != nil {
		panic(err)
	}

	var (
		// repositores initialization
		accountRepository       = persistence.NewAccountMongoRepository(client)
		serviceRepository       = persistence.NewServiceMongoRepository(client)
		establishmentRepository = persistence.NewEstablishmentMongoRepository(client)

		// usecases initialization
		accountSignup = usecase.NewSignup(accountRepository, establishmentRepository)
		accountSignin = usecase.NewAccountSignin(accountRepository)
		createService = usecase.NewCreateService(serviceRepository)
		findServices  = usecase.NewFindServices(serviceRepository)

		// handlers initialization
		authHandler    = handlers.NewAuthHandler(*accountSignup, *accountSignin)
		serviceHandler = handlers.NewServiceHandler(*createService, *findServices)

		// http server initialization
		app  = fiber.New()
		auth = app.Group("/auth")
		api  = app.Group("/api/v1", middleware.JWTAuth)
	)

	// auth
	auth.Post("/signup", authHandler.HandleSignup)
	auth.Post("/signin", authHandler.HandleSignin)

	api.Post("/service", serviceHandler.HandleCreateService)
	api.Get("/service", serviceHandler.HandleFindServicesByEstablishment)

	app.Listen(*listenAddr)
}
