package username

func search9Gag(username string) *string {
	return checkStatusAndReturnUrl("https://www.9gag.com/u/"+username, 200)
}

func searchAboutMe(username string) *string {
	return checkStatusAndReturnUrl("https://about.me/"+username, 200)
}

func searchAskFedora(username string) *string {
	return checkStatusAndReturnUrl("https://discussion.fedoraproject.org/u/"+username+"/summary", 200)
}

func searchBitwardenForum(username string) *string {
	return checkStatusAndReturnUrl("https://community.bitwarden.com/u/"+username, 200)
}

func searchBraveCommunity(username string) *string {
	return checkStatusAndReturnUrl("https://community.brave.com/u/"+username, 200)
}

func searchBuyMeACoffee(username string) *string {
	return checkStatusAndReturnUrl("https://buymeacoff.ee/"+username, 200)
}

func searchBuzFeed(username string) *string {
	return checkStatusAndReturnUrl("https://www.buzzfeed.com/"+username, 200)
}

func searchCNet(username string) *string {
	return checkStatusAndReturnUrl("https://www.cnet.com/profiles/"+username+"/", 200)
}

func searchCarTalkCommunity(username string) *string {
	return checkStatusAndReturnUrl("https://community.cartalk.com/u/"+username+"/summary", 200)
}

func searchDockerHub(username string) *string {
	respUrl := "https://hub.docker.com/u/" + username
	return csAndReturnCustomURL("https://hub.docker.com/v2/users/"+username, 200, &respUrl)
	//return checkStatusAndReturnUrl("https://hub.docker.com/u/"+username, 200)
}

func searchPasteBin(username string) *string {
	return checkStatusAndReturnUrl("https://pastebin.com/u/"+username, 200)
}

func searchReddit(username string) *string {
	retUrl := "https://www.reddit.com/user/" + username
	return csAndReturnCustomURL("https://old.reddit.com/user/"+username, 200, &retUrl)
	//return checkStatusAndReturnUrl("https://www.reddit.com/user/"+username, 200)
}

func searchRumble(username string) *string {
	return checkStatusAndReturnUrl("https://rumble.com/user/"+username, 200)
}

func searchScratch(username string) *string {
	return checkStatusAndReturnUrl("https://scratch.mit.edu/users/"+username, 200)
}

func searchSnapchat(username string) *string {
	return checkStatusAndReturnUrl("https://www.snapchat.com/add/"+username, 200)
}
