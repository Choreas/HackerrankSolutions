package implementation

//GetMoneySpent get max price for keyboard and drive, -1 for insufficient funds.
func GetMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var spend int32 = -1
	for _, keyboardCost := range keyboards {
		for _, driveCost := range drives {
			if keyboardCost+driveCost <= b && keyboardCost+driveCost > spend {
				spend = keyboardCost + driveCost
			}
		}
	}
	return spend
}
