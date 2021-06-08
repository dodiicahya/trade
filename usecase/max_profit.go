package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Calculation for response
type Calculation struct {
	MaxProfit int `json:"max_profit"`
	BuyHour   int `json:"buy_hour"`
	SellHour  int `json:"sell_hour"`
}

func Max(ctx context.Context) error {
	var (
		client = &http.Client{}
	)

	request, err := http.NewRequest("GET", "https://gist.githubusercontent.com/Jekiwijaya/c72c2de532203965bf818e5a4e5e43e3/raw/2631344d08b044a4b833caeab8a42486b87cc19a/gistfile1.txt", nil)
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
		e := fmt.Errorf("error read body : %v", err.Error())
		return e
	}

	bodyString := string(dataBody)
	bodyArray := strings.Split(bodyString, " ")

	arr := make([]int, len(bodyArray))

	for i, v := range bodyArray {
		arr[i], err = strconv.Atoi(v)
		if err != nil {
			e := fmt.Errorf("error parsing data :%v", err.Error())
			return e
		}
	}
	calc := Calculate(arr)

	res, _ := json.Marshal(calc)
	fmt.Println(string(res))

	return nil
}

// Calculate for calculation profit
func Calculate(prices []int) *Calculation {
	calc := &Calculation{}
	// initiate temp values
	profit := 0
	buy_idx := 0
	sell_idx := 0
	low_idx := 0

	for i := 0; i < len(prices); i++ {

		if prices[i]-prices[low_idx] > profit {
			buy_idx = low_idx
			sell_idx = i
			profit = prices[i] - prices[buy_idx]
		}
		if prices[i] < prices[low_idx] {
			low_idx = i
		}
	}
	calc.MaxProfit = profit
	calc.BuyHour = buy_idx + 1
	calc.SellHour = sell_idx + 1

	return calc
}
