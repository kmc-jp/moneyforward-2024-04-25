# 全体共通の情報

## 導入

[INSTRUCTION_ja.md](./INSTRUCTION_ja.md)
[INSTRUCTION_en.md](./INSTRUCTION_ja.md)

## code-server でEbitengineを動かすためのTips

VCS (version control system) の準備が必要

```sh
git init
git config --global --add safe.directory /home/coder/work
```

環境変数が多すぎる問題のworkaroundとしてcleanenvを利用して起動する

```sh
go install github.com/agnivade/wasmbrowsertest/cmd/cleanenv@latest
PATH="$(go env GOPATH)/bin":$PATH
cleanenv -remove-prefix CODEBOX -- go run github.com/hajimehoshi/wasmserve@latest -http ":8000" .
```
