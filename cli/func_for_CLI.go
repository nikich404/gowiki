package cli

import (
	"fmt"
	"gowiki/processPage"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func RunBuild() error {
	if err := os.MkdirAll("output", 0755); err != nil {
		return err
	}
	return filepath.WalkDir("content", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}

		fmt.Println("processing:", path)
		if err := processPage.ProcessPage(path); err != nil {
			fmt.Println("error processing", path, ":", err)
			return nil
		}
		return nil
	})
}

func RunNew(title string) error {

	slug := strings.ToLower(strings.ReplaceAll(title, " ", "-"))
	filename := filepath.Join("content", slug+".md")
	date := time.Now().Format("2006-01-02")
	content := fmt.Sprintf("---\ntitle: \"%s\"\ndate: %s\n---\n\n# %s\n\nТекст статьи.\n", title, date, title)
	return os.WriteFile(filename, []byte(content), 0644)
}

func RunServer() {
	fs := http.FileServer(http.Dir("output"))
	http.Handle("/", fs)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error:", err)
	}

}
