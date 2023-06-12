package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
)

type App struct {
	port string
	test string
}

func main() {
	port, exists := os.LookupEnv("PORT")

	if !exists {
		port = "8080"
	}

	test, exist := os.LookupEnv("TEST")
	if !exist {
		test = "default"
	}

	app := App{port, test}

	log.Println("⏱️  starting on port " + port)
	http.HandleFunc("/", app.handler)
	http.ListenAndServe(":"+port, nil)
}

func (a *App) handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-store")
	randomStr := generateRandomString(10)
	response := fmt.Sprintf("Success \n port: %s \n test: %s \n random-string: %s", a.port, a.test, randomStr)
	fmt.Fprint(w, response)
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(b)
}
