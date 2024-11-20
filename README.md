# GO-CARDS 🃏

Go でのシンプルなカード デッキの実装。Go プログラミングの基本概念を示す。

## 🎯 Features

- Create a new deck of cards
- Shuffle cards
- Deal cards
- Save deck to file
- Load deck from file
- Print deck contents

## 📦 Build & Run

```
make build  # ビルドのみ
make run    # ビルドして実行
make clean  # バイナリと出力ファイルを削除
make test   # テストを実行
```

## 🚀 Directory Structure

```
go-cards/
├── bin/           # コンパイルされたバイナリの出力先
├── cmd/
│   └── main.go
├── internal/
│   └── deck/
│       ├── deck.go
│       └── deck_test.go
├── out/           # カードデータの出力先
├── go.mod
└── Makefile      # ビルドコマンドを簡略化
```
