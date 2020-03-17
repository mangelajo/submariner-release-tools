package github

import "strings"

func IsAlreadyExistError(err error) bool {

	//fmt.Println(err.(*goGithub.ErrorResponse))

	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), "already_exists")
}
