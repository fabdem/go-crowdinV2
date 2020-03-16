package crowdin

import (
	"encoding/json"
	"fmt"
	"strconv"
)


// CheckProjectBuildStatus - Check Project Build Status api call
// {protocol}://{host}/api/v2/projects/{projectId}/translations/builds/{buildId}
func (crowdin *Crowdin) CheckProjectBuildStatus(options *CheckProjectBuildStatusOptions) (*ResponseCheckProjectBuildStatus, error) {

	crowdin.log("CheckProjectBuildStatus()")

	response, err := crowdin.get(&getOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"projects/%v/translations/builds/%v", crowdin.config.projectId, options.BuildId),
	})
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI ResponseCheckProjectBuildStatus
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.buildProgress = responseAPI.Data.Progress // Keep track % of progress

	return &responseAPI, nil
}



// BuildProjectTranslation - Build a project
// {protocol}://{host}/api/v2/projects/{ProjectId}/translations/builds
func (crowdin *Crowdin) BuildProjectTranslation(options *BuildProjectTranslationOptions) (*ResponseBuildProjectTranslation, error) {

	// Prepare URL and params
	var p postOptions
	p.urlStr = fmt.Sprintf(crowdin.config.apiBaseURL+"projects/%v/translations/builds", crowdin.config.projectId)
	p.body = options
	response, err := crowdin.post(&p)

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI ResponseBuildProjectTranslation
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
	}

	return &responseAPI, nil
}

// DownloadProjectTranslations - Download Project Translations api call
// {protocol}://{host}/api/v2/projects/{projectId}/translations/builds/{buildId}/download
func (crowdin *Crowdin) DownloadProjectTranslations(options *DownloadProjectTranslationsOptions) (*ResponseDownloadProjectTranslations, error) {

	response, err := crowdin.get(&getOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"projects/%v/translations/builds/%v/download", crowdin.config.projectId, options.BuildId),
	})

	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI ResponseDownloadProjectTranslations
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil
}

// ListProjectBuilds - List Project Builds API call. List the project builds
// {protocol}://{host}/api/v2/projects/{projectId}/translations/builds
func (crowdin *Crowdin) ListProjectBuilds(options *ListProjectBuildsOptions) (*ResponseListProjectBuilds, error) {

		var branchId string
		if options.BranchId >0 {
			branchId = strconv.Itoa(options.BranchId)
		}

		var limit string
		if options.Limit >0 {
			limit = strconv.Itoa(options.Limit)
		}

		var offset string
		if options.Offset >0 {
			offset = strconv.Itoa(options.Offset)
		}

	response, err := crowdin.get(&getOptions{
		urlStr: fmt.Sprintf(crowdin.config.apiBaseURL+"projects/%v/translations/builds", crowdin.config.projectId),
		params: map[string]string{
			"branchId"		: branchId,
			"limit"			: limit,
			"offset"		: offset,
		},
	})

	if err != nil {
		fmt.Printf("\nREPONSE:%s\n", response)
		crowdin.log(err)
		return nil, err
	}

	crowdin.log(string(response))

	var responseAPI ResponseListProjectBuilds
	err = json.Unmarshal(response, &responseAPI)
	if err != nil {
		crowdin.log(err)
		return nil, err
	}

	return &responseAPI, nil
}
