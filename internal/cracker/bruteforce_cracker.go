package cracker

import (
	"context"
	check "go-jwt-cracker/internal/check"
	"sync"
)

func BruteForceCracker(alg, token string) string {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	keys := GenerateAlnumKeys(ctx)
	for key := range keys {
		if check.CheckHMAC(alg, key, token) {
			return key
		}
	}

	return ""
}

// GenerateAlnumKeys 生成1-5位字母数字组合的密钥流
func GenerateAlnumKeys(ctx context.Context) <-chan string {
	charset := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	keys := make(chan string)
	var wg sync.WaitGroup

	// 创建goroutine分别处理1-5位所有组合
	for length := 1; length <= 5; length++ {
		wg.Add(1)
		go func(l int) {
			defer wg.Done()
			generateCombinations(ctx, keys, "", charset, l)
		}(length)
	}

	// 关闭通道
	go func() {
		wg.Wait()
		close(keys)
	}()

	return keys
}

// 递归生成所有组合
func generateCombinations(ctx context.Context, keys chan<- string, prefix, charset string, length int) {
	select {
	case <-ctx.Done():
		return
	default:
		if length == 0 {
			keys <- prefix
			return
		}
		for _, c := range charset {
			generateCombinations(ctx, keys, prefix+string(c), charset, length-1)
		}
	}
}
