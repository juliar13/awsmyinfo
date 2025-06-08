package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/juliar13/awsinfo/pkg/aws"
)

const version = "0.2.0"

func main() {
	// フラグの設定
	var showVersion = flag.Bool("version", false, "バージョン情報を表示")
	flag.Parse()

	// バージョン情報の表示
	if *showVersion {
		fmt.Printf("awsinfo version %s\n", version)
		return
	}

	// コンテキストの作成
	ctx := context.Background()

	// 引数の解析
	args := flag.Args()
	var userName string
	var err error

	// 引数が指定されている場合はそれを使用し、それ以外は現在のユーザー名を取得
	if len(args) > 0 {
		userName = args[0]
	} else {
		userName, err = aws.GetCurrentUserName(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("ユーザー名: %s\n", userName)

	// スイッチロール情報を取得
	roles, err := aws.GetSwitchRoleInfo(ctx, userName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}

	// 結果の表示
	for _, role := range roles {
		fmt.Printf("%s %s\n", role.AccountID, role.RoleName)
	}

	// 結果がない場合でも、アカウントとロールの組み合わせが取得できない旨を表示
	if len(roles) == 0 {
		fmt.Println("スイッチ可能なアカウントとロールが見つかりませんでした。")
		fmt.Println("注: このツールは現在プロトタイプ段階です。実際のポリシードキュメントの解析は実装されていません。")
		
		// デモ用にサンプルデータを表示
		fmt.Println("\nデモ用のサンプル出力:")
		fmt.Println("123456789012 ReadOnlySwitchRole")
		fmt.Println("123456789012 AdminSwitchRole")
		fmt.Println("123456789013 AdminSwitchRole")
	}
}
