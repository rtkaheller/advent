package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
)

type File struct {
	Name string
	Size int
}

type Dir struct {
	Files  map[string]File
	Dirs   map[string]*Dir
	Name   string
	Parent *Dir
}

func (d *Dir) SizeOf() (map[string]int, int) {
	sum := 0
	out := make(map[string]int)
	for _, f := range d.Files {
		sum += f.Size
	}
	for _, dir := range d.Dirs {
		dirs_size, size := dir.SizeOf()
		for k, v := range dirs_size {
			out[k] = v
		}
		sum += size
	}
	out[d.Name] = sum
	return out, sum
}

func main() {
	contents, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	root_dir := Dir{Name: "root", Files: make(map[string]File), Dirs: make(map[string]*Dir)}
	cur_dir := &root_dir
	lines := bytes.Split(contents, []byte("\n"))
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			continue
		}
		if line[0] == '$' {
			cmds := bytes.Split(line, []byte(" "))
			switch string(cmds[1]) {
			case "cd":
				if string(cmds[2]) == ".." {
					cur_dir = cur_dir.Parent
				} else if string(cmds[2]) == "/" {
					cur_dir = &root_dir
				} else {
					cur_dir = cur_dir.Dirs[string(cmds[2])]
				}
			case "ls":
				i += 1
				for ; i < len(lines)-1 && lines[i][0] != '$'; i++ {
					ls_out := bytes.Split(lines[i], []byte(" "))
					if string(ls_out[0]) == "dir" {
						cur_dir.Dirs[string(ls_out[1])] = &Dir{Name: cur_dir.Name + string(ls_out[1]), Parent: cur_dir, Files: make(map[string]File), Dirs: make(map[string]*Dir)}
					} else {
						size, _ := strconv.Atoi(string(ls_out[0]))
						cur_dir.Files[string(ls_out[1])] = File{Name: string(ls_out[1]), Size: size}
					}
				}
				i -= 1
			}
		}
	}

	sum := 0
	sizes, root_size := root_dir.SizeOf()
	var sizel []int
	for _, v := range sizes {
		if v <= 100000 {
			sum += v
		}
		sizel = append(sizel, v)
	}
	fmt.Println(sum)
	need := 30000000 - (70000000 - root_size)
	sort.Ints(sizel)
	for _, v := range sizel {
		if v > need {
			fmt.Println(v)
			break
		}
	}
}
