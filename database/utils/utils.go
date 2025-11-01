package utils

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
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
						WHERE id = ? AND expired_at > ?;
						`, id, now).Scan(&redirectUrl)

	if err == sql.ErrNoRows {
		fmt.Println("URL expired or not found")
		return "", nil
	} else if err != nil {
		log.Fatal(err)
	}

	return redirectUrl, nil

}

func IsUsedId(id string) bool {
	var exists bool
	err := database.DB.QueryRow(`SELECT EXISTS(SELECT 1 FROM short_urls WHERE ide = ? LIMIT 1)`, id).Scan(&exists)
	if err != nil {
		return false
	}
	return true
}

func GenerateShortId() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	if IsUsedId(string(b)) {
		return GenerateShortId()
	}

	return string(b)

}

func Create(url string, expire int) (models.NewUrl, error) {
	var newUrl models.NewUrl
	expiredAt := time.Now().Add(time.Duration(expire) * time.Second).Format("2006-01-02 15:04:05")
	newUrl.DeleteId = uuid.New().String()
	newUrl.ShortId = GenerateShortId()
	_, err := database.DB.Exec(`
		INSERT INTO short_urls
		(id,redirect_url,delete_id,expired_at)
		VALUES
		(?,?,?,?);
	`, newUrl.ShortId, url, newUrl.DeleteId, expiredAt)

	return newUrl, err

}

func Delete(deleteId string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	_, _ = database.DB.Exec(`
		UPDATE short_urls
		SET expired_at = ?
		WHERE delete_id = ?
	`, now, deleteId)
}
