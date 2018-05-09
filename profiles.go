package goinsta

import (
	"encoding/json"
	"fmt"
)

// Profiles allows user function interactions
type Profiles struct {
	inst *Instagram
}

func newProfiles(inst *Instagram) *Profiles {
	profiles := &Profiles{
		inst: inst,
	}
	return profiles
}

// ByName return a *User structure parsed by username
func (prof *Profiles) ByName(name string) (*User, error) {
	body, err := prof.inst.sendSimpleRequest(urlUserByName, name)
	if err == nil {
		resp := userResp{}
		err = json.Unmarshal(body, &resp)
		// user is not nil at this point
		user := &resp.User
		user.inst = prof.inst
		return user, err
	}
	return nil, err
}

// ByID returns a *User structure parsed by user id
func (prof *Profiles) ByID(id int64) (*User, error) {
	data, err := prof.inst.prepareData()
	if err != nil {
		return nil, err
	}

	body, err := prof.inst.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlUserById, id),
			Query:    generateSignature(data),
		},
	)
	if err == nil {
		resp := userResp{}
		err = json.Unmarshal(body, &resp)
		user := &resp.User
		user.inst = prof.inst
		return user, err
	}
	return nil, err
}
