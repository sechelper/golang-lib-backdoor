package backdoor

import (
	"os/exec"
)

func Calc() {
	// 执行计算器
	c := exec.Command("calc.exe")
	c.Run()
}
