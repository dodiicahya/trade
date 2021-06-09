package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type UniqueStringResponse struct {
	FirstOccurence    string `json:"first_occurence"`
	LexicoGraphically string `json:"smallest_lexicographical_order"`
}

func UniqueString(ctx context.Context) error {
	var (
		res    = &UniqueStringResponse{}
		client = &http.Client{}
	)
	request, err := http.NewRequest("GET", "https://gist.githubusercontent.com/Jekiwijaya/0b85de3b9ff551a879896dd78256e9b8/raw/e9d58da5d4df913ad62e6e8dd83c936090ee6ef4/gistfile1.txt", nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer response.Body.Close()

	dataBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		e := fmt.Errorf("error read body : %v", err.Error())
		return e
	}

	bodyString := string(dataBody)

	res.FirstOccurence = FirstOccurence(bodyString)
	res.LexicoGraphically = LexicoGraphically(bodyString)
	resp, _ := json.Marshal(res)
	fmt.Println(string(resp))
	return nil
}

func FirstOccurence(s string) string {
	r := []rune(s)
	fmt.Println(r)
	r = RemoveDuplicate(r)
	return string(r)
}

func LexicoGraphically(s string) string {
	r := []rune(s)
	r = RemoveDuplicate(r)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func RemoveDuplicate(sliceData []rune) []rune {
	keys := make(map[rune]bool)
	list := []rune{}
	for _, entry := range sliceData {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
