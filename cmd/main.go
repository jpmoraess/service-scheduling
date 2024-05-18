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
		customerRepository      = persistence.NewCustomerMongoRepository(client)
		schedulingRepository    = persistence.NewSchedulingMongoRepository(client)
		professionalRepository  = persistence.NewProfessionalMongoRepository(client)
		establishmentRepository = persistence.NewEstablishmentMongoRepository(client)
		passwordResetRepository = persistence.NewPasswordResetMongoRepository(client)

		// usecases initialization
		signup               = usecase.NewSignup(accountRepository, professionalRepository, establishmentRepository)
		signin               = usecase.NewSignin(accountRepository)
		createService        = usecase.NewCreateService(serviceRepository, establishmentRepository)
		findServices         = usecase.NewListServices(serviceRepository)
		createProfessional   = usecase.NewCreateProfessional(accountRepository, professionalRepository, establishmentRepository)
		getProfessional      = usecase.NewGetProfessional(professionalRepository)
		createCustomer       = usecase.NewCreateCustomer(customerRepository, establishmentRepository)
		createScheduling     = usecase.NewCreateScheduling(serviceRepository, customerRepository, professionalRepository, establishmentRepository, schedulingRepository)
		requestPasswordReset = usecase.NewRequestPasswordReset(accountRepository, passwordResetRepository)
		resetPassword        = usecase.NewResetPassword(accountRepository, passwordResetRepository)

		// handlers initialization
		authHandler          = handlers.NewAuthHandler(signup, signin)
		serviceHandler       = handlers.NewServiceHandler(createService, findServices)
		professionalHandler  = handlers.NewProfessionalHandler(createProfessional, getProfessional)
		customerHandler      = handlers.NewCustomerHandler(createCustomer)
		schedulingHandler    = handlers.NewSchedulingHandler(createScheduling)
		passwordResetHandler = handlers.NewPasswordResetHandler(resetPassword, requestPasswordReset)

		// http server initialization
		app  = fiber.New()
		auth = app.Group("/auth")
		api  = app.Group("/api/v1", middleware.JWTAuth(accountRepository))
	)

	// auth
	auth.Post("/signup", authHandler.HandleSignup)
	auth.Post("/signin", authHandler.HandleSignin)
	auth.Post("/request-password-reset", passwordResetHandler.HandleRequestPasswordReset)
	auth.Post("/reset-password", passwordResetHandler.HandleResetPassword)

	api.Post("/service", serviceHandler.HandleCreateService)
	api.Get("/service", serviceHandler.HandleListServicesByEstablishment)

	api.Post("/professional", professionalHandler.HandleCreateProfessional)
	api.Get("/professional/:id", professionalHandler.HandleGetProfessional)

	api.Post("/customer", customerHandler.HandleCreateCustomer)

	api.Post("/scheduling", schedulingHandler.HandleCreateScheduling)

	app.Listen(*listenAddr)
}
