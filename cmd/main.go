package main

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"reseller/config"
	"reseller/config/client"
	"reseller/internal/controller/handlers"
	"reseller/internal/repository"
	"reseller/internal/service"
)

// 2 function search bookByName // GET with params
// Buy
func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())

	cfg, err := loadConfig()
	if err != nil {
		//logger.Error().Msg("Cannot read config file: " + err.Error())
		return
	}

	postgresClient := client.NewPostgresClient(cfg.Postgres)

	bookRepository := repository.NewBookRepository(postgresClient)

	bookService := service.NewBookService(bookRepository)
	/*	book := new(entities.Book)
		res := postgresClient.NewSelect().Model(book).Where("id = ?", 1).Scan(context.Background())
		fmt.Println(res, book)*/
	handlerGB := &handlers.HandlerGetBook{BookService: bookService}

	e.GET("/echo", handlers.Echo)
	e.GET("/search", handlerGB.GetBook)

	e.Logger.Fatal(e.Start(cfg.Server.Endpoint))
}

func loadConfig() (config.Config, error) {
	k := koanf.New(".")
	if err := k.Load(rawbytes.Provider([]byte(config.Yaml)), yaml.Parser()); err != nil {
		return config.Config{}, err
	}

	return config.Bind(k), nil
}
