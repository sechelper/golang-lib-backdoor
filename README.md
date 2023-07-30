# backdoor

这是一个go语言包后门研究项目，在你的代码里引入这个项目，你就gg了

## 验证代码

```go
package main

import (
	_ "github.com/sechelper/backdoor"
)

func main() {

}
```