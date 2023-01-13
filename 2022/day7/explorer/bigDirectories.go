package explorer

func SumSmallDirectoriesSizes(root Directory) int {
	_, smallDirectorySizes := findSmallDirectoriesSizes(root)

	sum := 0
	for _, size := range smallDirectorySizes {
		sum += size
	}

	return sum
}

func findSmallDirectoriesSizes(current Directory) (int, []int) {
	fileSizes := 0
	for _, size := range current.FileSizes() {
		fileSizes += size
	}

	smallDescendants := []int{}
	subdirectoriesSize := 0
	for _, subdirectory := range current.SubDirectories() {
		subdirectorySize, subdirectorySmallDescendants := findSmallDirectoriesSizes(subdirectory)
		subdirectoriesSize += subdirectorySize
		smallDescendants = append(smallDescendants, subdirectorySmallDescendants...)
	}

	size := subdirectoriesSize + fileSizes
	if size < 100000 {
		return size, append(smallDescendants, size)
	} else {
		return size, smallDescendants
	}
}

type Directory interface {
	FileSizes() []int
	SubDirectories() []Directory
}
