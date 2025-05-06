# ファイル操作プログラム

このプログラムはGo言語で実装されたファイル操作ユーティリティです。テキストファイルの内容を様々な方法で操作することができます。

## 機能

このプログラムは以下の4つの機能を提供します：

1. **reverse**: 入力ファイルの内容を逆順にして出力ファイルに保存します
2. **copy**: 入力ファイルを別の場所にコピーします
3. **duplicate-contents**: 入力ファイルの内容を指定回数だけ複製して元のファイルに追記します
4. **replace-string**: ファイル内の特定の文字列を新しい文字列に置き換えます

## 使い方

```
go run file_manipulator.go [コマンド] [引数...]
```

各コマンドの使い方は以下の通りです：

### reverse

```
go run file_manipulator.go reverse [入力ファイルパス] [出力ファイルパス]
```

例：
```
go run file_manipulator.go reverse test.txt reversed.txt
```

### copy

```
go run file_manipulator.go copy [入力ファイルパス] [出力ファイルパス]
```

例：
```
go run file_manipulator.go copy test.txt test-copy.txt
```

### duplicate-contents

```
go run file_manipulator.go duplicate-contents [入力ファイルパス] [複製回数]
```

例：
```
go run file_manipulator.go duplicate-contents test.txt 3
```

### replace-string

```
go run file_manipulator.go replace-string [入力ファイルパス] [検索文字列] [置換文字列]
```

例：
```
go run file_manipulator.go replace-string test.txt "old" "new"
```

## バリデーション

このプログラムは引数の数や形式が正しいかをチェックするバリデーション機能を実装しています。引数が不足している場合や形式が間違っている場合は、適切なエラーメッセージが表示されます。

## ビルド方法

このプログラムを実行ファイルにビルドするには、以下のコマンドを実行します：

```
go build file_manipulator.go
```

ビルド後は、以下のように実行できます：

```
./file_manipulator [コマンド] [引数...]
``` 