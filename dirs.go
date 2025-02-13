package freezer

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/anicoll/straw"
)

func seqToPath(basepath string, seq int) string {
	seqStr := fmt.Sprintf("%014d", seq)
	path := basepath
	for len(seqStr) > 0 {
		path = filepath.Join(path, seqStr[0:2])
		seqStr = seqStr[2:]
	}
	return path
}

// This is the depth of parent directories within which the data is stored.
const dirDepth = 6

func nextSequence(ss straw.StreamStore, basedir string) (int, error) {
	_, err := ss.Stat(seqToPath(basedir, 0))
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return -1, err
	}

	dir := basedir
	total := 0
	for i := 0; i <= dirDepth; i++ {
		fis, err := ss.Readdir(dir)
		if err != nil {
			return -1, err
		}
		l := len(fis)
		if l == 0 {
			return -1, fmt.Errorf("'%s' does not contain enough folders", basedir)
		}
		fi := fis[l-1]
		if i < dirDepth && !fi.IsDir() {
			return -1, fmt.Errorf("'%s' is not a directory", fi.Name())
		}
		dir = filepath.Join(dir, fi.Name())
		num, err := strconv.Atoi(fi.Name())
		if err != nil {
			return -1, err
		}
		total *= 100
		total += num
	}
	return total + 1, nil
}
