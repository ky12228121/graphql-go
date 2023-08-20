# GraphQL用語解説

| 用語     | 説明                               |
| -------- | ---------------------------------- |
| Schema   | データ構造を定義したもの           |
| Model    | Schemaに従って生成されるGoの構造体 |
| Resolver | 実際にデータを取得する関数         |

# gqlgenの初期コード生成

## schemaの作成
プロジェクトルートに"schema.graphqls"というファイルを作成しスキーマ定義を書く。  
### 参考
公式: https://graphql.org/learn/schema/  
書き方: https://maku.blog/p/5s5jvfz/  
ベストプラクティス: https://maku.blog/p/4reqy9i/

## gqlgen設定ファイルの作成

プロジェクトルートに"gqlgen.yml"ファイルを作成し設定を記述する。

### 参考

公式: https://gqlgen.com/config/  

## コード生成

以下のコマンドを利用してGoのコードを生成する。
```zsh
gqlgen generate
```

## 自動生成ファイルの解説

- generated.go  
  **※編集してはいけないファイル**  
  リソルバをHTTP通信の橋渡しを行うコード。編集することはない。

- resolver.go  
  リソルバの型定義記述するファイル。  
  再生成されないためrepositoryやserviceの構造体ポインタを持たせる。

- schema.resolvers.go  
  リソルバ関数を実装するファイル。エンドポイント(クエリ、ミューテーション)の管理を行う。  
  コアロジックと接続することが多い。MVCのcontrollerなどの役割。

- models_gen.go  
  **※編集してはいけないファイル**  
  Schemaから定義されたGoの構造体。MVCでいうモデル定義。  
  modelパッケージ内の.goファイルはGraphQLと紐づけられる。

## エントリポイントの作成
通常は`server.go`という名前で作成する。goアプリケーションのエントリポイントとなり、HTTPサーバを開いてルーティングを行う。

# 開発の流れ

基本的にはResolverを作成していくだけ。
Schemaを変更した際は`gqlgen generate`コマンドでコードを再生成する。

## Service層の実装
DBとデータをやり取りするサービス層(package services)を作成する。  
ここは使用するDBによって実装が異なるが、ロジックの返り値はmodels_gen.goに定義されている各モデルのポインタにする。

## Resolverとサービス層の連携
schema.resolvers.goのコードは`gqlgen gemerate`コマンドで再生成されるため、ここにカスタムロジックを書くのはよろしくない。
resolver.goにあるResolver構造体に依存を注入する形で連携させる。

Resolver構造体にServiceの構造体を依存させ、エントリポイントのserver.goで生成したサービスのインスタンスを渡すことで各サービスを呼び出す。