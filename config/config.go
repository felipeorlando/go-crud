package config

// Configs on dev enviroments
const (
	// Database dev configs
	DbURIDev  string = "mongodb://mongo:27017"
	DbNameDev string = "users-crud-dev"
	// Databse test configs
	DbURITest  string = "mongodb://localhost:27017"
	DbNameTest string = "users-crud-test"
	// API Server port
	APIPort string = ":3000"
)
