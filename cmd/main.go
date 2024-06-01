package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// initialize repositories
	accountRepository := persistence.NewAccountMongoRepository(client)
	serviceRepository := persistence.NewServiceMongoRepository(client)
	customerRepository := persistence.NewCustomerMongoRepository(client)
	schedulingRepository := persistence.NewSchedulingMongoRepository(client)
	professionalRepository := persistence.NewProfessionalMongoRepository(client)
	establishmentRepository := persistence.NewEstablishmentMongoRepository(client)
	passwordResetRepository := persistence.NewPasswordResetMongoRepository(client)

	// initialize use cases
	signup := usecase.NewSignup(accountRepository, professionalRepository, establishmentRepository)
	signin := usecase.NewSignin(accountRepository)
	createService := usecase.NewCreateService(serviceRepository)
	findService := usecase.NewFindService(serviceRepository)
	createProfessional := usecase.NewCreateProfessional(accountRepository, professionalRepository)
	getProfessional := usecase.NewGetProfessional(professionalRepository)
	createCustomer := usecase.NewCreateCustomer(customerRepository)
	findCustomer := usecase.NewFindCustomer(customerRepository)
	findProfessional := usecase.NewFindProfessional(professionalRepository)
	createScheduling := usecase.NewCreateScheduling(serviceRepository, customerRepository, professionalRepository, establishmentRepository, schedulingRepository)
	requestPasswordReset := usecase.NewRequestPasswordReset(accountRepository, passwordResetRepository)
	resetPassword := usecase.NewResetPassword(accountRepository, passwordResetRepository)

	// initialize handlers
	authHandler := handlers.NewAuthHandler(signup, signin)
	serviceHandler := handlers.NewServiceHandler(createService, findService)
	professionalHandler := handlers.NewProfessionalHandler(createProfessional, getProfessional, findProfessional)
	customerHandler := handlers.NewCustomerHandler(createCustomer, findCustomer)
	schedulingHandler := handlers.NewSchedulingHandler(createScheduling)
	passwordResetHandler := handlers.NewPasswordResetHandler(resetPassword, requestPasswordReset)

	// initialize fiber app
	app := fiber.New()

	// set up cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type,Authorization",
	}))

	// set up routes
	setupRoutes(app, accountRepository, establishmentRepository, authHandler, serviceHandler, customerHandler, schedulingHandler, professionalHandler, passwordResetHandler)

	// start server
	if err := app.Listen(*listenAddr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func setupRoutes(
	app *fiber.App,
	accountRepository *persistence.AccountMongoRepository,
	establishmentRepository *persistence.EstablishmentMongoRepository,
	authHandler *handlers.AuthHandler,
	serviceHandler *handlers.ServiceHandler,
	customerHandler *handlers.CustomerHandler,
	schedulingHandler *handlers.SchedulingHandler,
	professionalHandler *handlers.ProfessionalHandler,
	passwordResetHandler *handlers.PasswordResetHandler,
) {
	auth := app.Group("/auth")
	api := app.Group("/api/v1", middleware.JWTAuth(accountRepository, establishmentRepository))

	// auth routes
	auth.Post("/signup", authHandler.HandleSignup)
	auth.Post("/signin", authHandler.HandleSignin)
	auth.Post("/request-password-reset", passwordResetHandler.HandleRequestPasswordReset)
	auth.Post("/reset-password", passwordResetHandler.HandleResetPassword)

	// service routes
	api.Post("/service", serviceHandler.HandleCreateService)
	api.Get("/service", serviceHandler.HandleFindServiceByEstablishment)

	// professional routes
	api.Post("/professional", professionalHandler.HandleCreateProfessional)
	api.Get("/professional", professionalHandler.HandleFindProfessional)
	api.Get("/professional/:id", professionalHandler.HandleGetProfessional)

	// customer routes
	api.Post("/customer", customerHandler.HandleCreateCustomer)
	api.Get("/customer", customerHandler.HandleFindCustomer)

	// scheduling routes
	api.Post("/scheduling", schedulingHandler.HandleCreateScheduling)
}
