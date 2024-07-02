package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jpmoraess/service-scheduling/configs"
	"github.com/jpmoraess/service-scheduling/internal/application/usecase"
	"github.com/jpmoraess/service-scheduling/internal/infra/http"
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
	signupUseCase := usecase.NewSignupUseCase(accountRepository, professionalRepository, establishmentRepository)
	signinUseCase := usecase.NewSigninUseCase(accountRepository)
	createServiceUseCase := usecase.NewCreateServiceUseCase(serviceRepository)
	findServiceUseCase := usecase.NewFindServiceUseCase(serviceRepository)
	createProfessionalUseCase := usecase.NewCreateProfessionalUseCase(accountRepository, professionalRepository)
	getProfessionalUseCase := usecase.NewGetProfessionalUseCase(professionalRepository)
	createCustomerUseCase := usecase.NewCreateCustomerUseCase(customerRepository)
	findCustomerUseCase := usecase.NewFindCustomerUseCase(customerRepository)
	findProfessionalUseCase := usecase.NewFindProfessionalUseCase(professionalRepository)
	createSchedulingUseCase := usecase.NewCreateSchedulingUseCase(serviceRepository, customerRepository, professionalRepository, establishmentRepository, schedulingRepository)
	requestPasswordResetUseCase := usecase.NewRequestPasswordResetUseCase(accountRepository, passwordResetRepository)
	resetPasswordUseCase := usecase.NewResetPasswordUseCase(accountRepository, passwordResetRepository)

	// initialize http handlers
	authHandler := http.NewAuthHandler(signupUseCase, signinUseCase)
	serviceHandler := http.NewServiceHandler(findServiceUseCase, createServiceUseCase)
	professionalHandler := http.NewProfessionalHandler(getProfessionalUseCase, findProfessionalUseCase, createProfessionalUseCase)
	customerHandler := http.NewCustomerHandler(findCustomerUseCase, createCustomerUseCase)
	schedulingHandler := http.NewSchedulingHandler(createSchedulingUseCase)
	passwordResetHandler := http.NewPasswordResetHandler(resetPasswordUseCase, requestPasswordResetUseCase)

	// initialize fiber application
	app := fiber.New()
	auth := app.Group("/auth")
	apiV1 := app.Group("/api/v1", middleware.JWTAuth(accountRepository, establishmentRepository))

	// set up cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type,Authorization",
	}))

	// set up routes
	auth.Post("/signup", authHandler.HandleSignup)
	auth.Post("/signin", authHandler.HandleSignin)
	auth.Post("/request-password-reset", passwordResetHandler.HandleRequestPasswordReset)
	auth.Post("/reset-password", passwordResetHandler.HandleResetPassword)

	apiV1.Post("/service", serviceHandler.HandleCreateService)
	apiV1.Get("/service", serviceHandler.HandleFindServiceByEstablishment)

	apiV1.Post("/professional", professionalHandler.HandleCreateProfessional)
	apiV1.Get("/professional", professionalHandler.HandleFindProfessional)
	apiV1.Get("/professional/:id", professionalHandler.HandleGetProfessional)

	apiV1.Post("/customer", customerHandler.HandleCreateCustomer)
	apiV1.Get("/customer", customerHandler.HandleFindCustomer)

	apiV1.Post("/scheduling", schedulingHandler.HandleCreateScheduling)

	// start http server
	if err := app.Listen(*listenAddr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
