package main

import (
	"flag"
	"fmt"
	cracker "go-jwt-cracker/internal/cracker"
	jwt "go-jwt-cracker/internal/jwt"
	"os"
)

func main() {
	// 定义命令行参数
	jwtToken := flag.String("jwt", "", "JWT token to crack")
	wordlist := flag.String("wordlist", "", "path to the wordlist file")
	mode := flag.String("mode", "", "crack mode")
	flag.Parse()

	// 参数检查
	if *jwtToken == "" || *mode == "" {
		fmt.Println("Usage: go-jwt-cracker --jwt <token> --mode <mode> [--wordlist <path>]")
		os.Exit(1)
	}

	// 根据mode调用不同的cracker进行爆破
	switch *mode {
	case "wordlist":
		if *wordlist == "" {
			fmt.Println("Please choose your wordlist path")
			os.Exit(1)
		}
		alg, err := jwt.ParseJWT(*jwtToken)
		if err != nil {
			fmt.Println("Error Parse JWT:", err)
			os.Exit(1)
		}

		result := cracker.WordlistCracker(alg, *jwtToken, *wordlist)
		if result != "" {
			fmt.Printf("Cracked secret: %s\n", result)
		} else {
			fmt.Println("Failed to crack the secret")
		}
	case "bruteforce":
		alg, err := jwt.ParseJWT(*jwtToken)
		if err != nil {
			fmt.Println("Error Parse JWT:", err)
			os.Exit(1)
		}

		result := cracker.BruteForceCracker(alg, *jwtToken)
		if result != "" {
			fmt.Printf("Cracked secret: %s\n", result)
		} else {
			fmt.Println("Failed to crack the secret")
		}

	default:
		fmt.Println("mode not exist")
		os.Exit(1)
	}
}
