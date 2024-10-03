package username

func searchChess(username string) *string {
	return checkStatusAndReturnUrl("https://www.chess.com/member/"+username, 200)
}

func searchFortnite(username string) *string {
	return checkStatusAndReturnUrl("https://fortnitetracker.com/profile/all/"+username, 200)
}

func searchSpeedrun(username string) *string {
	return checkStatusAndReturnUrl("https://www.speedrun.com/user/"+username, 200)
}

func searchSteam(username string) *string {
	return checkStatusAndReturnUrl("https://steamcommunity.com/id/"+username, 200)
}
