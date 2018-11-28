package libseat

func (s *libseatTestSuite) TestDiscordPlugin_GetDiscordMapping() {
	resp, err := s.c.GetDiscordMapping()
	if resp == nil || err != nil {
		s.FailNow("Unexpected GetGroup call failed", err.Error())
	}
}
