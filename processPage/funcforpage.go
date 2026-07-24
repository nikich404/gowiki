package processPage

import (
	"bytes"
	"fmt"
	"gowiki/yaml_struct"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
)

func ProcessPage(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("dont open found!")
		return err

	}

	//
	var meta yaml_struct.PageMeta
	body, err := frontmatter.Parse(bytes.NewReader(data), &meta)
	if err != nil {
		fmt.Println("no parse")
		return err
	}

	//
	var buf bytes.Buffer
	if err := goldmark.Convert(body, &buf); err != nil {
		fmt.Println("not convert")
		return err
	}
	//if err := os.WriteFile("output.html", buf.Bytes(), 0644); err != nil {
	//	fmt.Println("not open file to HTML")
	//}

	//
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Println("template parse error")
		return err

	}

	temp := yaml_struct.PageData{
		Title: meta.Title,
		Date:  meta.Date,
		Body:  template.HTML(buf.String()),
	}

	name := strings.TrimSuffix(filepath.Base(path), ".md")
	outputPath := filepath.Join("output", name+".html")

	f, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("cannot create file")
		return err
	}
	defer f.Close()
	if err := tmpl.Execute(f, temp); err != nil {
		fmt.Println("template execute error")
		return err
	}
	return nil
}
