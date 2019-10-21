package crowdin
import (
	"time"
)

// ListProjectBuilds - List Project Build API call
type ListProjectBuildsOptions struct {
	ProjectId 			int
	Body struct {
		BranchId	   	int	`json:hasManagerAccess,omitempty`
		Limit			int	`json:limit,omitempty`
		Offset			int	`json:offset,omitempty`
	}
}

type ResponseListProjectBuilds struct {
	Data []struct {
		Data struct {
			ID          int    `json:"id"`
			ProjectID   int    `json:"projectId"`
			BranchID    int    `json:"branchId"`
			LanguagesID []int  `json:"languagesId"`
			Status      string `json:"status"`
			Progress    struct {
				Percent           int `json:"percent"`
				CurrentLanguageID int `json:"currentLanguageId"`
				CurrentFileID     int `json:"currentFileId"`
			} `json:"progress"`
		} `json:"data"`
	} `json:"data"`
	Pagination []struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
}




// ListProjects - List Projects API call
type ListProjectsOptions struct {
	GroupId 						int	`json:groupId,omitempty`
	HasManagerAccess   	int	`json:hasManagerAccess,omitempty`
	Limit								int	`json:limit,omitempty`
	Offset							int	`json:offset,omitempty`
}

type ResponseListProjects struct {
	Data []struct {
		Data struct {
			Id                   int       `json:"id"`
			GroupId              int       `json:"groupId"`
			UserId               int       `json:"userId"`
			SourceLanguageId     int       `json:"sourceLanguageId"`
			TargetLanguageIds    []int     `json:"targetLanguageIds"`
			JoinPolicy           string    `json:"joinPolicy"`
			LanguageAccessPolicy string    `json:"languageAccessPolicy"`
			Type                 int       `json:"type"`
			Name                 string    `json:"name"`
			Cname                string    `json:"cname"`
			Identifier           string    `json:"identifier"`
			Description          string    `json:"description"`
			Visibility           string    `json:"visibility"`
			Logo                 []byte    `json:"logo"`
			Background           string    `json:"background"`
			IsExternal           bool      `json:"isExternal"`
			ExternalType         string    `json:"externalType"`
			AdvancedWorkflowId   int       `json:"advancedWorkflowId"`
			HasCrowdsourcing     bool      `json:"hasCrowdsourcing"`
			CreatedAt            time.Time `json:"createdAt"`
			UpdatedAt            time.Time `json:"updatedAt"`
		} `json:"data"`
	} `json:"data"`
	Pagination []struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
}




// GetProjectBuildProgressOptions are options for Check Project Build Status api call
type GetBuildProgressOptions struct {
	// Project Identifier.
	ProjectId int
	BuildId   int
}

type ResponseGetBuildProgress struct {
	Data struct {
		Id         int    `json:"id"`
		ProjectId  int    `json:"projectId"`
		BranchId   int    `json:"branchId"`
		LanguageId []int  `json:"languageId"`
		Status     string `json:"status"`
		Progress   struct {
			Percent           int `json:"percent"`
			CurrentLanguageId int `json:"currentLanguageId"`
			CurrentFileId     int `json:"currentFileId"`
		} `json:"progress"`
	} `json:"data"`
}


// DownloadProjectTranslationsOptions are options for  DownloadProjectTranslations api call
type DownloadProjectTranslationsOptions struct {
	// Project Identifier.
	ProjectId int
	// Build Identifier.
	BuildId int
}

type ResponseDownloadProjectTranslations struct {
  Data struct {
    Url 			string `json:"url"`
    ExpireIn 	string `json:"expireIn"`
  } `json: "data"`
}

// GetProjectBuildsOptions are options for GetProjectBuilds api call
type GetProjectBuildsOptions struct {
	// Project Identifier.
	ProjectId int
}

type ResponseGetProjectBuilds struct {
	Data []struct {
    Data struct {
      Id 					int 		`json:"id"`
			ProjectId 	int 		`json:"projectId"`
			BranchId 		int 		`json:"branchId"`
			LanguageId 	[]int		`json:"languageId"`
      Status 			string 	`json:"status"`
      Progress struct {
	        Percent 					int	`json:"percent"`
	        CurrentLanguageId int	`json:"currentLanguageId"`
			    CurrentFileId 		int	`json:"currentFileId"`
			} `json:"progress"`
    } `json:"data"`
  }  `json:"data"`
  Pagination []struct {
      Offset 			int	`json:"offset"`
      Limit 			int	`json:"limit"`
  } `json:"pagination"`
}


// GetLanguageProgressOptions are options for GetLanguageProgress api call
type GetLanguageProgressOptions struct {
	// Project Identifier.
	ProjectId int
}

type ResponseGetLanguageProgress struct {
	Data []struct {
		Data struct {
			LanguageId                int `json:"languageId"`
			PhrasesCount              int `json:"phrasesCount"`
			PhrasesTranslatedCount    int `json:"phrasesTranslatedCount"`
			PhrasesApprovedCount      int `json:"phrasesApprovedCount"`
			PhrasesTranslatedProgress int `json:"phrasesTranslatedProgress"`
			PhrasesApprovedProgress   int `json:"phrasesApprovedProgress"`
		} `json:"data"`
	} `json:"data"`
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
}

// BuildProjectOptions are options for BuildProject api call
type BuildProjectOptions struct {
	ProjectId int		 // Project Identifier.
	Body      struct {
		BranchId int `json:"branchId,omitempty"` // Branch Identifier. - optional
		// Specify target languages for build.
		// Leave this field empty to build all target languages
		Languages []string `json:"targetLanguagesId,omitempty"`
	}
}

type ResponseBuildProject struct {
	Data struct {
		Id          int      `json:"id"`
		ProjectId   int      `json:"ProjectId"`
		BranchId    int      `json:"branchId"`
		LanguagesId []string `json:"languagesId"`
		Status      string   `json:"status"`
		Progress    struct {
			Percent           int `json:"percent"`
			CurrentLanguageId int `json:"currentLanguageId"`
			CurrentFileId     int `json:"currentFileId"`
		} `json:"progress"`
	} `json:"data"`
}

type responseGeneral struct {
	Success bool `json:"success"`
}
