package controller

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/people/v1"
)

func Oauth(c echo.Context) (err error) {
	code := c.QueryParam("code")
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Println("Unable to read client secret file: %v", err)
	}


	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b,
		people.ContactsReadonlyScope,
		people.UserinfoProfileScope,
		people.UserinfoEmailScope,
		people.UserBirthdayReadScope,
	)
	if err != nil {
		log.Println("Unable to parse client secret file to config: %v", err)
	}


	if code == "" {
		authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		http.Redirect(c.Response(), c.Request(), authURL, 302)
		return nil
	}


	tok, err := config.Exchange(context.TODO(), code)
	if err != nil {
		log.Println("Unable to retrieve token from web: %v", err)
	}
	

	client := config.Client(context.Background(), tok)
	srv, err := people.New(client)
	if err != nil {
		log.Println("Unable to create people Client %v", err)
	}


	profile, err := srv.People.Get("people/me").PersonFields("names,emailAddresses,addresses,coverPhotos,photos," +
		"birthdays,phoneNumbers").Do()
	if err != nil {
		log.Println("Unable to retrieve people. %v", err)
	}

	return c.JSON(200, profile)
}
