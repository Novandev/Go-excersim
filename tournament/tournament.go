package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

//TestVersion is the verion of the unit tests that this will pass
const TestVersion = 2

//scoreBoard keeps track of the score of the current team
type scoreBoard map[string]*team

// this struct is for kkeeping the win/ loss record of the current team
type team struct {
	name                            string
	played, win, loss, draw, points int
}

// For each eason this will tally up all of the wins and losses
func Tally(reader io.Reader, writer io.Writer) error {
	board := make(scoreBoard)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// error handler
		if err := board.addGame(line); err != nil {
			return err
		}
	}

	teams := board.getTeams()
	sort.Sort(byScore(teams))
	header := "Team                           | MP |  W |  D |  L |  P\n"
	io.WriteString(writer, header)
	for _, team := range teams {
		io.WriteString(writer, team.String()+"\n")
	}
	return nil
}
// Adds a game to the scoreboard
func (b scoreBoard) addGame(game string) error {
	fields := strings.Split(game, ";")
	if len(fields) != 3 {
		return fmt.Errorf("Game not properly formatted: %s", game)
	}
	home, homeOk := b[fields[0]]
	away, awayOk := b[fields[1]]
	if !homeOk {
		home = &team{name: fields[0]}
		b[fields[0]] = home
	}
	if !awayOk {
		away = &team{name: fields[1]}
		b[fields[1]] = away
	}
	switch fields[2] {
	case "win":
		home.addWin()
		away.addLoss()
	case "loss":
		home.addLoss()
		away.addWin()
	case "draw":
		home.addDraw()
		away.addDraw()
	default:
		return fmt.Errorf("Unknown win condition: %s", game)
	}
	return nil
}

// takes the scoreboard from the struct and turns it into a list
func (b scoreBoard) getTeams() []team {
	var teams []team
	for _, team := range b {
		teams = append(teams, *team)
	}
	return teams
}

// This sorts the team a-z fashin and then by score, not the cleanest but ittl work
type byScore []team

func (t byScore) Len() int      { return len(t) }
func (t byScore) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t byScore) Less(i, j int) bool {
	if t[i].points != t[j].points {
		return t[i].points > t[j].points
	} else if t[i].win != t[j].win {
		return t[i].win > t[j].win
	}
	return t[i].name < t[j].name
}

// adds a win to the team
func (t *team) addWin() {
	t.played++
	t.win++
	t.points += 3
}
// Adds a loss to the team
func (t *team) addLoss() {
	t.played++
	t.loss++
}

// If the team is a draw, it adds a draw to the team
func (t *team) addDraw() {
	t.played++
	t.draw++
	t.points++
}

// gets the cruurent team recors and dislays woth | as delinators
func (t *team) String() string {
	fmtString := "%-31s| %2d | %2d | %2d | %2d | %2d"
	return fmt.Sprintf(fmtString, t.name, t.played, t.win, t.draw, t.loss, t.points)
}