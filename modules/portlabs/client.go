package portlabs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var Client PortClient

type AccessTokenRequest struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type AccessTokenResponse struct {
	Ok          bool   `json:"ok"`
	AccessToken string `json:"accessToken"`
	ExpiresIn   int64  `json:"expiresIn"`
	TokenType   string `json:"tokenType"`
}

type PortClient struct {
	token string
}

func NewPortClient(clientId string, clientSecret string) (*PortClient, error) {
	client := &PortClient{}
	err := client.authenticate(clientId, clientSecret)

	if err != nil {
		return nil, fmt.Errorf("authentication failed: %v", err)
	}

	return client, nil
}

func (c PortClient) GetCluster(id string) (*ClusterEntity, error) {
	cluster := &ClusterEntity{}

	err := c.request(
		fmt.Sprintf("https://api.getport.io/v1/blueprints/clusters/entities/%s", id),
		cluster,
	)
	if err != nil {
		return nil, err
	}

	return cluster, nil
}

func (c PortClient) request(url string, v any) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", c.token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("request failed with status %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func (c *PortClient) authenticate(clientId string, clientSecret string) error {
	body, err := json.Marshal(AccessTokenRequest{
		ClientID:     clientId,
		ClientSecret: clientSecret,
	})
	if err != nil {
		return err
	}

	httpResponse, err := http.Post("https://api.getport.io/v1/auth/access_token", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	response := &AccessTokenResponse{}
	err = json.NewDecoder(httpResponse.Body).Decode(response)
	if err != nil {
		return err
	}

	if !response.Ok {
		return fmt.Errorf("access token request failed because of not being OK")
	}

	c.token = response.AccessToken

	return nil
}
