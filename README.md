与えられたディレクトリ(なければカレントディレクトリ)のlsをフルパスで2秒毎に表示していく
<C-C>が押されるまで無限ループ

方向性
1. go langでtxtディレクトリ(以下txtDir)を監視
2. txtDirに新しいファイルが入ったら、そのファイル名を引数にpythonのグラフ化スクリプトを走らせる
> pythonのグラフ化だけじゃなくて汎用的に多くの外部コマンドを使いたい
> テスト用に、例えば `ls -ltr [filename]`

使用例:
```
watchd --in dir --out dir command
$ watchd --in ./txt  --out ./pic python -c "textparse_exportpic.py"
```

