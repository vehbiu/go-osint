package username

import (
	"fmt"
	"github.com/fatih/color"
	"sync"
)

var platforms = []string{
	/* Learning */
	"duolingo",
	"geeksforgeeks",

	/* Professional */
	"github",
	"gitlab",
	"gitea",
	"codeberg",
	"codecademy",
	"codechef",
	"codepen",
	"devcommunity",
	"cloudflarecommunity",
	"pypi",
	"npm",
	"replit",

	/* Games */
	"chess",
	"fortnite",
	"speedrun",
	"steam",

	/* Social */
	"9gag",
	"aboutme",
	"askfedora",
	"bitwardenforum",
	"bravecommunity",
	"buymeacoffee",
	"buzzfeed",
	"cnet",
	"cartalkcommunity",
	"dockerhub",
	"pastebin",
	"reddit",
	"rumble",
	"scratch",
	"snapchat",
}

type SearchResult struct {
	Platform string
	Output   *string
}

func searchPlatform(username, platform string) *string {
	var output *string
	switch platform {
	case "duolingo":
		output = searchDuoLingo(username)
	case "geeksforgeeks":
		output = searchGeeksForGeeks(username)
	case "github":
		output = searchGithub(username)
	case "gitlab":
		output = searchGitlab(username)
	case "gitea":
		output = searchGitea(username)
	case "codeberg":
		output = searchCodeberg(username)
	case "codecademy":
		output = searchCodecademy(username)
	case "codechef":
		output = searchCodeChef(username)
	case "codepen":
		output = searchCodePen(username)
	case "devcommunity":
		output = searchDevCommunity(username)
	case "cloudflarecommunity":
		output = searchCloudflareCommunity(username)
	case "pypi":
		output = searchPyPi(username)
	case "npm":
		output = searchNpm(username)
	case "replit":
		output = searchReplit(username)
	case "chess":
		output = searchChess(username)
	case "fortnite":
		output = searchFortnite(username)
	case "speedrun":
		output = searchSpeedrun(username)
	case "steam":
		output = searchSteam(username)
	case "9gag":
		output = search9Gag(username)
	case "aboutme":
		output = searchAboutMe(username)
	case "askfedora":
		output = searchAskFedora(username)
	case "bitwardenforum":
		output = searchBitwardenForum(username)
	case "bravecommunity":
		output = searchBraveCommunity(username)
	case "buymeacoffee":
		output = searchBuyMeACoffee(username)
	case "buzzfeed":
		output = searchBuzFeed(username)
	case "cnet":
		output = searchCNet(username)
	case "cartalkcommunity":
		output = searchCarTalkCommunity(username)
	case "dockerhub":
		output = searchDockerHub(username)
	case "pastebin":
		output = searchPasteBin(username)
	case "reddit":
		output = searchReddit(username)
	case "rumble":
		output = searchRumble(username)
	case "scratch":
		output = searchScratch(username)
	case "snapchat":
		output = searchSnapchat(username)
	default:
		fmt.Println(color.RedString("invalid platform! must be one of: %s", platforms))
	}

	return output
}

func Search(username, platform string) *string {
	output := searchPlatform(username, platform)
	if output != nil {
		fmt.Println(color.GreenString("[✔]"), color.YellowString(platform), *output)
		return output
	} else {
		fmt.Println(
			color.RedString("[✖]"),
			color.YellowString(platform),
			"not found",
		)
		return nil
	}
}

func SearchAll(username string) *[]string {
	resultsChan := make(chan SearchResult, len(platforms))
	var wg sync.WaitGroup

	// Launch goroutines for each platform
	for _, platform := range platforms {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			output := searchPlatform(username, p)
			resultsChan <- SearchResult{Platform: p, Output: output}
		}(platform)
	}

	// Close channel when all goroutines are done
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	items := make([]string, len(resultsChan))

	// Print results as they come in
	for result := range resultsChan {
		if result.Output != nil {
			items = append(items, *result.Output)
			fmt.Println(color.GreenString("[✔]"), color.YellowString(result.Platform), *result.Output)
		} else {
			fmt.Println(
				color.RedString("[✖]"),
				color.YellowString(result.Platform),
				"not found",
			)
		}
	}

	if len(items) == 0 {
		return nil
	}

	return &items
}
