package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/jeffersonbraster/apigo/configs"
	_ "github.com/jeffersonbraster/apigo/docs"
	"github.com/jeffersonbraster/apigo/internal/entity"
	"github.com/jeffersonbraster/apigo/internal/infra/database"
	"github.com/jeffersonbraster/apigo/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Jefferson Brandão
// @contact.url    http://www.jeffersonbrandao.com.br
// @contact.email  oi@jeffersonbrandao.com.br

// @license.name   Jefferson Brandao License
// @license.url    http://www.jeffersonbrandao.com.br

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
		configs, err := configs.LoadConfig(".")
		if err != nil {
				panic(err)
		}

		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
				panic(err)
		}

		db.AutoMigrate(&entity.Product{}, &entity.User{})
		productDB := database.NewProduct(db)
		productHandle := handlers.NewProductHandle(productDB)

		userDB := database.NewUser(db)
		userHandle := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JWTExpiresIn)

		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		//r.Use(LogRequest)

		r.Route("/products", func(r chi.Router) {
			r.Use(jwtauth.Verifier(configs.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Post("/", productHandle.CreateProduct)
			r.Get("/", productHandle.GetProducts)
			r.Get("/{id}", productHandle.GetProduct)		
			r.Put("/{id}", productHandle.UpdateProduct)
			r.Delete("/{id}", productHandle.DeleteProduct)
		})
		

		r.Post("/users", userHandle.CreateUser)
		r.Post("/users/login", userHandle.GetJwt)

		r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

		http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

