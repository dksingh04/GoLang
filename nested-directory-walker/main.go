package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	args := []string{"."}

	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	//fmt.Println(args)
	for _, arg := range args {
		err := visit(arg)
		if err != nil {
			fmt.Printf("Error in visiting the Directories for paths %s: %v", arg, err)
		}
	}
	for _, arg := range args {
		err := tree(arg, "")
		if err != nil {
			fmt.Printf("Error in printing the Directories in tree structure for paths %s: %v", arg, err)
		}
	}
}

// visit all nested directories and the print the directories in tree structure and ignoring any
// folder or file starting with ., using Walk method
func visit(root string) error {
	err := filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		//fmt.Println(path)
		if err != nil {
			fmt.Printf("Failure accessing the Dir: %s and because of : %v", root, err)
			return err
		}

		rel, err := filepath.Rel(root, path)
		dir, fName := filepath.Split(rel)
		if ok := strings.Contains(dir, "."); ok || fName[0] == '.' {
			return nil
		}

		if err != nil {
			return fmt.Errorf("Could not rel (%s, %s): %v ", root, path, err)
		}
		depth := len(strings.Split(rel, string(filepath.Separator)))
		fmt.Println(strings.Repeat("  ", depth) + fi.Name())
		return nil
	})

	if err != nil {
		fmt.Printf("Unable acess the file path %s, because of the issue %v", root, err)
		return err
	}
	return nil
}

// Use recursive method to print the directory structure as it get printed through tree command.
// Thanks to Francesc Campoy for creating a wonderful Videos and sharing his knowledge
// Refer his video on youtube https://www.youtube.com/watch?v=XbKSssBftLM
func tree(root, indent string) error {
	fi, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("Unable to get the Stat %s: %v", root, err)
	}
	fmt.Println(fi.Name())
	fis, err := ioutil.ReadDir(root)
	var names []string
	for _, f := range fis {
		if f.Name()[0] != '.' {
			names = append(names, f.Name())
		}
	}
	for i, name := range names {
		add := "│  "
		if i == len(names)-1 {
			fmt.Printf(indent + "└──")
			add = "   "
		} else {
			fmt.Printf(indent + "├──")
		}
		// fmt.Printf(indent)
		if err := tree(filepath.Join(root, name), indent+add); err != nil {
			return err
		}
	}
	return nil
}
