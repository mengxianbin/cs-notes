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

var contentDir = "content"
var ignoredDirReg = regexp.MustCompile(fmt.Sprintf("(\\.|%s)", contentDir))
var ignoredFileReg = regexp.MustCompile("(README|index)")

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

func writeMarkdown(parent string, fileName string, home string, parents *list.List) (string, error) {
	// 读取原始文件
	fi, err := os.Open(fmt.Sprintf("%s/%s", parent, fileName))
	if err != nil {
		log.Printf("Markdown reading error: %#v\n", err)
		panic(err)
	}
	defer func() { _ = fi.Close() }()
	fd, err := ioutil.ReadAll(fi)
	content := string(fd)

	// 创建多级目录
	newParent := fmt.Sprintf("./%s/%s", contentDir, parent[2:])
	err = os.MkdirAll(newParent, 0777)
	if err != nil {
		log.Printf("Directory making error: %#v\n", err)
		panic(err)
	}

	// 创建 markdown
	newPath := newParent + "/" + fileName
	file, err := os.Create(newPath)
	if err != nil {
		log.Printf("Markdown creating error: %#v\n", err)
		panic(err)
	}

	// 添加上级目录链接
	pathLink := toPathLink(home, parents)
	_, err = file.WriteString(pathLink)

	// 添加原始文件内容
	_, err = file.WriteString("\n\n")
	_, err = file.WriteString(content)

	// 关闭文件
	err = file.Close()

	return newPath, err
}

func toPathLink(home string, parents *list.List) string {
	parent := home
	link := fmt.Sprintf("[Home](%s) /", parent)
	for p := parents.Front(); p != nil; p = p.Next() {
		parent = fmt.Sprintf("%s/%s", parent, p.Value)
		link += fmt.Sprintf("\n[%s](%s) /", toTitle(p.Value.(string)), parent)
	}

	return link
}

// GenerateIndex 为目录递归生成 index.md
func GenerateIndex(path string, home string, parents *list.List) (err error) {
	// 创建 index.md
	indexFile, err := os.Create(path + "/index.md")
	if err != nil {
		log.Printf("Index creating error: %#v\n", err)
		panic(err)
	}

	// 添加上级目录链接
	pathLink := toPathLink(home, parents)
	_, err = indexFile.WriteString(pathLink)

	// 定义结束操作
	itemCount := 0
	defer func() {
		if itemCount == 0 {
			_, err = indexFile.WriteString("\n\n# TO DO")
		}

		_, err = indexFile.WriteString("\n")
		err = indexFile.Close()
		log.Printf("Index file closed: path=%s, itemCount=%d.", path, itemCount)
	}()

	// 读取文件列表
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Printf("Directory reading error: %#v\n", err)
		panic(err)
	}

	// 遍历文件名，追加到 index.md
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		var uri string
		if file.IsDir() { // 处理目录
			if ignoredDirReg.MatchString(file.Name()) {
				continue
			}

			parents.PushBack(file.Name())
			err = GenerateIndex(path+"/"+file.Name(), home, parents)
			parents.Remove(parents.Back())

			uri = "./" + file.Name()
		} else { // 处理文件
			if filepath.Ext(file.Name()) != ".md" || ignoredFileReg.MatchString(file.Name()) {
				continue
			}

			var mdPath string
			mdPath, err = writeMarkdown(path, file.Name(), home, parents)
			uri = fmt.Sprintf("%s/%s/%s", home, parents.Front().Value, mdPath[2:])
		}

		// 累计有效文件数
		itemCount++

		// 生成链接
		link := fmt.Sprintf("## [%s](%s)", toTitle(file.Name()), uri)
		log.Printf("Link: %s\n", link)

		// 写入链接
		_, err := indexFile.WriteString("\n\n" + link)
		if err != nil {
			err := indexFile.Close()
			log.Printf("Link writing error: %#v\n", err)
			panic(err)
		}
	}

	return err
}

func main() {
	parents := list.New()
	parents.PushBack("cs-note")
	_ = GenerateIndex(".", "https://mengxianbin.github.io", parents)
}
