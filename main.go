package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/djherbis/times"
)

type args struct {
	Quiet   bool     `arg:"-q,--quiet" help:"ヘッダーを非表示にする"`
	Search  string   `arg:"-s,required" help:"検索文字列(例: *.txt)"`
	BaseDir string   `arg:"-d" default:"." help:"検索ディレクトリ(省略時はカレントディレクトリ)"`
	Exclude []string `arg:"-v,separate" help:"除外したいパターン(部分一致)。複数指定可"`
}

func init() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {
	var args args
	arg.MustParse(&args)

	matches, err := findMatchingFiles(args.BaseDir, args.Search, args.Exclude)
	if err != nil {
		log.Fatalf("ファイル検索エラー: %v", err)
	}

	if !args.Quiet {
		printHeader()
	}

	for _, path := range matches {
		printDurationLine(path)
	}
}

func printHeader() {
	fmt.Println("作成\t更新\t差分\tファイル")
}

func findMatchingFiles(baseDir, pattern string, exclude []string) ([]string, error) {
	var results []string
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if match, _ := filepath.Match(pattern, filepath.Base(path)); !match {
			return nil
		}
		for _, ex := range exclude {
			if strings.Contains(path, ex) {
				return nil // 除外
			}
		}
		abs, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		results = append(results, abs)
		return nil
	})
	return results, err
}

func printDurationLine(filePath string) {
	t, err := times.Stat(filePath)
	if err != nil {
		log.Printf("ファイル取得失敗: %s (%v)", filePath, err)
		return
	}

	var ctime time.Time
	if t.HasBirthTime() {
		ctime = t.BirthTime()
	} else {
		ctime = t.ChangeTime()
	}

	mtime := t.ModTime()
	diff := mtime.Sub(ctime)

	h := int(diff.Hours())
	m := int(diff.Minutes()) % 60
	s := int(diff.Seconds()) % 60
	diffStr := fmt.Sprintf("%02d:%02d:%02d", h, m, s)

	fmt.Printf("%s\t%s\t%s\t%s\n",
		ctime.Format("2006-01-02 15:04:05"),
		mtime.Format("2006-01-02 15:04:05"),
		diffStr,
		filePath)
}
