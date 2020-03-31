package lerps

import (
	// "fmt"

	"fmt"
	"testing"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/animation/tweening"
	"github.com/wdevore/RangerGo/engine/maths"

	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

func TestRunner(t *testing.T) {
	tweenTest()
}

func linearTest(t *testing.T) {
	equa := tweening.NewLinearEquation(api.EaseInOut)

	for ti := 0.0; ti < 1.0; ti += 0.1 {
		fmt.Println(equa.Compute(ti))
	}

}

func tweenTest() {
	tween := tweening.NewTween(-5.0, 10.0, 3.0, api.EquationExpo, api.EaseInOut)

	isFinished := false
	value := 0.0
	dt := 1.0 / 33.3333333 //0.167 // milliseconds per frame

	for !isFinished {
		value, isFinished = tween.Update(dt)
		fmt.Println("Value: ", value, ", Finshed: ", isFinished, ", Elapsed: ", tween.Elapsed())
	}
}

func expoTest() {
	fmt.Println("EaseIn -----------------")
	equa := tweening.NewExpoEquation(api.EaseIn)

	// Parameters
	duration := 3.0
	begin := -5.0
	end := 10.0

	// Controls
	dt := 0.167 // milliseconds per frame
	et := 0.0   // elapsed time
	interpolation := 0.0
	pInter := 0.0
	step := dt / duration
	dif := end - begin

	// Animation
	for t := 0.0; interpolation < 1.0; t += step {
		interpolation = equa.Compute(t)
		interpolation = maths.Clamp(interpolation, 0.0, 1.0)
		et += dt
		value := interpolation*dif + begin
		fmt.Println(interpolation, " : ", interpolation-pInter, " ", et, " v: ", value)
		pInter = interpolation
	}

	// fmt.Println("EaseOut -----------------")
	// equa = tweening.NewExpoEquation(api.EaseOut)

	// dt = 0.0
	// for ti := 0.0; ti <= 1.0; ti += 0.1 {
	// 	fmt.Println(equa.Compute(dt))
	// 	dt += 0.1
	// }

	// dt = 0.0
	// fmt.Println("EaseInOut -----------------")
	// equa = tweening.NewExpoEquation(api.EaseInOut)

	// for ti := 0.0; ti <= 1.0; ti += 0.1 {
	// 	fmt.Println(equa.Compute(dt))
	// 	dt += 0.1
	// }
}

func gweenTest() {
	duration := float32(3.0)
	var tween = gween.New(-5, 10, duration, ease.InExpo)

	dt := float32(0.167)
	isFinished := false
	current := float32(0)
	pCurrent := float32(0.0)
	et := float32(0.0)

	for !isFinished {
		current, isFinished = tween.Update(dt)
		et += dt
		fmt.Println(current, " : ", current-pCurrent, " ", et, " ", isFinished)
		pCurrent = current
	}
}
