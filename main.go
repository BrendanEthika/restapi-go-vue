package main

import "os"

func main() {
	app := App{}
	os.Setenv("TEST_DB_USERNAME", "restapi2")
	os.Setenv("TEST_DB_PASSWORD", "assword")
	os.Setenv("TEST_DB_NAME", "restgovue")
	os.Setenv("TEST_DB_HOST", "localhost")

	app.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"),
		os.Getenv("TEST_DB_HOST"),
		"5432",
		"disable")

	port := os.Getenv("PORT")
	if port == "" {
		port = "5678"
	}
	app.Run(":" + port)
}
