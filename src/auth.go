/*
Copyright © 2024 Donald Gifford dgifford06@gmail.com
*/
package src

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/smartdevicemanagement/v1"
)

type AuthConfig struct {
	c    context.Context
	conf oauth2.Config
}

func InitConfig(ctx context.Context) AuthConfig {
	clientID := viper.GetString("auth.client_id")
	clientSecret := viper.GetString("auth.client_secret")
	projectID := viper.GetString("nest.project_id")

	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  viper.GetString("auth.redirect_uri"),
		Scopes: []string{
			smartdevicemanagement.SdmServiceScope,
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("https://nestservices.google.com/partnerconnections/%s/auth", projectID),
			TokenURL: google.Endpoint.TokenURL,
		},
	}

	return AuthConfig{c: ctx, conf: *conf}
}

func localTokenExists() bool {
	tokenFileName := viper.GetString("auth.token_file_name")
	if _, err := os.Stat(tokenFileName); errors.Is(err, os.ErrNotExist) {
		log.Println("No token file found")
		return false
	} else {
		log.Println("Token file found")
		return true
	}
}

func (a *AuthConfig) GetToken() *oauth2.Token {
	var localToken *oauth2.Token
	tokenFileName := viper.GetString("auth.token_file_name")
	localToken = a.ReadTokenFromFile(tokenFileName)
	fmt.Println(localToken)
	if localToken.Expiry.Before(time.Now()) {
		fmt.Println("local token expired")

		src := a.conf.TokenSource(a.c, localToken)
		fmt.Println(src.Token())
		newToken, err := src.Token()
		if err != nil {
			log.Fatal(err)
		}
		if newToken.AccessToken != localToken.AccessToken {
			a.SaveTokenToFile(newToken)
			localToken = newToken
		}
	}
	return localToken
}

func (a *AuthConfig) SaveTokenToFile(oAuthToken *oauth2.Token) {
	// Save the token to a file
	file, err := os.Create("token.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p, err := json.Marshal(oAuthToken)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(p)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *AuthConfig) ReadTokenFromFile(tokenFileName string) *oauth2.Token {
	// Read the token from a file
	file, err := os.ReadFile(tokenFileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(file))
	var tt *oauth2.Token

	err = json.Unmarshal(file, &tt)
	if err != nil {
		log.Fatal(err)
	}

	return tt
}

func Authenticate(conf AuthConfig) *oauth2.Token {
	// check for local token file first
	if localTokenExists() {
		// read local token, check expiry, refresh if needed
		return conf.GetToken()
	} else {
		// create config, genererate token using web flow
		return conf.authenticate()
	}
}

func (a *AuthConfig) authenticate() *oauth2.Token {
	log.Print("Your browser will be opened to authenticate with Google")
	log.Print("Hit enter to confirm and continue")
	url := a.conf.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	codeChannel := make(chan string)
	mux := http.NewServeMux()
	server := http.Server{Addr: ":8080", Handler: mux}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<p style=font-size:xx-large;text-align:center>return to your terminal</p>"))
		codeChannel <- r.URL.Query().Get("code")
	})
	openBrowser(url)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	code := <-codeChannel
	log.Print("Shutting down server")
	server.Shutdown(context.Background())
	log.Print("Exchanging token")
	log.Print("Code: ", code)
	token, err := a.conf.Exchange(a.c, code, oauth2.AccessTypeOffline)
	if err != nil {
		log.Fatal(err)
	}
	a.SaveTokenToFile(token)
	return token
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
