# example-grpc-api

## installation

1. GitHub のアクセストークンを取得する
   token の権限は `repo`, `read:packages` にチェックを入れる
2. 以下コマンドで初期化する

```
$ make init
```

3. app.env に GITHUB_TOKEN アクセストークンを書く

```
$ GITHUB_TOKEN=${GITHUB_TOKEN}
```

4. 以下コマンドで実行する

```
$ make serve
```
