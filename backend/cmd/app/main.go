package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nrf24l01/k8s-pobeda/internal/stats"
	transporthttp "github.com/nrf24l01/k8s-pobeda/internal/transport/http"
)

func main() {
	if err := loadDotEnv(".env"); err != nil {
		log.Printf(".env not loaded: %v", err)
	}

	addr := envOrDefault("HTTP_ADDR", ":8080")
	corsAllowOrigin := envOrDefault("CORS_ALLOW_ORIGIN", "*")

	provider, err := stats.NewKubernetesProvider()
	if err != nil {
		log.Fatalf("failed to initialize kubernetes provider: %v", err)
	}

	handler := transporthttp.NewServer(provider, corsAllowOrigin)

	log.Printf("server listening on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func loadDotEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		if key == "" {
			continue
		}

		if _, exists := os.LookupEnv(key); exists {
			continue
		}

		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, "\"")
		_ = os.Setenv(key, value)
	}

	return scanner.Err()
}
