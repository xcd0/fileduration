# fileduration

基本chatgpt製。

ファイルの「作成日時」「更新日時」およびその差分を表示するCLIツールです。  

## 🕊️ 概要

```sh
$ ./fileduration -h
Usage: fileduration.exe [--quiet] --search SEARCH [--basedir BASEDIR] [--exclude EXCLUDE]

Options:
  --quiet, -q            ヘッダーを非表示にする
  --search SEARCH, -s SEARCH
                         検索文字列(例: *.txt)
  --basedir BASEDIR, -d BASEDIR
                         検索ディレクトリ(省略時はカレントディレクトリ) [default: .]
  --exclude EXCLUDE, -v EXCLUDE
                         除外したいパターン(部分一致)。複数指定可
  --help, -h             display this help and exit
$ ./fileduration.exe -s "*" -v ".git"
作成    更新    差分    ファイル
2025-05-23 10:06:21     2025-05-23 10:09:11     00:02:50        /home/user/fileduration/go.mod
2025-05-23 10:06:55     2025-05-23 10:09:11     00:02:16        /home/user/fileduration/go.sum
2025-05-23 10:06:34     2025-05-23 10:34:41     00:28:07        /home/user/fileduration/main.go
2025-05-23 10:23:29     2025-05-23 10:28:44     00:05:14        /home/user/fileduration/readme.md
```

## 📦 特徴

- ファイルの作成日時・更新日時・差分（hh:mm:ss）を表示
- フルパスでの出力
- 検索パターン（例: "*.txt"）指定可能
- 検索基準ディレクトリの指定（省略時はカレントディレクトリ）
- 除外パターン（部分一致）指定可能
- ヘッダーの有無切替可能

## 🛠 使用例

```sh
# カレントディレクトリ以下から build_log_vs2019.txt を検索
$ ./fileduration -s "build_log_vs2019.txt"
作成               更新               差分       ファイル
2025-05-23 09:30:00	2025-05-23 11:00:00	01:30:00	/home/user/build_log_vs2019.txt
# ./logs 配下の .txt ファイルを検索し、ヘッダーを非表示
$ ./fileduration -q -s "*.txt" -d ./logs
2025-05-23 08:00:00	2025-05-23 08:10:30	00:10:30	/home/user/logs/log1.txt
2025-05-23 08:05:00	2025-05-23 08:07:00	00:02:00	/home/user/logs/log2.txt
# .log ファイルを検索し、tmpとbackupを含むパスを除外
$ ./fileduration -s "*.log" -v tmp -v backup
作成               更新               差分       ファイル
2025-05-23 07:00:00	2025-05-23 09:00:00	02:00:00	/home/user/logs/run.log
```

## ⚠ 注意

* 作成日時（BirthTime）はOSとファイルシステムに依存します。
	* WindowsやmacOSでは通常取得可能です。
	* Linux（ext4など）では取得できない場合があり、その場合はinodeの変更時刻（ctime）を代用しています。

## 📝 ライセンス

[CC0](https://creativecommons.org/publicdomain/zero/1.0/deed.ja)

