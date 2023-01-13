package directory

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertCommands(commands []Command) (Directory, error) {
	root := newDirectory("", nil)

	var convertCommands func(commands []Command, currentDirectory Directory) error
	convertCommands = func(commands []Command, currentDirectory Directory) error {
		if len(commands) == 0 {
			return nil
		}

		castedDirectory, ok := currentDirectory.(*directoryImpl)
		if !ok {
			return fmt.Errorf("Could not access dictionary implementation")
		}

		currentCommand := commands[0]
		nextIteration := func(currentDirectory Directory) error {
			return convertCommands(commands[1:], currentDirectory)
		}

		if currentCommand.Name() == "cd" {
			return castCommand(currentCommand, func(castedCommand CdCommand) error {
				switch targetDir := castedCommand.TargetDirectory(); targetDir {
				case "/":
					return nextIteration(root)
				case "..":
					return nextIteration(currentDirectory.Parent())
				default:
					newDirectory := newDirectory(targetDir, currentDirectory)
					castedDirectory.subdirectories = append(castedDirectory.subdirectories, newDirectory)
					return nextIteration(newDirectory)
				}

			})
		} else if currentCommand.Name() == "ls" {
			return castCommand(currentCommand, func(castedCommand LsCommand) error {
				files := []File{}

				listingElems := strings.Split(castedCommand.Rawlisting(), "\n")
				for _, listingElem := range listingElems {
					if listingElem == "" {
						continue
					}
					listingElemParsed := strings.Split(listingElem, " ")
					left, right := listingElemParsed[0], listingElemParsed[1]

					if left != "dir" {
						size, err := strconv.Atoi(left)
						if err != nil {
							return fmt.Errorf("Could not parse size %d: %w", size, err)
						}
						files = append(files, newFile(right, size))
					}

				}

				castedDirectory.files = files
				return nextIteration(currentDirectory)
			})
		}

		return fmt.Errorf("Could not parse command of type %s", currentCommand.Name())
	}

	if err := convertCommands(commands, root); err != nil {
		return nil, err
	}

	return root, nil
}

func castCommand[T any](command Command, funcIfSuccess func(T) error) error {
	castedCommand, successCast := command.(T)
	if !successCast {
		return fmt.Errorf("Could not parse command. Command of type %s could not be cast to its corresponding type", command.Name())
	}
	return funcIfSuccess(castedCommand)
}

/*
 Input types
*/

type Command interface {
	Name() string
}

type LsCommand interface {
	Rawlisting() string
}

type CdCommand interface {
	TargetDirectory() string
}

/*
 Output types
*/

// Directory enables access to a directory's files and subdirectories
type Directory interface {
	Files() []File
	Parent() Directory
	Subdirectories() []Directory
	Name() string
}

// File - self explanatory
type File interface {
	Name() string
	Size() int
}
