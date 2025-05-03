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
		reverseFile(args)
	case "copy":
		copyFile(args)
	case "duplicate-contents":
		duplicateContents(args)
	case "replace-string":
		replaceStringInFile(args)
	default:
		fmt.Fprintln(os.Stderr, "コマンドは、[reverse|copy|duplicate-contents|replace-string]のいずれかを選択してください。")
		os.Exit(1)
	}

	os.Exit(0)
}

func validateArgsCount(args []string, count int, format string) error {
	if len(args) < count {
		return fmt.Errorf("引数を正しく指定してください。\nex) %s", format)
	}
	return nil
}

func getFileData(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラーが発生しました: %v\n", err)
		os.Exit(1)
	}
	return data
}

func reverseFile(args []string) {
	argsErr := validateArgsCount(args, 4, "reverse <input file path> <output file path>")
	if argsErr != nil {
		fmt.Fprintln(os.Stderr, argsErr)
		os.Exit(1)
	}
	input, output := args[2], args[3]

	data := getFileData(input)
	slices.Reverse(data)
	err := os.WriteFile(output, data, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラーが発生しました。: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%sの内容を反転させたものを%sに出力しました。\n", input, output)
}

func copyFile(args []string) {
	argsErr := validateArgsCount(args, 4, "copy <input file path> <output file path>")
	if argsErr != nil {
		fmt.Fprintln(os.Stderr, argsErr)
		os.Exit(1)
	}
	input, output := args[2], args[3]

	data := getFileData(input)
	err := os.WriteFile(output, data, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラーが発生しました。: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%sの内容をコピーして%sに出力しました。\n", input, output)
}

func duplicateContents(args []string) {
	argsErr := validateArgsCount(args, 4, "duplicate-contents <input file path> <repeat count>")
	if argsErr != nil {
		fmt.Fprintln(os.Stderr, argsErr)
		os.Exit(1)
	}
	input := args[2]
	repeatCount, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Fprintf(os.Stderr, "repeat countは整数を入力してください。: %v\n", err)
		os.Exit(1)
	}

	data := getFileData(input)
	f, err := os.OpenFile(input, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラーが発生しました: %v\n", err)
	}
	for i := 0; i < repeatCount; i++ {
		_, err := f.Write(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラーが発生しました: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("%sの内容を%d回複製して%sに追記しました。\n", input, repeatCount, input)
}

func replaceStringInFile(args []string) {
	argsErr := validateArgsCount(args, 5, "replace-string <input file path> <old string> <new string>")
	if argsErr != nil {
		fmt.Fprintln(os.Stderr, argsErr)
		os.Exit(1)
	}
	input, oldStr, newStr := args[2], args[3], args[4]

	data := getFileData(input)
	content := string(data)
	replaceContent := strings.ReplaceAll(content, oldStr, newStr)
	err := os.WriteFile(input, []byte(replaceContent), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラーが発生しました: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%sの内容から%sを検索して%sに置き換えました。\n", input, oldStr, newStr)
}
