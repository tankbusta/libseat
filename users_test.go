package libseat

func (s *libseatTestSuite) TestUsers_GetGroups() {
	resp, err := s.c.GetGroups()
	if resp == nil || err != nil {
		s.FailNow("Unexpected GetGroups call failed", err)
	}
}

func (s *libseatTestSuite) TestUsers_GetGroup() {
	resp, err := s.c.GetGroup(1)
	if resp == nil || err != nil {
		s.FailNow("Unexpected GetGroup call failed", err)
	}
}
