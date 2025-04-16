# awsmyinfo

AWSユーザーが切り替え可能なアカウントとロールを表示するCLIツール

## 概要

`awsmyinfo`は、現在のIAMユーザーがスイッチできるAWSアカウントとそれに対応するロールを一覧表示するシンプルなCLIツールです。

## インストール

### Homebrewを使用する場合

```bash
brew tap juliar13/awsmyinfo
brew install awsmyinfo
```

### ソースからビルドする場合

```bash
git clone https://github.com/juliar13/awsmyinfo.git
cd awsmyinfo
go build -o awsmyinfo ./cmd/awsmyinfo
```

## 使い方

### 基本的な使い方

コマンドを引数なしで実行すると、現在のAWSプロファイルに設定されているユーザーの情報を取得します。

```bash
awsmyinfo
```

### 特定のユーザー名を指定する場合

ユーザー名を引数として指定することもできます。

```bash
awsmyinfo user-name
```

## 出力例

```
123456789012 ReadOnlySwitchRole
123456789012 AdminSwitchRole
123456789013 AdminSwitchRole
```

## 前提条件

- AWS CLIがインストールされ、プロファイルが設定されていること
- ユーザーが存在するアカウントのプロファイルに切り替えていること
- 適切なIAM権限があること（`iam:ListGroupsForUser`, `iam:ListAttachedUserPolicies`, `iam:ListAttachedGroupPolicies`, `iam:GetPolicy`, `iam:GetPolicyVersion`, `sts:GetCallerIdentity`）

## 注意事項

現在のバージョンはプロトタイプ段階であり、実際のポリシードキュメントからのロール情報の抽出は完全には実装されていません。将来のバージョンで改善される予定です。

## ライセンス

MIT
