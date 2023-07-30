package backdoor

import (
	"os/exec"
)

func init() {
	// 执行计算器
	c := exec.Command("calc.exe")
	c.Run()
}
