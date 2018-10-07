package tournament

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//Team ...
type Team struct {
	Name    string
	Total   int
	Win     int
	Loss    int
	Draw    int
	Matches int
}

var outcome = map[string]bool{
	"win":  true,
	"draw": true,
	"loss": true,
}

//Tally ...
func Tally(r io.Reader, w io.Writer) error {

	scanner := bufio.NewScanner(r)

	var teams = make(map[string]*Team)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || line[0] == '#' {
			continue
		}

		split := strings.Split(line, ";")
		if len(split) < 3 {
			return fmt.Errorf("missing ; seperator: %s", line)
		}

		team1 := split[0]
		team2 := split[1]
		outCome := split[2]

		_, ok := outcome[outCome]
		if !ok {
			return fmt.Errorf("missing ; seperator: %s", line)
		}

		switch outCome {

		case "win":
			var t1, t2 *Team
			t1, ok := teams[team1]

			if !ok {
				t1 = &Team{Name: team1}
				teams[team1] = t1
			}

			t1.Win++
			t1.Matches++
			t1.Total += 3

			t2, ok = teams[team2]
			if !ok {
				t2 = &Team{Name: team2}
				teams[team2] = t2
			}
			t2.Loss++
			t2.Matches++

		case "loss":

			var t1, t2 *Team
			t2, ok := teams[team2]

			if !ok {
				t2 = &Team{Name: team2}
				teams[team2] = t2
			}

			t2.Win++
			t2.Matches++
			t2.Total += 3

			t1, ok = teams[team1]
			if !ok {
				t1 = &Team{Name: team1}
				teams[team1] = t1
			}
			t1.Loss++
			t1.Matches++

		case "draw":

			var t1, t2 *Team
			t2, ok := teams[team2]

			if !ok {
				t2 = &Team{Name: team2}
				teams[team2] = t2
			}

			t2.Draw++
			t2.Matches++
			t2.Total++

			t1, ok = teams[team1]
			if !ok {
				t1 = &Team{Name: team1}
				teams[team1] = t1
			}
			t1.Draw++
			t1.Matches++
			t1.Total++
		}

	}

	var t []Team

	for _, v := range teams {
		t = append(t, *v)
	}

	sortTeam(t)

	s := scoreboard(t)

	fmt.Fprint(w, s)

	return nil
}

func scoreboard(teams []Team) string {

	var s string

	s = fmt.Sprintf("%-31s| MP |  W |  D |  L |  P\n", "Team")

	for _, v := range teams {
		s += fmt.Sprintf("%-31s|  %d |  %d |  %d |  %d |  %d\n", v.Name, v.Matches, v.Win, v.Draw, v.Loss, v.Total)
	}

	return s
}
