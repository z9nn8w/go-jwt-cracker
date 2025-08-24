package cracker

import (
	"bufio"
	"fmt"
	check "go-jwt-cracker/internal/check"
	"os"
	"strings"
)

func WordlistCracker(alg, token, wordlist string) string {
	// 打开字典文件
	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Println("Error opening wordlist:", err)
		return ""
	}
	defer file.Close()

	// 按行读取字典文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key := strings.TrimSpace(scanner.Text()) // 去掉空白字符
		if check.CheckHMAC(alg, key, token) {
			return key
		}
	}

	return ""
}
