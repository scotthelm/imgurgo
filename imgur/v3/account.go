package v3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Account struct {
	Id             int     `json:"id"`
	Url            string  `json:"url"`
	Bio            string  `json:"bio"`
	Reputation     float64 `json:"reputation"`
	CreatedSeconds int64   `json:"created"`
	ProExpiration  bool    `json:"pro_expiration"`
}

type AccountResponse struct {
	Data    Account `json:"data"`
	Status  int     `json:"status"`
	Success bool    `json:"success"`
}

func (cl *ImgurClient) GetAccount(username string) (Account, error) {
	ar := AccountResponse{}
	request, _ := cl.prepareRequest("GET", "account/me")
	response, err := cl.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return Account{}, err
	}
	return ar.Data, err
}
