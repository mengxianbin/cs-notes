package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var ignoreReg = regexp.MustCompile("(README|index)")

// 将目录名转换为标题格式
func toTitle(source string) string {
	// 去除扩展名
	name := strings.ReplaceAll(source, ".md", "")

	// 文件名按下划线拆分
	parts := strings.Split(name, "_")

	// 单词首字母大写
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	// 拼接链接
	return strings.Join(parts, " ")
}

// GenerateIndex 为目录递归生成 index.md
func GenerateIndex(path string, home string, parents *list.List) {
	// 创建 index.md
	indexFile, err := os.Create(path + "/index.md")
	if err != nil {
		log.Fatalf("Index creating error: %#v\n", err)
	}

	// 添加上级目录链接
	parent := home
	indexFile.WriteString(fmt.Sprintf("[Home](%s) /", parent))
	for p := parents.Front(); p != nil; p = p.Next() {
		parent = fmt.Sprintf("%s/%s", parent, p.Value)
		indexFile.WriteString(fmt.Sprintf("\n[%s](%s) /", toTitle(p.Value.(string)), parent))
	}

	// 定义结束操作
	itemCount := 0
	finish := func(f *os.File, itemCount int) {
		if itemCount == 0 {
			f.WriteString("\n\n# TO DO")
		}

		f.WriteString("\n")
		f.Close()
	}

	// 读取文件列表
	files, err := ioutil.ReadDir(path)
	if err != nil {
		finish(indexFile, itemCount)
		log.Fatalf("Directory reading error: %#v\n", err)
	}

	// 遍历文件名，追加到 index.md
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		if file.IsDir() {
			// 处理子目录
			parents.PushBack(file.Name())
			GenerateIndex(path+"/"+file.Name(), home, parents)
			parents.Remove(parents.Back())
		} else {
			if filepath.Ext(file.Name()) != ".md" || ignoreReg.MatchString(file.Name()) {
				continue
			}
		}

		// 累计有效文件数
		itemCount++

		// 生成链接
		link := fmt.Sprintf("## [%s](./%s)", toTitle(file.Name()), file.Name())
		log.Printf("Link: %s\n", link)

		// 写入链接
		_, err := indexFile.WriteString("\n\n" + link)
		if err != nil {
			indexFile.Close()
			log.Fatalf("Link writing error: %#v\n", err)
		}
	}

	// 结束 index.md 写操作
	finish(indexFile, itemCount)
}

func main() {
	parents := list.New()
	parents.PushBack("cs-note")
	GenerateIndex(".", "https://mengxianbin.github.io", parents)
}
