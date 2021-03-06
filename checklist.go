// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import "fmt"

type Checklist struct {
	client     *Client
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	IDBoard    string      `json:"idBoard,omitempty"`
	IDCard     string      `json:"idCard,omitempty"`
	Pos        float64     `json:"pos,omitempty"`
	CheckItems []CheckItem `json:"checkItems,omitempty"`
}

type CheckItem struct {
	client      *Client
	IDCard      string
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	State       string  `json:"state"`
	IDChecklist string  `json:"idChecklist,omitempty"`
	Pos         float64 `json:"pos,omitempty"`
}

// Manifestation of CheckItem when it appears in CheckItemStates
// on a Card.
type CheckItemState struct {
	IDCheckItem string `json:"idCheckItem"`
	State       string `json:"state"`
}

func (c *CheckItem) SetPos(newPos int) error {
	path := fmt.Sprintf("cards/%s/checklist/%s/checkItem/%s", c.IDCard, c.IDChecklist, c.ID)
	return c.client.Put(path, Arguments{"pos": fmt.Sprintf("%d", newPos)}, c)
}

func (c *CheckItem) SetNameAndState(name string, state string) error {
	path := fmt.Sprintf("cards/%s/checklist/%s/checkItem/%s", c.IDCard, c.IDChecklist, c.ID)
	return c.client.Put(path, Arguments{"name": name, "state": state}, c)
}

func (c *Checklist) AddCheckItem(name string) error {
	var checkitem CheckItem
	path := fmt.Sprintf("checklists/%s/checkItems", c.ID)
	args := Arguments{"name": name, "pos": "top", "checked": "false"}

	return c.client.Post(path, args, &checkitem)
}
