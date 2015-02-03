// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Key represents a public SSH key used to authenticate a user or deploy script.
type Key struct {
	ID    int    `json:"id,omitempty"`
	Key   string `json:"key,omitempty"`
	URL   string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
}

type AddSSHKeyOption struct {
	Title string `json:"title"`
	Key   string `json:"key"`
}

// ListMyKeys lists all repositories for the authenticated user that has access to.
func (c *Client) ListMyKeys() ([]*Key, error) {
	keys := make([]*Key, 0, 10)
	err := c.getParsedResponse("GET", "/user/keys", nil, nil, &keys)
	return keys, err
}

// CreateKey creates a ssh key for authenticated user.
func (c *Client) CreateKey(opt AddSSHKeyOption) (*Key, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	key := new(Key)
	return key, c.getParsedResponse("POST", "/user/keys",
		http.Header{"content-type": []string{"application/json"}}, bytes.NewReader(body), key)
}

func (c *Client) DeleteKey(id int) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/user/keys/%s", id), nil)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	return err
}
