package engine

import (
	"app/config"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	TemplatesPath      = "../../templates/**/*"
	StaticPath         = "../../static"
	StaticRelativePath = "static"
	EnvFilesPath       = "../../envs/"
)

func Setup(envFile string) *gin.Engine {
	err := godotenv.Load(EnvFilesPath + envFile)
	if err != nil {
		panic("Could not load environment file")
	}
	fmt.Println("Environment:", os.Getenv("ENV"))

	e := gin.Default()
	e.LoadHTMLGlob(TemplatesPath)
	e.Static(StaticRelativePath, StaticPath)

	config.SetupMiddlewares(e)
	config.SetupRoutes(e)

	return e
}
