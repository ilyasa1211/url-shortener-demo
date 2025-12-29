package dto

type CreateSiteRequest struct {
	AliasUrl  string `json:"alias_url"`
	TargetUrl string `json:"target_url"`
}

type UpdateSiteRequest struct {
	TargetUrl string `json:"target_url`
}
