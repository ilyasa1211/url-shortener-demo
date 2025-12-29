package domain

type Site struct {
	ID        int    `json:"id"`
	AliasUrl  string `json:"alias_url"`
	TargetUrl string `json:"target_url"`
}
