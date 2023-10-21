package utils

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"url-shortner/internal/app/model"

	"github.com/go-redis/redis"
)

type Utility interface {
	ConvertLongURLToShortURL(input string) (*model.URLShortenResponse, error)
}

type UtilRequest struct {
	file *os.File
}

func NewUtilityService(redisClient *redis.Client) Utility {
	fileExists, err := os.Open("/Users/user1/go/src/github.com/url-shortner/pkg/data/url_store.json")
	if err != nil {
		log.Fatal("failed to locate url store file")
	}

	return &UtilRequest{
		file: fileExists,
	}
}

func (u *UtilRequest) ConvertLongURLToShortURL(input string) (*model.URLShortenResponse, error) {
	urlStore, err := u.readURLStore()
	if err != nil {
		return nil, err
	}

	if val, ok := urlStore[input]; ok {
		return val, nil
	}

	output := u.baseConversion(input)
	response := &model.URLShortenResponse{
		ShortURL:  output,
		CreatedAt: time.Now(),
	}

	return response, nil
}

func (u *UtilRequest) baseConversion(input string) string {
	charSet := "ABCDEFGHIIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var (
		sb    strings.Builder
		index = 7
	)
	for index > 0 {
		sb.WriteString(string(charSet[rand.Intn(67)%62]))
		index--
	}

	return sb.String()
}

func (u *UtilRequest) readURLStore() (map[string]*model.URLShortenResponse, error) {
	fileRead, err := io.ReadAll(u.file)
	if err != nil {
		return nil, err
	}

	var urlStore = make(map[string]*model.URLShortenResponse)
	err = json.Unmarshal(fileRead, &urlStore)
	if err != nil {
		return nil, err
	}

	return urlStore, nil
}

func (u *UtilRequest) writeURLStore(input map[string]*model.URLShortenResponse) error {
	data, _ := json.Marshal(input)
	_, err := io.WriteString(u.file, string(data))
	if err != nil {
		return err
	}

	return nil
}
