package main

import (
	"fmt"

	"github.com/MihaiChirilov/go-redactor/redactor"
)

func main() {
    // Example usage
    config := map[string]interface{}{
        "someKey": "someValue",
        "password": "secret",
        "url": "http://example.com",
    }

    anonymizedConfig, err := redactor.Anonymize(config)
    if err != nil {
        fmt.Println("Error anonymizing config:", err)
        return
    }

    fmt.Println("Anonymized config:", anonymizedConfig)

    redactedConfig, err := redactor.RemoveCredentials(config)
    if err != nil {
        fmt.Println("Error removing credentials:", err)
        return
    }

    fmt.Println("Redacted config:", redactedConfig)
}
