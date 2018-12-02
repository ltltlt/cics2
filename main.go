package main

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"

	"cytus2.rocks/cics2/models/user"
)

func main() {
	ConfigRuntime()
	StartWorkers()
	StartGin()
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartWorkers start starsWorker by goroutine.
func StartWorkers() {
	go statsWorker()
}

// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	accounts := getAllAcounts()

	router := gin.New()
	router.Use(gin.BasicAuth(accounts), gin.Recovery())
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")
	router.GET("/", index)
	router.GET("/room/:roomid", roomGET)
	router.POST("/room-post/:roomid", roomPOST)
	router.GET("/stream/:roomid", streamRoom)

	router.Run(":8080")
}

func getAllAcounts() map[string]string {
	users := user.GetAllUsers()
	accounts := make(map[string]string, len(users))
	for _, user := range users {
		accounts[user.ID] = user.Pass
	}
	return accounts
}
