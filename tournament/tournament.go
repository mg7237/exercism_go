package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type club struct {
	name   string
	win    int
	loss   int
	draw   int
	points int
}

var clubs []club
var err error

// Tally generates soccer tournament tally
func Tally(reader io.Reader, buffer io.Writer) error {
	var results []string
	rd := bufio.NewReader(reader)

	for {
		line, err := rd.ReadString('\n')

		if err == io.EOF {
			results = append(results, line)
			break
		}
		if err != nil {
			return err
		}
		results = append(results, line)
	}
	for _, result := range results {
		if len(result) > 1 && string(result[0]) != "#" {
			s := strings.Split(result, ";")
			j, ok := findClub(s[0])
			if !ok {
				clubs = append(clubs, club{
					name: s[0],
					win:  0,
					loss: 0,
					draw: 0,
				})
				j = len(clubs) - 1
			}

			err = update(j, s)
			if err != nil {
				return err
			}

			k, ok := findClub(s[1])
			if !ok {
				clubs = append(clubs, club{
					name: s[1],
					win:  0,
					loss: 0,
					draw: 0,
				})
				k = len(clubs) - 1
			}
			err = update(k, s)
			if err != nil {
				return err
			}
		}
	}
	sortClubs()
	_, err := io.WriteString(buffer, "Team                           | MP |  W |  D |  L |  P\n")
	if err != nil {
		return err
	}
	for _, club := range clubs {
		s := fmt.Sprintf("%-31v|%3v |%3v |%3v |%3v |%3v\n", club.name, club.win+club.loss+club.draw, club.win, club.draw, club.loss, club.points)
		_, err := io.WriteString(buffer, s)
		if err != nil {
			return err
		}
	}
	clubs = nil
	return nil
}
func update(i int, r []string) error {
	if len(r) < 3 {
		return errors.New("Incomplete result")
	}

	result := strings.Replace(r[2], "\n", "", -1)

	if clubs[i].name == r[0] && result == "win" {
		clubs[i].win++
		clubs[i].points = clubs[i].points + 3
		return nil
	} else if clubs[i].name == r[0] && result == "loss" {
		clubs[i].loss++
		return nil
	} else if clubs[i].name == r[1] && result == "win" {
		clubs[i].loss++
		return nil
	} else if clubs[i].name == r[1] && result == "loss" {
		clubs[i].win++
		clubs[i].points = clubs[i].points + 3
		return nil
	} else if result == "draw" {
		clubs[i].draw++
		clubs[i].points++
		return nil
	}
	return errors.New("No if condition captured")
}

func findClub(val string) (int, bool) {
	for i, club := range clubs {
		if club.name == val {
			return i, true
		}
	}
	return -1, false
}

func sortClubs() {
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
