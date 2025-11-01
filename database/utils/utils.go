package utils

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/isa0-gh/urlshorter/database"
	"github.com/isa0-gh/urlshorter/models"
)

func GetUrl(id string) (string, error) {
	var redirectUrl string
	now := time.Now().Format("2006-01-02 15:04:05")
	err := database.DB.QueryRow(`
						SELECT redirect_url 
						FROM short_urls
						WHERE id = ? AND expired_at > ?
						`, id, now).Scan(&redirectUrl)

	if err == sql.ErrNoRows {
		fmt.Println("URL expired or not found")
		return "", nil
	} else if err != nil {
		log.Fatal(err)
	}

	return redirectUrl, nil

}

func Create(url string, expire int) models.NewUrl {
	var newUrl models.NewUrl
	expiredAt := time.Now().Add(time.Duration(expire) * time.Second).Format("2006-01-02 15:04:05")
	deleteId := uuid.New().String()
	shortId := GenerateShortId()
	err := database.Exec(`
		INSERT INTO short_urls
		(id,redirect_url,delete_id,expired_at)
		VALUES
		(?,?,?,?)
	`, shortId, url, deleteId, expiredAt)

}
