package http

import (
	"encoding/base64"
	urlverifier "github.com/davidmytton/url-verifier"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
	"url-shortener-api/db"
	"url-shortener-api/db/repository"
)

type Request struct {
	Id   int64  `form:"id" json:"id"`
	Code string `form:"code" json:"code"`
	Url  string `form:"url" json:"url"`
}

func (s *Server) GenerateShortUrl(ginCtx *gin.Context) {
	request, err := s.parseRequest(ginCtx)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, err.Error())
		log.Println(err.Error())
		return
	}

	isValid, err := isUrlValid(request.Url)
	if err != nil || !isValid {
		ginCtx.JSON(http.StatusBadRequest, err)
		return
	}

	urlsRepository := repository.NewUrlsRepository(db.GetInstance().DB)
	urlEntity, err := urlsRepository.Get(repository.Urls{Url: request.Url})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, "error on getting url from database: "+err.Error())
		return
	}

	if urlEntity != nil {
		ginCtx.JSON(http.StatusOK, urlEntity)
	} else {
		urlCode, _ := generateRandomString(10)
		// TODO - Validate urk code is exists in DB
		newUrlRecord := repository.Urls{Code: urlCode, Url: request.Url, IP: ginCtx.ClientIP(), CreatedAt: time.Now()}
		_, err = urlsRepository.Create(newUrlRecord)
		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, "error on create url in database: "+err.Error())
			return
		}

		ginCtx.JSON(http.StatusOK, newUrlRecord)
	}
}

func (s *Server) Redirect(ginCtx *gin.Context) {
	code := ginCtx.Param("code")
	if code == "" {
		ginCtx.JSON(http.StatusNoContent, "No results")
		return
	}

	// Get from DB
	urlsRepository := repository.NewUrlsRepository(db.GetInstance().DB)
	url, err := urlsRepository.Get(repository.Urls{Code: code})
	if err != nil {
		log.Println(err.Error())
		ginCtx.JSON(http.StatusInternalServerError, "error on getting urls from database: "+err.Error())
		return
	}

	if url != nil {
		log.Println("url found in database")

		visitorsRepository := repository.NewVisitorsRepository(db.GetInstance().DB)
		_, err = visitorsRepository.Create(repository.Visitors{
			IP:        ginCtx.ClientIP(),
			UrlId:     url.ID,
			VisitedAt: time.Now(),
		})
		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, "error on create visitor in database: "+err.Error())
			return
		}

		log.Println("return urls from database")
		ginCtx.Redirect(http.StatusPermanentRedirect, url.Url)
	} else {
		log.Println("not found urls in database")

		// Redirect
		ginCtx.JSON(http.StatusNoContent, "No results")
	}
}

func (s *Server) parseRequest(ginCtx *gin.Context) (*Request, error) {
	req := &Request{}
	err := ginCtx.BindJSON(&req)
	if err != nil {
		return nil, err
	}

	return req, err
}

// isUrlValid verify slice url as syntax and existence.
// Retrieves a true or false.
// If occur error will be returned too.
func isUrlValid(url string) (bool, error) {
	urlVerifier := urlverifier.NewVerifier()
	//urlVerifier.EnableHTTPCheck()
	verifiedUrl, err := urlVerifier.Verify(url)
	if err != nil {
		return false, err
	}

	return verifiedUrl.IsURL, nil
}

func generateRandomString(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buffer)[:length], nil
}
