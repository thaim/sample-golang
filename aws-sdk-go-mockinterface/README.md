aws-sdk-go-v2 mockinterface
==============================
aws-sdk-go-v2においてmock interfaceを用いたモックを生成してユニットテストを実装するサンプルコード。

`make build` で実行可能なコードがビルドされる。
(対象S3バケットおよびオブジェクトキーはハードコードしているので修正が必要)。

このコードに対するテストは `make test`で動作する。