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
	var err error
	switch command {
	case "reverse":
		err = reverseFile(args)
	case "copy":
		err = copyFile(args)
	case "duplicate-contents":
		err = duplicateContents(args)
	case "replace-string":
		err = replaceStringInFile(args)
	default:
		err = fmt.Errorf("コマンドは、[reverse|copy|duplicate-contents|replace-string]のいずれかを選択してください。")
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v", err)
		os.Exit(1)
	}
}

func validateArgsCount(args []string, count int, format string) error {
	if len(args) < count {
		return fmt.Errorf("引数を正しく指定してください。\nex) %s", format)
	}
	return nil
}

func getFileData(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ファイルの読み込みに失敗しました。%w", err)
	}
	return data, nil
}

func reverseFile(args []string) error {
	err := validateArgsCount(args, 4, "reverse <input file path> <output file path>")
	if err != nil {
		return err
	}
	input, output := args[2], args[3]

	data, err := getFileData(input)
	if err != nil {
		return err
	}

	slices.Reverse(data)
	err = os.WriteFile(output, data, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("%sの内容を反転させたものを%sに出力しました。\n", input, output)
	return nil
}

func copyFile(args []string) error {
	err := validateArgsCount(args, 4, "copy <input file path> <output file path>")
	if err != nil {
		return err
	}
	input, output := args[2], args[3]

	data, err := getFileData(input)
	if err != nil {
		return err
	}

	err = os.WriteFile(output, data, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("%sの内容をコピーして%sに出力しました。\n", input, output)
	return nil
}

func duplicateContents(args []string) error {
	err := validateArgsCount(args, 4, "duplicate-contents <input file path> <repeat count>")
	if err != nil {
		return err
	}
	input := args[2]
	repeatCount, err := strconv.Atoi(args[3])
	if err != nil {
		return fmt.Errorf("repeat countは整数の入力してください。")
	}

	data, err := getFileData(input)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(input, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	for i := 0; i < repeatCount; i++ {
		_, err := f.Write(data)
		if err != nil {
			return err
		}
	}

	fmt.Printf("%sの内容を%d回複製して%sに追記しました。\n", input, repeatCount, input)
	return nil
}

func replaceStringInFile(args []string) error {
	err := validateArgsCount(args, 5, "replace-string <input file path> <old string> <new string>")
	if err != nil {
		return err
	}
	input, oldStr, newStr := args[2], args[3], args[4]

	data, err := getFileData(input)
	if err != nil {
		return err
	}

	content := string(data)
	replaceContent := strings.ReplaceAll(content, oldStr, newStr)
	err = os.WriteFile(input, []byte(replaceContent), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("%sの内容から%sを検索して%sに置き換えました。\n", input, oldStr, newStr)
	return nil
}
