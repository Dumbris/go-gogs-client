// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

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

//TODO Add functions with getParsedResponse()
