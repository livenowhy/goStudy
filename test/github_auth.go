package main

import (
	"fmt"
"github.com/livenowhy/dbModel/oauth"
)

func main() {

	repos, _, err := oauth.RepoList("94ab9a03687dbe24b1fc0504d29bc2884f86ba55")

	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(len(repos))

	for i, repo := range repos {
		fmt.Println(i)
		fmt.Println(*repo.FullName)
	}

}
