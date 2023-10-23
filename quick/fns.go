package quick

type ChatRoom struct {
	Usernames         []string
	TotalActionOnline int
	// Average Time Per user is a
	// product of usernames length divided
	// by TotalActionOnline
	AverageActionPerUser int
}

func GetAverageAction(c *ChatRoom) int {
	return c.TotalActionOnline / len(c.Usernames)
}

func (c *ChatRoom) CalculateAverageAction() {
	c.AverageActionPerUser = c.TotalActionOnline / len(c.Usernames)
}
