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
