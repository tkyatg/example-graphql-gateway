# example-graphql-grpc

## what is this repository?

- 仕事で使う参考用 repository として作成
  - api gateway として graphql、domain api として grpc を利用
  - 実装している処理
    - 認証
    - 認証したユーザーを取得する（jwt token から取得）

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
