package formula

import "math"

// CalculationBase 基础公式: 当前水位/目标水位 * 当前机器数 * (1 + 涨幅比) -- 向上取整
func CalculationBase(noWater float64, muWater float64, nowMachine int, increase float64) (foreMachine int) {
	result := noWater / muWater * float64(nowMachine) * (1 + increase)
	foreMachine = int(math.Floor(result + 0.5))
	return
}

// CalculationLevelSix 进阶公式: 当前水位/目标水位 * 当前机器数 * (1 + 涨幅比) 向上取整到6的倍数
func CalculationLevelSix(noWater float64, muWater float64, nowMachine int, increase float64) (foreMachine int) {
	result := noWater / muWater * float64(nowMachine) * (1 + increase)
	foreMachine = int(math.Floor(result + 0.5))
	if foreMachine%6 != 0 {
		foreMachine = foreMachine/6*6 + 6
	}
	return
}

// CalculationLevelThree 进阶公式: 当前水位/目标水位 * 当前机器数 * (1 + 涨幅比) 向上取整到3的倍数
func CalculationLevelThree(noWater float64, muWater float64, nowMachine int, increase float64) (foreMachine int) {
	result := noWater / muWater * float64(nowMachine) * (1 + increase)
	foreMachine = int(math.Floor(result + 0.5))
	if foreMachine%3 != 0 {
		foreMachine = foreMachine/3*3 + 3
	}
	return
}

// CalculationLevelSixNext 向上取整6倍数
func CalculationLevelSixNext(nowMachine int, increase float64) (foreMachine int) {
	result := float64(nowMachine) * increase
	foreMachine = int(math.Floor(result + 0.5))
	if foreMachine%6 != 0 {
		foreMachine = foreMachine/6*6 + 6
	}
	return
}

// CalculationLevelThreeNext 向上取整3倍数
func CalculationLevelThreeNext(nowMachine int, increase float64) (foreMachine int) {

	result := float64(nowMachine) * increase
	foreMachine = int(math.Floor(result + 0.5))
	if foreMachine%3 != 0 {
		foreMachine = foreMachine/3*3 + 3
	}
	return
}

// CheckThree 向上取整
func CheckThree(num int) (numA int) {
	numA = num
	for {
		if numA%3 == 0 {
			break
		}
		numA++
	}
	return
}

// CheckSix check six
func CheckSix(num int) (numA int) {
	numA = num
	for {
		if numA%6 == 0 {
			break
		}
		numA++
	}
	return
}

// GetUpInt get up int
func GetUpInt(num, check int) (result int) {
	result = num
	for {
		if result%check == 0 {
			break
		}
		result++
	}
	return
}
