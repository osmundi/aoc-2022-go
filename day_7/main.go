package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type file struct {
	name string
	size int
}

type directory struct {
	name        string
	parent      *directory
	files       []file
	directories []*directory
}

var file_system = directory{
	name:        "/",
	parent:      nil,
	files:       []file{},
	directories: nil,
}

func (dir directory) GetSize() int {
	size := 0
	for _, file := range dir.files {
		size += file.size
	}

	for _, sub_dir := range dir.directories {
		size += sub_dir.GetSize()
	}

	return size
}

func (dir directory) SumDirectoriesUnder100kb() int {
	sum := 0

	if dir.GetSize() <= 100000 {
		sum += dir.GetSize()
	}

	for _, sub_dir := range dir.directories {
		sum += sub_dir.SumDirectoriesUnder100kb()
	}

	return sum
}

func (dir directory) GetAllSubDirecotries(directories *[]directory) []directory {

    // append current dir to list
    *directories = append(*directories, dir)

    // loop  through all sub directories
	for _, sub_dir := range dir.directories {
        *directories = sub_dir.GetAllSubDirecotries(directories)
	}

    return *directories
}

func (dir directory) Move(cmd string) *directory {
	if cmd == ".." {
		return dir.parent
	} else if cmd == "/" {
		return &file_system
	} else {
		for _, d := range dir.directories {
			if d.name == cmd {
				return d
			}
		}
	}
	return nil
}

func create_directory(parent *directory, name string) directory {
	return directory{name: name, parent: parent}
}

func create_file(size int, name string) file {
	return file{name: name, size: size}
}

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

	current_directory := &file_system
	reading_directory := false

	for _, command := range data {
		if strings.HasPrefix(command, "$") {

			if reading_directory {
				reading_directory = !reading_directory
			}

			command_list := strings.Split(command, " ")

			switch command_list[1] {
			case "cd":
				current_directory = current_directory.Move(command_list[2])
			case "ls":
				reading_directory = true
				continue
			}
		}

		if reading_directory {
			command_list := strings.Split(command, " ")

			switch command_list[0] {
			case "dir":
				dir := create_directory(current_directory, command_list[1])
				current_directory.directories = append(current_directory.directories, &dir)
			default: 
                // file
				size_of_the_file, err := strconv.Atoi(command_list[0])
				if err == nil {
					file := create_file(size_of_the_file, command_list[1])
					current_directory.files = append(current_directory.files, file)
				}
			}
		}

	}

	if part == "1" {
		fmt.Println(file_system.SumDirectoriesUnder100kb())
	} else {

		total_disk_size := 70000000
		unused_spaced_needed := 30000000
		current_size_of_file_system := file_system.GetSize()
		space_needed_to_free_up := unused_spaced_needed - (total_disk_size - current_size_of_file_system)
        var directories_candidate_for_removal []directory
        var all_directories []directory
        all_directories = file_system.GetAllSubDirecotries(&all_directories)

        for _,dir := range all_directories {
            if dir.GetSize() > space_needed_to_free_up {
                directories_candidate_for_removal = append(directories_candidate_for_removal, dir)
            }
        }

        sort.Slice(directories_candidate_for_removal, func(i, j int) bool {
          return directories_candidate_for_removal[i].GetSize() < directories_candidate_for_removal[j].GetSize()
        })

        fmt.Println(directories_candidate_for_removal[0].name)
        fmt.Println(directories_candidate_for_removal[0].GetSize())

	}

}
