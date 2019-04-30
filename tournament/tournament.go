package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name string
	win  int
	draw int
	loss int
}

func (t *team) matchPlayed() int {
	return t.win + t.draw + t.loss
}

func (t *team) point() int {
	return t.win*3 + t.draw
}

type byPoint []*team

func (t byPoint) Len() int {
	return len(t)
}

func (t byPoint) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t byPoint) Less(i, j int) bool {
	return t[i].point() > t[j].point()
}

type teams map[string]*team

func (t teams) formResult(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" || text[0] == '#' {
			continue
		}
		match := strings.Split(text, ";")
		if len(match) != 3 {
			return fmt.Errorf("wrong field count for line: %s (got %d fields)", text, len(match))
		}
		t1, t2, result := match[0], match[1], match[2]
		if _, ok := t[t1]; !ok {
			t[t1] = &team{name: t1}
		}
		if _, ok := t[t2]; !ok {
			t[t2] = &team{name: t2}
		}
		switch result {
		case "win":
			t[t1].win++
			t[t2].loss++
		case "loss":
			t[t1].loss++
			t[t2].win++
		case "draw":
			t[t1].draw++
			t[t2].draw++
		default:
			return fmt.Errorf("invalid result %s", result)
		}
	}
	return nil
}

func Tally(reader io.Reader, writer io.Writer) error {

	t := make(teams)
	err := t.formResult(reader)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(writer, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	if err != nil {
		return err
	}

	names := make([]string, 0, len(t))
	for _, t := range t {
		names = append(names, t.name)
	}
	sort.Strings(names)
	tms := make([]*team, 0, len(t))
	for _, name := range names {
		tms = append(tms, t[name])
	}

	sort.Sort(byPoint(tms))

	for _, t := range tms {
		_, err = fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			t.name, t.matchPlayed(), t.win, t.draw, t.loss, t.point())
		if err != nil {
			return err
		}
	}

	return nil
}
