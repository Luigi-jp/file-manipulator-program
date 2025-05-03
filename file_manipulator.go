package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "引数を指定してください。\n- reverse <input file path> <output file path>\n- copy <input file path> <output file path>\n-duplicate-contents <input file path> <repeat count>\n- replace-string <input file path> <old string> <new string>")
		os.Exit(1)
	}

	command := args[1]
	switch command {
	case "reverse":
		// 引数のチェック
		// 引数フォーマット: reverse <input file path> <output file path>
		if len(args) < 4 {
			fmt.Println("引数が不足しています。\nex) reverse <input file path> <output file path>")
			os.Exit(1)
		}
		input := args[2]
		output := args[3]

		// input file pathの内容を取得
		data, err := os.ReadFile(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラーが発生しました: %v\n", err)
			os.Exit(1)
		}

		// input file pathの内容を反転させる
		slices.Reverse(data)

		// input file pathの内容を反転させた内容をoutput file pathに出力する
		// output file pathがなければファイルを作成。あれば上書きをする
		err = os.WriteFile(output, data, 0644)

		if err != nil {
			fmt.Fprintf(os.Stderr, "エラーが発生しました。: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%sの内容を反転させたものを%sに出力しました。\n", input, output)
	case "copy":
		// 引数チェック
		// 引数フォーマット： copy <input file path> <output file path>
		if len(args) < 4 {
			fmt.Fprintln(os.Stderr, "引数が不足しています。\nex) copy <input file path> <output file path>")
			os.Exit(1)
		}
		input := args[2]
		output := args[3]

		// input file pathの内容を取得
		data, err := os.ReadFile(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラーが発生しました: %v", err)
			os.Exit(1)
		}

		// input file pathの内容をコピーして、output file pathに出力する
		// output file pathがなければファイルを作成。あれば上書きする
		err = os.WriteFile(output, data, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラーが発生しました。: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%sの内容をコピーして%sに出力しました。\n", input, output)
	case "duplicate-contents":
		// 引数チェック
		// 引数フォーマット： duplicate-contents <input file path> <repeat count>
		if len(args) < 4 {
			fmt.Fprintln(os.Stderr, "引数が不足しています。\nex) duplicate-contents <input file path> <repeat count>です。")
			os.Exit(1)
		}

		// 繰り返し回数（repeat count）が数値で入力されているかチェック
		input := args[2]
		repeatCount, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "repeat countは整数を入力してください。: %v\n", err)
			os.Exit(1)
		}

		// input file pathの内容を取得
		data, err := os.ReadFile(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラーが発生しました: %v\n", err)
			os.Exit(1)
		}

		// input file pathの内容を{repeat count}回複製して、input file pathに追記する
		f, err := os.OpenFile(input, os.O_APPEND|os.O_WRONLY, 0644)
		defer f.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラーが発生しました: %v\n", err)
		}
		for i := 0; i < repeatCount; i++ {
			_, err := f.Write(data)
			if err != nil {
				fmt.Fprintf(os.Stderr, "エラーが発生しました: %v\n", err)
				os.Exit(0)
			}
		}

		fmt.Printf("%sの内容を%d回複製して%sに追記しました。\n", input, repeatCount, input)
	case "replace-string":
		// 引数チェック
		// 引数フォーマット: replace-string <input file path> <old string> <new string>
		if len(args) < 5 {
			fmt.Fprintln(os.Stderr, "引数が不足しています。\nex) replace-string <input file path> <old string> <new string>")
			os.Exit(1)
		}

		input := args[2]
		oldStr := args[3]
		newStr := args[4]

		// input file pathの内容を取得
		data, err := os.ReadFile(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラーが発生しました: %v\n", err)
			os.Exit(1)
		}

		// input file pathの内容から文字列{oldStr}を検索し、{oldStr}を全て{newStr}に変換する
		content := string(data)
		replaceContent := strings.ReplaceAll(content, oldStr, newStr)

		// 変換した内容でinput file pathを置き換える
		err = os.WriteFile(input, []byte(replaceContent), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラーが発生しました: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%sの内容から%sを検索して%sに置き換えました。\n", input, oldStr, newStr)
	default:
		fmt.Fprintln(os.Stderr, "コマンドは、[reverse|copy|duplicate-contents|replace-string]のいずれかを選択してください。")
		os.Exit(1)
	}

	os.Exit(0)
}
