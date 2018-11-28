package libseat

import "fmt"

type GroupUser struct {
	CharacterID int    `json:"character_id"`
	Active      bool   `json:"bool"`
	Name        string `json:"name"`
}

type Group struct {
	ID    int         `json:"id"`
	Users []GroupUser `json:"users"`
}

type GroupsResponse struct {
	Data []Group `json:"data"`
}

type GroupResponse struct {
	Data Group `json:"data"`
}

// GetGroups returns list of groups with their associated character_id's
func (s *Client) GetGroups() (*GroupsResponse, error) {
	var resp GroupsResponse
	if err := s.performRequest(requestApi{
		Method: "GET",
		Path:   "/api/v2/users/groups",
		Output: &resp,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGroup returns information about a specific character group
func (s *Client) GetGroup(id int) (*GroupResponse, error) {
	var resp GroupResponse
	if err := s.performRequest(requestApi{
		Method: "GET",
		Path:   fmt.Sprintf("/api/v2/users/groups/%d", id),
		Output: &resp,
	}); err != nil {
		return nil, err
	}
	return &resp, nil
}
