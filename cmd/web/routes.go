package main

import (
	"embed"
	"net/http"
	"time"

	"github.com/yarlson/chistaticmiddleware/static"

	"github.com/go-chi/chi/v5"
)

//go:embed static
var staticFiles embed.FS

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(SessionLoad)

	staticConfig := static.Config{
		Fs:            staticFiles,
		Root:          "static",
		FilePrefix:    "/static",
		CacheDuration: 24 * time.Hour, // Optional: Cache for 24 hours
	}

	mux.Use(static.Handler(staticConfig))

	mux.Get("/", app.Home)
	mux.Get("/virtual-terminal", app.VirtualTerminal)
	mux.Post("/virtual-terminal-payment-succeeded", app.VirtualTerminalPaymentSucceeded)
	mux.Get("/virtual-terminal-receipt", app.VirtualTerminalReceipt)

	mux.Get("/widget/{id}", app.ChargeOnce)
	mux.Post("/payment-succeeded", app.PaymentSucceeded)
	mux.Get("/receipt", app.Receipt)

	mux.Get("/plans/bronze", app.BronzePlan)

	// auth routes
	mux.Get("/login", app.LoginPage)

	mux.Get("/receipt/bronze", app.BronzePlanReceipt)
	// fileServer := http.FileServer(http.Dir("./static"))
	// mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
