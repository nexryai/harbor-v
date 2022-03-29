# Harbor-V
systemd-nspawn helper

## これは何
systemd-nspawnのヘルパー。具体的には... <br>
・コンテナの構築 <br>
・コンテナの初期設定 <br>
・コンテナのバックアップ、エクスポート、インポート <br>
<br>
を行う簡易的なスクリプト

## tips
・Harborはディストリビューション関係なくネットワーク関連の設定をsystemd-networkdで設定します。 <br>
・Harborはシンプルに扱えるように設計されています。バイナリ一つにまとまっており、複雑な依存関係はありません。

## 使い方
`./harbor [distribution-name].[distribution-version] [container_name] [username] [network_interface_name_for_container]` <br>

### 第1引数 [distribution-name].[distribution-version]
`debian.bullseye`や`ubuntu.focal`のように指定します。現在debianのみ対応しています。

### 第2引数 [container_name]
コンテナ名を指定します。スペースや日本語が入っていないかつ被っていないものなら何でも良いですが、あんまり打ちにくい名前にすると後で後悔しますよ。

### 第3引数 [username]
その名の通りコンテナ内で使うユーザーのユーザー名を指定します。

### 第4引数  [network_interface_name_for_container]
現在使用しているネットワークのインターフェース名を指定します。`em1`や`eth0`、`enp8s0`などです。
