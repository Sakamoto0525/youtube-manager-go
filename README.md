# youtube-manager-go
Go × Nuxt 動画投稿アプリ作成チュートリアル

## セットアップ

### API

```bash:
# apiディレクトリに移動
cd api
# depがインストールされていなければインストール
go get -u github.com/golang/dep/cmd/dep
# or
brew install dep

# depで依存関係をインストール（yarnみたいな感じ）
cd api/src/project
dep ensure
```

### Front

```bash:
# frontディレクトリに移動
cd front
npm install
```

## 起動

### API

```bash:
cd api/src/project
go run server.go

# ホットリロード
cd api/src/project
../../bin/realize start --run --no-config
```

### Front

```bash:
cd front
npm run dev
```

## メモ

### Goのimportについて

Goのimportは、GOPATHからの相対パスで指定する
相対パスは.envrcで設定している

今回でいうと`../youtube-manager-go/api`と設定しているので、
`youtube-manager-go/api/src/project/sample.go`をimportしたい場合は、
`import "project/sample"`とする。

相対パスを変更する場合は、.envrcのパスを書き換え、以下のコマンドを叩く

```bash:
$ direnv allow
```
