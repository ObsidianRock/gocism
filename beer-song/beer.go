package beer

import (
	"fmt"
)

func Verse(verse int) (string, error) {

	if verse > 99 || verse < 0 {
		return "", fmt.Errorf("upper and lower must be between 99-0")
	}

	switch verse {
	case 0:
		return fmt.Sprintf("No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"), nil
	case 1:
		return fmt.Sprintf("1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"), nil
	case 2:
		return fmt.Sprintf("2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"), nil
	default:
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", verse, verse, verse-1), nil
	}

}

func Verses(upper, lower int) (string, error) {

	if upper < lower {
		return "", fmt.Errorf("upper less than lower")
	}

	var ver string
	for i := upper; i >= lower; i-- {
		s, err := Verse(i)
		if err != nil {
			return "", err
		}
		ver += s
		ver += string('\n')
	}

	return ver, nil
}

func Song() string {
	v, _ := Verses(99, 0)
	return v
}
