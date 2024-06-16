package main

import (
    "log"
    "github.com/philipphahmann/rating-app/router"
)

func main() {
    r := router.SetupRouter()
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
