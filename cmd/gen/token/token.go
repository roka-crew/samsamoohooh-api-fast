package token

import (
	"flag"
	"fmt"
	"os"

	"github.com/roka-crew/pkg/config"
	"github.com/roka-crew/pkg/token"
)

var userID = flag.Uint("user_id", 0, "사용자의 ID 값을 입력해주세요.")
var secretKey = flag.String("secret_key", "", "JWT 토큰을 생성하기 위한 시크릿 키를 입력해주세요.")

func Run() {
	err := flag.CommandLine.Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		return
	}
	flag.Parse()

	cfg := &config.Config{
		Token: config.Token{
			SecretKey: []byte(*secretKey),
		},
	}

	tokenService := token.New(cfg)
	tokenString, err := tokenService.GenerateToken(*userID)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Generated:", tokenString)
}
