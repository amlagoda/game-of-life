package universe

import (
	"math/rand"
	"time"
)

type Universe struct {
	generation       [][]bool
	generationNumber int
}

func NewUniverse(size int) *Universe {
	if size < 3 {
		return nil
	}

	generation := make([][]bool, size)
	for i := range generation {
		generation[i] = make([]bool, size)
	}

	return &Universe{generation, 0}
}

func (unv *Universe) NextGeneration() [][]bool {
	if unv.generationNumber == 0 {
		unv.generate()
	} else {
		unv.regenerate()
	}

	unv.generationNumber++
	return unv.generation
}

func (unv *Universe) CountAlive() int {
	cnt := 0

	for i := range unv.generation {
		for j := range unv.generation[i] {
			if unv.generation[i][j] {
				cnt++
			}
		}
	}

	return cnt
}

func (unv *Universe) NumberGeneration() int {
	return unv.generationNumber
}

func (unv *Universe) generate() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range unv.generation {
		for j := range unv.generation[i] {
			if random.Intn(2) != 0 {
				unv.generation[i][j] = true
			}
		}
	}
}

func (unv *Universe) regenerate() {
	rgn := unv.generation

	for i := range unv.generation {
		for j := range unv.generation[i] {
			rgn[i][j] = unv.isAlive(i, j)
		}
	}

	unv.generation = rgn
}

func (unv *Universe) isAlive(x int, y int) bool {
	cnt := 0
	size := len(unv.generation)

	top := size - 1
	if x-1 >= 0 {
		top = x - 1
	}

	bottom := 0
	if x+1 < size {
		bottom = x + 1
	}

	left := size - 1
	if y-1 >= 0 {
		left = y - 1
	}

	right := 0
	if y+1 < size {
		right = y + 1
	}

	cnt += toInt(unv.generation[top][y])
	cnt += toInt(unv.generation[top][left])
	cnt += toInt(unv.generation[top][right])
	cnt += toInt(unv.generation[x][left])
	cnt += toInt(unv.generation[x][right])
	cnt += toInt(unv.generation[bottom][y])
	cnt += toInt(unv.generation[bottom][left])
	cnt += toInt(unv.generation[bottom][right])

	return (unv.generation[x][y] && (cnt == 2 || cnt == 3)) ||
		(!unv.generation[x][y] && cnt == 3)
}

func toInt(b bool) int {
	if b {
		return 1
	}

	return 0
}
