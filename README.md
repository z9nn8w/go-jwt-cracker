# Go-JWT-Cracker

## Introduction

A JWT weak password brute force cracking and dictionary brute force cracking tool written in Go language, using Go's ARM feature to optimize JWT brute force cracking speed

Supported algorithms: **HS256,HS384,HS512**     (the program will automatically recognize)

## Usage

In the same directory as the program `go-jwt-cracker.exe` use command:
```
go-jwt-cracker --jwt <token> --mode <mode> [--wordlist <path>]
```

Parameter description:
```
token: JWT token
mode: cracker mode (bruteforce or wordlist) 
path: the path of wordlist file
```

## Function

`bruteforce` mode is used to bruteforce a key consisting of __1-5__ digits and a mixture of uppercase and lowercase letters. (e.g. 1a39P)

`wordlist` mode is used to extract the key from the dictionary file line by line and bruteforce it.

## Example

we have a JWT token like:
`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoic3hjIiwiYWRtaW4iOnRydWV9.l1hoVX-GMZ1D1Ug-MyHCaZsE5bRKsi2OL3cO4pFtr1I`

if we use wordlist mode:
```cmd
go-jwt-cracker --jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoic3hjIiwiYWRtaW4iOnRydWV9.l1hoVX-GMZ1D1Ug-MyHCaZsE5bRKsi2OL3cO4pFtr1I --mode wordlist --wordlist dicts/rockyou.txt
```

result:
```
Header: {"alg":"HS256","typ":"JWT"}
Payload: {"name":"sxc","admin":true}
Cracked secret: 12345
```

if we use bruteforce mode:
```
go-jwt-cracker.exe --jwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoic3hjIiwiYWRtaW4iOnRydWV9.l1hoVX-GMZ1D1Ug-MyHCaZsE5bRKsi2OL3cO4pFtr1I --mode bruteforce
```

result:
```
Header: {"alg":"HS256","typ":"JWT"}
Payload: {"name":"sxc","admin":true}
Cracked secret: 12345
```

## Future

In the future, I will consider adding features such as blasting progress bar and time consumption