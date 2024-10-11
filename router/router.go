package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/adapters/api"
	"github.com/ppwlsw/sa-project-backend/adapters/database"
	"github.com/ppwlsw/sa-project-backend/usecases"
	"gorm.io/gorm"
)

func SetUpRouters(app *fiber.App, db *gorm.DB) {

	userRepo := database.ProvideUserPostgresRepository(db)
	userService := usecases.ProvideUserService(userRepo)
	userHandler := api.ProvideUserHandler(userService)

	productRepo := database.InitiateProductPostGresRepository(db)
	productService := usecases.InitiateProductsService(productRepo)
	productHandler := api.InitiateProductHandler(productService)

	transactionRepo := database.InitiateTransactionPostGresRepository(db)
	transactionService := usecases.InitiateTransactionService(transactionRepo)
	transactionHandler := api.InitiateTransactionHandler(transactionService)

	authService := usecases.ProvideAuthService(userRepo)
	authHandler := api.ProvideAuthHandler(authService)

	shipmentRepo := database.InitiateShipmentPostgresRepository(db)
	shipmentService := usecases.InitiateShipmentService(shipmentRepo)
	shipmentHandler := api.InitiateShipmentHandler(shipmentService)

	orderRepo := database.InitiateOrderPostgresRepository(db)
	orderService := usecases.InitiateOrderService(orderRepo)
	orderHandler := api.InitiateOrderHandler(orderService)

	packageRepo := database.InitiatePackagePostgresRepository(db)
	packageService := usecases.InitiatePackageService(packageRepo)
	packageHandler := api.InitiatePackageHandler(packageService)

	handlers := api.ProvideHandlers(
		userHandler, productHandler, transactionHandler, authHandler, shipmentHandler, orderHandler, packageHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//User
	app.Get("/users", handlers.UserHandler.GetAllUsers)
	app.Get("/users/:id", handlers.UserHandler.GetUserByID)

	//Product
	app.Post("/products", handlers.ProductHandler.CreateProduct)
	app.Get("/products", handlers.ProductHandler.GetAllProducts)
	app.Get("product/:id", handlers.ProductHandler.GetProductByID)
	app.Put("products/:id", handlers.ProductHandler.UpdateProduct)

	//Transaction
	app.Post("/transaction", handlers.TransactionHandler.CreateTransaction)
	app.Post("/transactions", handlers.TransactionHandler.CreateTransaction)
	app.Get("/transactions", handlers.TransactionHandler.GetAllTransactions)
	app.Get("/transaction/:id", handlers.TransactionHandler.GetTransactionById)
	app.Put("/transaction/:id", handlers.TransactionHandler.UpdateTransaction)

	//Auth
	app.Post("/register", handlers.AuthHandler.Register)
	app.Post("/login", handlers.AuthHandler.Login)

	//Shipment
	app.Post("/shipment", handlers.ShipmentHandler.CreateShipment)
	app.Post("/shipments", handlers.ShipmentHandler.CreateShipment)
	app.Get("/shipments", handlers.ShipmentHandler.GetAllShipments)
	app.Get("/shipment/:id", handlers.ShipmentHandler.GetShipmentByID)
	app.Put("/shipment/:id", handlers.ShipmentHandler.UpdateShipment)

	//Order
	app.Post("/order", handlers.OrderHandler.CreateOrder)
	app.Post("/orders", handlers.OrderHandler.CreateOrder)
	app.Get("/orders", handlers.OrderHandler.GetAllOrders)
	app.Get("/order/:id", handlers.OrderHandler.GetOrderByID)
	app.Put("/order/:id", handlers.OrderHandler.UpdateOrder)

	//Package
	app.Post("/packages", packageHandler.CreatePackage)
	app.Get("/packages/:id", packageHandler.GetPackageByID)
	app.Get("/orders/:orderID/packages", packageHandler.GetAllPackagesByOrderID)
	app.Get("/packages", packageHandler.GetAllPackages)

}
