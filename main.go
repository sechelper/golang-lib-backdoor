package backdoor

import (
	"os/exec"
)

func Calc() {
	// 执行计算器
	exec.Command("calc.exe")
}
