# GO-CARDS 🃏

A simple card deck implementation in Go, demonstrating basic Go programming concepts.

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
