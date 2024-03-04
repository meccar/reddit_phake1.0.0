package util

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// var (
// 	DB_URL, DB_DRIVER, JWT_SECRET_KEY, PORT string
// )

var (
	idGeneratorMu sync.Mutex
	idCounter     int32
)

// func init() {
// 	err := godotenv.Load("./environment.env")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	DB_URL = os.Getenv("DB_URL")
// 	fmt.Printf("\n DB_URL : %+v\n", DB_URL)

// 	DB_DRIVER = os.Getenv("DB_DRIVER")
// 	fmt.Printf("\n DB_DRIVER : %+v\n", DB_DRIVER)

// 	PORT = os.Getenv("PORT")
// 	fmt.Printf("\n PORT : %+v\n", PORT)

// 	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
// 	fmt.Printf("\n DB_URL : %+v\n", JWT_SECRET_KEY)

// }
// type Config struct {
// 	DBURL        string
// 	DBDriver     string
// 	Port         string
// 	JWTSecretKey string
// }

// func Init() (*Config, error) {
// 	err := godotenv.Load("./config.env")
// 	if err != nil {
// 		return nil, err
// 	}

// 	config := &Config{
// 		DBURL:        os.Getenv("DB_URL"),
// 		DBDriver:     os.Getenv("DB_DRIVER"),
// 		Port:         os.Getenv("PORT"),
// 		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
// 	}

// 	fmt.Printf("\n DB_URL: %s\n", config.DBURL)
// 	fmt.Printf("\n DB_DRIVER: %s\n", config.DBDriver)
// 	fmt.Printf("\n PORT: %s\n", config.Port)
// 	fmt.Printf("\n JWT_SECRET_KEY: %s\n", config.JWTSecretKey)

// 	return config, nil
// }

type Config struct {
	Environment          string        `mapstructure:"ENVIRONMENT"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

const validChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// func InitializeTemplates() error {
// 	templates, err := template.ParseGlob("../templates/*.html")
// 	if err != nil {
// 		return err
// 	}

// 	Templates = templates

// 	return nil
// }

// GenerateID generates a unique int32 ID using a combination of timestamp and counter.
func GenerateID() string {
	idGeneratorMu.Lock()
	defer idGeneratorMu.Unlock()

	// Use a combination of timestamp, counter, and random characters
	id := int32(time.Now().UnixNano() / 1e6) // Milliseconds
	idCounter++

	// Generate random characters
	randomChars := make([]byte, 8)
	for i := range randomChars {
		randomChars[i] = validChars[rand.Intn(len(validChars))]
	}

	// Combine timestamp, counter, and random characters
	idString := fmt.Sprintf("%08d%s", id, string(randomChars))

	return idString[:8]
}

func generateUniqueID() string {
	return uuid.New().String()
}
