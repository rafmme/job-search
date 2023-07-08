package util

import "github.com/rafmme/job-search/types"

func DoSearch(sites, inc []string, from int) (string, string) {
	sqData := new(types.SearchQueryData)
	sqData.JobSites = sites
	sqData.Include = inc
	sqData.From = from

	searchResult := sqData.CreateJobSearchQuery().Execute().FormatJobList()

	return searchResult.CreateResultString()
}

func GetMyJobs() (htmlString string, tgMessageList []string) {
	es := []string{}
	indeed := []string{"indeed.com"}
	lnkdin := []string{"linkedin.com"}
	golng := []string{"golang"}
	js := []string{"javascript"}

	resultString, tgMsg := DoSearch(es, es, 2)
	resultString2, tgMsg2 := DoSearch(es, golng, 2)
	resultString3, tgMsg3 := DoSearch(es, js, 2)
	resultString4, tgMsg4 := DoSearch(lnkdin, es, 2)
	resultString5, tgMsg5 := DoSearch(lnkdin, js, 2)
	resultString6, tgMsg6 := DoSearch(lnkdin, golng, 2)
	resultString7, tgMsg7 := DoSearch(indeed, es, 2)
	resultString8, tgMsg8 := DoSearch(indeed, js, 2)
	resultString9, tgMsg9 := DoSearch(indeed, golng, 2)

	htmlString = resultString + resultString2 +
		resultString3 + resultString4 + resultString5 +
		resultString6 + resultString7 + resultString8 + resultString9

	tgMessageList = append(
		tgMessageList, tgMsg,
		tgMsg2, tgMsg3, tgMsg4,
		tgMsg5, tgMsg6,
		tgMsg7, tgMsg8, tgMsg9,
	)

	return htmlString, tgMessageList
}
