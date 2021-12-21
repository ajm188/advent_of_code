package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/ajm188/advent_of_code/pkg/cli"
)

var (
	path = flag.String("path", "input.txt", "")

	debug = flag.Bool("debug", false, "")
)

func makering(capacity int) *RingSlice {
	b := make([]int, capacity)
	for i := 0; i < len(b); i++ {
		b[i] = i + 1
	}

	return &RingSlice{b: b}
}

func forAllPlayers(players []*Player, condition func(p *Player) bool) bool {
	for _, p := range players {
		if !condition(p) {
			return false
		}
	}

	return true
}

func sum(xs []int) (n int) {
	for _, x := range xs {
		n += x
	}

	return n
}

type Player struct {
	ID       int
	Position int
	Scores   []int // Scores[i] is the score the player had after i rolls.
}

func (player *Player) Score() int {
	if len(player.Scores) == 0 {
		return 0
	}

	return player.Scores[len(player.Scores)-1]
}

type PlayersByID []*Player

func (players PlayersByID) Len() int           { return len(players) }
func (players PlayersByID) Swap(i, j int)      { players[i], players[j] = players[j], players[i] }
func (players PlayersByID) Less(i, j int) bool { return players[i].ID < players[j].ID }

func main() {
	flag.Parse()

	data, err := cli.GetInput(*path)
	cli.ExitOnError(err)

	var players []*Player
	inputRegexp := regexp.MustCompile(`^Player (\d+) starting position: (\d+)$`)
	for i, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		m := inputRegexp.FindStringSubmatch(line)
		if m == nil {
			cli.ExitOnError(fmt.Errorf("line %d does not match %s", i, inputRegexp))
		}

		id, err := strconv.ParseInt(m[1], 10, 64)
		cli.ExitOnErrorf(err, "line %d could not parse player id: %s", i, err)

		pos, err := strconv.ParseInt(m[2], 10, 64)
		cli.ExitOnErrorf(err, "line %d could not parse starting position: %s", i, err)

		players = append(players, &Player{
			ID:       int(id),
			Position: int(pos),
		})
	}

	sort.Sort(PlayersByID(players))
	dice := makering(100)
	positions := makering(10)

	if *debug {
		log.Printf("%v", positions.b)
	}

	var (
		wg     sync.WaitGroup
		_range = [2]int{0, 3}
	)
	for _, player := range players {
		wg.Add(1)
		go func(player *Player, _range [2]int) {
			defer wg.Done()

			for player.Score() < 1000 {
				prevPos := player.Position
				rolls := dice.GetSlice(_range[0], _range[1])
				player.Position = positions.Get(player.Position + sum(rolls) - 1)
				player.Scores = append(player.Scores, player.Score()+player.Position)

				if *debug && player.ID == 1 {
					log.Printf("Player %d roll #%d: %v => moved from %d to %d for total score of %d", player.ID, len(player.Scores)-1, rolls, prevPos, player.Position, player.Score())
				}

				_range[0] += 6
				_range[1] += 6
			}
		}(player, [2]int{_range[0], _range[1]})

		_range[0] += 3
		_range[1] += 3
	}

	wg.Wait()

	var winner *Player
	for i := 0; forAllPlayers(players, func(p *Player) bool { return i < len(p.Scores) }); i++ {
		for _, p := range players {
			if p.Scores[i] >= 1000 {
				winner = p
				break
			}
		}

		if winner != nil {
			break
		}
	}

	rolls := 0
	loserScore := 0
	for _, p := range players {
		if p.ID <= winner.ID {
			rolls += 3 * len(winner.Scores)
		} else {
			rolls += 3 * (len(winner.Scores) - 1)
		}

		if p.ID != winner.ID {
			loserScore += p.Scores[len(winner.Scores)-2]
		}
	}

	fmt.Println(rolls * loserScore)
}
