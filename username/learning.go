package username

func searchDuoLingo(username string) *string {
	url := "https://www.duolingo.com/profile/" + username
	return ctAndReturnCustomURL("https://www.duolingo.com/2017-06-30/users?fields=users%7Bid%7D&username="+username, "{\"users\":[{\"id\":", &url)
}

func searchGeeksForGeeks(username string) *string {
	return checkStatusAndReturnUrl("https://auth.geeksforgeeks.org/user/"+username, 200)
}
