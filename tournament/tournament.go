package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type clubScore struct {
	win    int
	loss   int
	draw   int
	points int
}

type sortClub struct {
	name   string
	points int
}

var club map[string]*clubScore
var err error

// Tally generates soccer tournament tally
func Tally(reader io.Reader, buffer io.Writer) error {
	club = map[string]*clubScore{}
	rd := bufio.NewReader(reader)

	for {
		line, err := rd.ReadString('\n')

		if err == io.EOF {
			if err = fillMap(line); err != nil {
				return err
			}
			break
		}

		if err = fillMap(line); err != nil {
			return err
		}
	}

	sortC := make([]sortClub, 0, len(club))
	for k := range club {
		sortC = append(sortC, sortClub{name: k, points: club[k].points})
	}
	sortClubs(sortC)

	_, err := io.WriteString(buffer, "Team                           | MP |  W |  D |  L |  P\n")
	if err != nil {
		return err
	}
	for _, c := range sortC {
		s := fmt.Sprintf("%-31v|%3v |%3v |%3v |%3v |%3v\n", c.name, club[c.name].win+club[c.name].loss+club[c.name].draw, club[c.name].win, club[c.name].draw, club[c.name].loss, club[c.name].points)
		if _, err := io.WriteString(buffer, s); err != nil {
			return err
		}
	}
	return nil
}

func update(c string, r []string) error {
	if len(r) < 3 {
		return errors.New("Incomplete result")
	}

	result := strings.Replace(r[2], "\n", "", -1)
	if c == r[0] && result == "win" {
		club[c].win++
		club[c].points += 3
		return nil
	} else if c == r[0] && result == "loss" {
		club[c].loss++
		return nil
	} else if c == r[1] && result == "win" {
		club[c].loss++
		return nil
	} else if c == r[1] && result == "loss" {
		club[c].win++
		club[c].points += 3
		return nil
	} else if result == "draw" {
		club[c].draw++
		club[c].points++
		return nil
	}
	return errors.New("No if condition captured")
}

func sortClubs(clubs []sortClub) {
	sort.SliceStable(clubs, func(i, j int) bool {
		ci, cj := clubs[i], clubs[j]
		switch {
		case ci.points != cj.points:
			return ci.points > cj.points
		default:
			return ci.name < cj.name
		}
	})
}

func fillMap(result string) error {

	if len(result) > 1 && string(result[0]) != "#" {
		s := strings.Split(result, ";")
		if _, ok := club[s[0]]; !ok {
			club[s[0]] = &clubScore{}
		}
		if err = update(s[0], s); err != nil {
			return err
		}

		if _, ok := club[s[1]]; !ok {
			club[s[1]] = &clubScore{}
		}

		if err = update(s[1], s); err != nil {
			return err
		}
	}
	return nil
}
