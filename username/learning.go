package username

func searchDuoLingo(username string) *string {
	return checkStatusAndReturnUrl("https://www.duolingo.com/profile/"+username, 200)
}

func searchGeeksForGeeks(username string) *string {
	return checkStatusAndReturnUrl("https://auth.geeksforgeeks.org/user/"+username, 200)
}
