package main

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/yehormironenko/reseller/config"
	"github.com/yehormironenko/reseller/config/client"
	"github.com/yehormironenko/reseller/internal/controller/handlers"
	"github.com/yehormironenko/reseller/internal/repository"
	"github.com/yehormironenko/reseller/internal/service"
)

// 2 function search bookByName // GET with params
// Buy
func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Logger.SetPrefix("reseller")
	e.Use(middleware.Logger())

	cfg, err := loadConfig()
	if err != nil {
		//logger.Error().Msg("Cannot read config file: " + err.Error())
		return
	}

	postgresClient := client.NewPostgresClient(cfg.Postgres)

	bookRepository := repository.NewBookRepository(postgresClient)

	bookService := service.NewBookService(bookRepository)

	// custom handlers
	handlerWithBookService := &handlers.HandlerWithService{BookService: bookService}

	e.GET("/echo", handlers.Echo)
	e.GET("/search", handlerWithBookService.GetBook)
	e.POST("/buy", handlerWithBookService.BuyBook)

	e.Logger.Fatal(e.Start(cfg.Server.Endpoint))
}

func loadConfig() (config.Config, error) {
	k := koanf.New(".")
	if err := k.Load(rawbytes.Provider([]byte(config.Yaml)), yaml.Parser()); err != nil {
		return config.Config{}, err
	}

	return config.Bind(k), nil
}
