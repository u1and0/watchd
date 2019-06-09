/*
なにをやる
与えられたディレクトリ(なければカレントディレクトリ)のlsをフルパスで2秒毎に表示していく
<C-C>が押されるまで無限ループ

方向性
1. go langでtxtディレクトリ(以下txtDir)を監視
2. txtDirに新しいファイルが入ったら、そのファイル名を引数にpythonのグラフ化スクリプトを走らせる
> pythonのグラフ化だけじゃなくて汎用的に多くの外部コマンドを使いたい
> テスト用に、例えば `ls -ltr [filename]`

仕様例:
```
watchd [--dir dir] command
$ watchd --dir echo newTextFileList
```
dir指定がなければ実行ディレクトリ下
*/

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

func main() {
	/* thisfile, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(thisfile)
	*/
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// 引数が与えられていれば、そのディレクトリ
	if len(os.Args) > 1 {
		_dir := os.Args[1]
		f, err := os.Stat(_dir)
		if err != nil {
			log.Fatal(err)
		}
		if f.IsDir() {
			dir = _dir
		}
	}
	fmt.Println("////" + dir + "////")

	// シグナル待機中にやりたい処理 ※goroutine(並行処理)で書く
	go func() {
		for {
			files, err := filepath.Glob(dir + "/*")
			if err != nil {
				log.Fatal(err)
			}
			// /root以下のディレクトリ/ファイルを表示する
			for _, f := range files {
				fmt.Print(f)
			}
			fmt.Println("処理中...")
			time.Sleep(2 * time.Second)
		}
	}()

	// シグナル用のチャネル定義
	quit := make(chan os.Signal)
	// 受け取るシグナルを設定
	signal.Notify(quit, os.Interrupt)

	<-quit // ここでシグナルを受け取るまで以降の処理はされない
	// シグナルを受け取った後にしたい処理を書く
	fmt.Println("Keyboard interrupt!")
}
