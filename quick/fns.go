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

	users := len(c.Usernames)

	if users == 0 {
		return 0
	}

	return c.TotalActionOnline / users
}

func (c *ChatRoom) CalculateAverageAction() {
	c.AverageActionPerUser = c.TotalActionOnline / len(c.Usernames)
}
