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

	mapset "github.com/deckarep/golang-set"
)

// basename : ファイルパスpathからディレクトリと拡張子を取り去る
func basename(path string) string {
	filename := filepath.Base(path)
	ext := filepath.Ext(filename)
	base := filename[0 : len(filename)-len(ext)]
	return base
}

// dirとdir1には行っているファイルの差分を2秒ごとに表示
func main() {
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

	dir = "test"
	dir1 := "test1"

	fmt.Println("////" + dir + "////")

	// シグナル待機中にやりたい処理 * goroutine(並行処理)で書く
	go func() {
		for {
			files, err := filepath.Glob(dir + "/*")
			if err != nil {
				log.Fatal(err)
			}
			// ディレクトリのファイル郡をset化
			fileset := mapset.NewSet()
			for _, f := range files {
				fileset.Add(basename(f))
			}

			files1, err := filepath.Glob(dir1 + "/*")
			if err != nil {
				log.Fatal(err)
			}
			// ディレクトリのファイル郡をset化
			fileset1 := mapset.NewSet()
			for _, f := range files1 {
				fileset1.Add(basename(f))
			}

			// /dirとdir1以下のディレクトリ/ファイルを表示する
			fmt.Print(fileset)
			fmt.Println("処理中...")
			fmt.Print(fileset1)
			fmt.Println("処理中1...")
			// 差分表示
			fmt.Print(fileset.Difference(fileset1))
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
