package RecentFileActivity

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func RecentFileActivityCheck() {
	recdir := filepath.Join(os.Getenv("APPDATA"), "microsoft", "windows", "recent")
	files, err := ioutil.ReadDir(recdir)
	if err != nil {
		return
	}
	if len(files) < 20 {
		os.Exit(-1)
	}
}
