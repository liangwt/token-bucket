### 简单令牌桶实现限流

Feature:
- 使用带缓冲通道作为桶实现
- 使用定时器触发令牌发放
- 线程安全可实现多并发

Example:

```go
package main

import (
	"github.com/liangwt/token-bucket"
	"log"
	"time"
)

func main() {
	bucket := bucket.New(10, time.Second, 100, true)
	isLimit := bucket.Consume(100)
	log.Println(isLimit)

	time.Sleep(2 * time.Second)
	isLimit = bucket.Consume(10)
	log.Println(isLimit)
}
```