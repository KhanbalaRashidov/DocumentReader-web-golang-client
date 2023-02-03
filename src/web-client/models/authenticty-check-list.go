package models

type AuthenticityCheckList struct {
	Count int
	List  []AuthenticityCheckResult
}

func NewAuthenticityCheckList(count int, list []AuthenticityCheckResult) *AuthenticityCheckList {
	return &AuthenticityCheckList{
		Count: count,
		List:  list,
	}
}
