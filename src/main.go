package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong111")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	port := getEnv()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	if err := r.Run(port); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

type Config struct {
	Production EnvConfig `toml:"production"`
	Testing    EnvConfig `toml:"testing"`
	Dev        EnvConfig `toml:"dev"`
}

type EnvConfig struct {
	Port string `toml:"port"`
}

func getEnv() string {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	var config Config
	if _, err := toml.DecodeFile("env/config.toml", &config); err != nil {
		log.Fatalf("Error decoding config.toml: %s", err)
	}

	switch env {
	case "production":
		return config.Production.Port
	case "test":
		return config.Testing.Port
	default:
		return config.Dev.Port
	}
}
