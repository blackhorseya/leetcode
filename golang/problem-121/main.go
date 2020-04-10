package main

func Solve(prices []int) int {
	if prices == nil || len(prices) <= 1 {
		return 0
	}

	buyFlag := 0
	profit := 0
	for index, price := range prices[1:] {
		if prices[index] > price && price < prices[buyFlag] {
			buyFlag = index + 1
		}

		if price-prices[buyFlag] > profit {
			profit = price - prices[buyFlag]
		}
	}

	return profit
}

func main() {

}
