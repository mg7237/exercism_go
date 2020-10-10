package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type teamScore struct {
	name   string
	win    int
	loss   int
	draw   int
	points int
	mp     int
}

var err error

// Tally generates soccer tournament tally
func Tally(reader io.Reader, buffer io.Writer) error {
	team := make(map[string]teamScore)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		s := strings.Split(line, ";")
		if len(s) != 3 {
			return errors.New("Invalid line: Data provided = " + line + " ; expecting line with exactly 2 semicolons")
		}
		a := team[s[0]]
		a.name = s[0]
		b := team[s[1]]
		b.name = s[1]
		switch s[2] {
		case "win":
			a.win++
			a.points += 3
			b.loss++
		case "loss":
			b.win++
			b.points += 3
			a.loss++
		case "draw":
			a.draw++
			a.points++
			b.draw++
			b.points++
		default:
			return errors.New("Invalid result: Value Provided = " + s[2] + "; Expected either win,loss or draw ")
		}
		a.mp++
		b.mp++
		team[s[0]] = a
		team[s[1]] = b

	}
	if err := scanner.Err(); err != nil {
		return err
	}

	sortC := make([]teamScore, 0, len(team))
	for k := range team {
		sortC = append(sortC, team[k])
	}

	sort.Slice(sortC, func(i, j int) bool {
		ci, cj := sortC[i], sortC[j]
		if ci.points != cj.points {
			return ci.points > cj.points
		}
		return ci.name < cj.name
	})

	_, err := fmt.Fprintf(buffer, "Team                           | MP |  W |  D |  L |  P\n")
	if err != nil {
		return err
	}

	for _, c := range sortC {
		_, err := fmt.Fprintf(buffer, "%-31v|%3v |%3v |%3v |%3v |%3v\n", c.name, c.mp, c.win, c.draw, c.loss, c.points)
		if err != nil {
			return err
		}
	}
	return nil
}
