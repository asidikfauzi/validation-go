package app

import (
	"fmt"
	modelDomain "test-prepare/model/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetEnv(key string) string {
	var appConfig map[string]string
	appConfig, err := godotenv.Read()

	if err != nil {
		fmt.Println("Error reading .env file")
	}
	return appConfig[key]
}

func InitConfig() (*gorm.DB, error) {
	var err error
	postgresCredential := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		GetEnv("DB_HOST"),
		GetEnv("DB_USER"),
		GetEnv("DB_PASSWORD"),
		GetEnv("DB_NAME"),
		GetEnv("DB_PORT"),
		GetEnv("SSL_MODE"),
	)

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: postgresCredential,
	}), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
		return nil, err
	}

	fmt.Println("CONNECT DATABASE SUCCESSFULLY")
	//InitMigrate()

	return DB, nil
}

func InitMigrate() {
	DB.AutoMigrate(&modelDomain.Users{})
	DB.Where("1=1").Delete(&modelDomain.Users{})
}
