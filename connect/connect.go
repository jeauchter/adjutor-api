package connect

import (
	"fmt"
	"log"
	"os"

	"github.com/jeremyauchter/adjutor/models/auth"
	"github.com/jeremyauchter/adjutor/models/products"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	Auth    *gorm.DB
	Product *gorm.DB
}

var productModels = []interface{}{
	&products.Address{},
	&products.Audience{},
	&products.Class{},
	&products.Country{},
	&products.Department{},
	&products.ItemVariant{},
	&products.ProductType{},
	&products.Tag{},
	&products.TagProductMap{},
	&products.Item{},
	&products.Product{},
	&products.Style{},
	&products.Vendor{},
}

func (server *Server) Connect() {
	var err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env in connect, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.InitializeAuth(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"))
	server.Auth.Debug().AutoMigrate(&auth.User{}) //database migration
	server.InitializeProduct(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"))
	server.Product.Debug().AutoMigrate(productModels...)
}

func (server *Server) InitializeAuth(DbUser, DbPassword, DbPort, DbHost string) {
	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, "auth")
	server.Auth, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err == nil {
		fmt.Printf("We are connected to the %s database", "auth")
	} else {
		fmt.Printf("Cannot connect to %s database", "auth")
		fmt.Print(DBURL)
		log.Fatal("This is the error:", err)
	}
}

func (server *Server) InitializeProduct(DbUser, DbPassword, DbPort, DbHost string) {
	var err error
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, "product")
	server.Product, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", "product")
		fmt.Print(DBURL)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", "product")
	}
}
