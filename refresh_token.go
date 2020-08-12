package contaazul

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

func refreshToken(cli *Client) error {
	url := fmt.Sprintf("%s/oauth2/token?grant_type=refresh_token&refresh_token=%s", cli.baseURL, cli.refreshToken)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Printf("could not create request, err: %v", err)
		return err
	}

	data := fmt.Sprintf("%s:%s", cli.clientID, cli.clientSecret)
	dataEncoded := b64.StdEncoding.EncodeToString([]byte(data))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", dataEncoded))

	resp, err := cli.client.Do(req)
	if err != nil {
		log.Printf("could get a refreshed token, err: %v", err)
		return err
	}

	defer resp.Body.Close() // nolint: errcheck
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("could not read response body, err: %v", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("could not refresh token, invalid status code: %d", resp.StatusCode)
		log.Printf("%s, body: %s", msg, body)
		return fmt.Errorf(msg)
	}

	atr := AccessTokenResponse{}
	err = json.Unmarshal(body, &atr)

	if err != nil {
		msg := fmt.Sprintf("could not unmarshal access token response, error: %v, payload: %s", err, body)
		log.Println(msg)
		return fmt.Errorf(msg)
	}

	cli.accessToken = atr.AccessToken
	cli.refreshToken = atr.RefreshToken
	log.Printf("Token sucessfully refreshed, %s | %s", cli.accessToken, cli.refreshToken)

	return err
}
