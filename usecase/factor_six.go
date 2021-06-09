package usecase

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FactorRequest struct {
	Number int `json:"number"`
}

type FactorResponse struct {
	MatchedFactorCount int `json:"mathched_factor_count"`
}

func FactorSixCount(ctx context.Context) error {

	request := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := request.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	response := &FactorResponse{}
	response.MatchedFactorCount, _ = FindFactorSix(text)
	res, _ := json.Marshal(response)
	fmt.Println(string(res))
	return nil
}

func isMatchFactor(n int, match int) bool {
	count := 0

	for t := n; t > 0; t-- {
		if (n/t)*t == n {
			count++
		}
		if count > match {
			return false
		}
	}
	if count == match {
		return true
	}
	return false
}

func FindFactorSix(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	result := 0
	matchFactor := 6
	for i := 1; i <= num; i++ {
		isMatch := isMatchFactor(i, matchFactor)
		if isMatch {
			result++
		}
	}
	return result, nil
}
