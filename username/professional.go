package username

func searchGithub(username string) *string {
	returnUrl := "https://github.com/" + username
	return csAndReturnCustomURL("https://api.github.com/users/"+username, 200, &returnUrl)
}

func searchGitlab(username string) *string {
	return checkStatusAndReturnUrl("https://gitlab.com/"+username, 200)
}

func searchGitea(username string) *string {
	return checkStatusAndReturnUrl("https://gitea.com/"+username, 200)
}

func searchCodeberg(username string) *string {
	return checkStatusAndReturnUrl("https://codeberg.org/"+username, 200)
}

func searchCodecademy(username string) *string {
	return checkTextAndReturnUrl("https://www.codecademy.com/profiles/"+username, "&#x27;s avatar")
}

func searchCodeChef(username string) *string {
	return checkTextAndReturnUrl("https://www.codechef.com/users/"+username, "alt='user profile'")
}

func searchCodePen(username string) *string {
	return checkStatusAndReturnUrl("https://codepen.io/"+username, 200)
}

func searchDevCommunity(username string) *string {
	return checkStatusAndReturnUrl("https://dev.to/"+username, 200)
}

func searchCloudflareCommunity(username string) *string {
	return checkStatusAndReturnUrl("https://community.cloudflare.com/u/"+username+"/summary", 200)
}

func searchPyPi(username string) *string {
	return checkStatusAndReturnUrl("https://pypi.org/user/"+username, 200)
}

func searchNpm(username string) *string {
	return checkStatusAndReturnUrl("https://www.npmjs.com/~"+username, 200)
}

func searchReplit(username string) *string {
	return checkStatusAndReturnUrl("https://replit.com/@"+username, 200)
}
