package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func (e *Env) authen(c *gin.Context) {
	token, _ := ioutil.ReadAll(c.Request.Body)

	environment := getEnv("ENV", Environment)

	tokenString, _ := url.QueryUnescape(string(token))
	log.Println("Token = ", string(tokenString))

	jsonData := []byte(`{"data": {"token": "` + string(tokenString) + `"}}`)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	request, _ := http.NewRequest("POST", "authURL", bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)

	auth := AuthRes{}

	if err != nil {
		log.Println("authen error: The HTTP request failed with error = ", err)
	} else {
		defer response.Body.Close()
		json.NewDecoder(response.Body).Decode(&auth)
		if auth.Data.IsLoggedIn {
			log.Println("user: " + auth.Data.UserID + " has been logged in!")
		} else {
			if environment != "local" {
				log.Println("Authenticated error!, the system can't extract user data from token")
			}
		}

		accessToken, err := createToken(auth.Data.UserID)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"isLoggedIn": auth.Data.IsLoggedIn,
			"userId":     auth.Data.UserID,
			"roles":      auth.Data.Roles,
			"token":      accessToken,
		})
	}
}

// getHardCodeDropdown : get all hard code drop down list
func getHardCodeDropdown(c *gin.Context) {
	dropdowns, err := getHcDropdown(c.Request.URL.Path)
	if err != nil {
		log.Println("getHardCodeDropdown error: ", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": dropdowns,
	})
}
