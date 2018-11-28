package libseat

type DiscordUser struct {
	SeatGroup int    `json:"group_id"`
	DiscordID int64  `json:"discord_id"`
	Nickname  string `json:"nick"`
}

// GetDiscordMapping returns a list of Discord Users along with their SeAT group ID
// See https://github.com/warlof/seat-discord-connector/pull/14
func (s *Client) GetDiscordMapping() ([]DiscordUser, error) {
	var resp []DiscordUser
	if err := s.performRequest(requestApi{
		Method: "GET",
		Path:   "/api/v2/discord-connector/mapping",
		Output: &resp,
	}); err != nil {
		return nil, err
	}

	return resp, nil
}
