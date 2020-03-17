package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var ignoreReg = regexp.MustCompile("(README|index)")

// 转换链接标题
func convertLinkTitle(source string) (target string) {
	// 文件名按下划线拆分
	parts := strings.Split(source, "_")

	// 单词首字母大写
	for i, part := range parts {
		parts[i] = strings.Title(part)
	}

	// 拼接链接标题
	subTitle := strings.Join(parts, " ")
	return fmt.Sprintf("## [%s](./%s)", subTitle, source)
}

func updateIndex(path string, parent *string) {
	// 创建 index.md
	indexFile, err := os.Create(path + "/index.md")
	if err != nil {
		log.Fatalf("Index creating error: %#v\n", err)
		return
	}

	// 读取文件列表
	files, err := ioutil.ReadDir(path)
	if err != nil {
		// empty directory
		indexFile.Close()
		return
	}

	// 添加上级目录链接
	if parent == nil {
		indexFile.WriteString("## [..](../index.md)\n\n")
	} else {
		indexFile.WriteString(*parent + "\n\n")
	}

	// 遍历文件名，追加到 index.md
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		fileName := file.Name()
		if file.IsDir() {
			// 处理子目录
			updateIndex(path+"/"+file.Name(), nil)
		} else {
			if filepath.Ext(file.Name()) != ".md" || ignoreReg.MatchString(file.Name()) {
				continue
			}
			// 截取文件名
			fileName = file.Name()[:len(fileName)-3]
		}

		// 链接名称转换
		link := convertLinkTitle(fileName)
		log.Printf("Link: %v\n", link)

		// 写入链接
		_, err := indexFile.WriteString(link + "\n\n")
		if err != nil {
			log.Fatalf("Link writing error: %#v\n", err)
			indexFile.Close()
			return
		}
	}

	// 关闭 index.md 文件
	indexFile.Close()

	log.Printf("Walking over: %v\n", path)
}

func main() {
	parent := "## [..](https://mengxianbin.github.io)"
	updateIndex(".", &parent)
}
