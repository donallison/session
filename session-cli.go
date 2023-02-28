package session

import (
	"encoding/json"
	"log"
	"os"
)

type Client struct {
	AuthToken string `json:"authToken"`
	Email     string `json:"email"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}

const (
	SessionFile = "session.json"
)

// LoadSession will open the session file
func (c *Client) LoadSession() (*Client, error) {
	log.Println("Function: LoadSession")
	var client *Client

	// check if SessionFile exists
	if _, err := os.Stat(SessionFile); err == nil {
		log.Printf("%s File exists\n", SessionFile)
	} else {
		log.Printf("%s File does not exist. Creating....\n", SessionFile)
		// create empty file
		err := c.SaveSession()
		if err != nil {
			log.Println(err)
			return client, err
		}
	}

	data, err := os.ReadFile(SessionFile)
	if err != nil {
		log.Println(err)
		return client, err
	}

	// log.Printf("Data from sessionFile: %s", data)
	err = json.Unmarshal(data, &client)

	if err != nil {
		log.Println(err)
		return client, err
	}

	return client, nil
}

func (c *Client) SaveSession() error {
	log.Println("Function: SaveSession")
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	// log.Printf("Data: %v\n", string(data))
	err = os.WriteFile(SessionFile, data, 0600)
	if err != nil {
		return err
	}

	return nil
}
