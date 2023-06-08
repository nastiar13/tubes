package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jedib0t/go-pretty/v6/table"
)

const NMAX = 3
const MAX_MATCHES = 6

type Club struct {
	name string
	// Match Played, Won, Draw, Loss, Goals For, Goals Againts, Goals Different, Points
	MP, W, D, L, GF, GA, GD, Pts int
}

type Clubs [NMAX]Club

type Match struct {
	home, away             string
	home_goals, away_goals int
	week                   int
	status                 bool
}

type Matches [MAX_MATCHES]Match

type League struct {
	classement   Clubs
	matches      Matches
	n_classement int
	n_matches    int
}

func header() {
	fmt.Println("\n ========================== Welcome ==========================")
	fmt.Println("\t\t     EPL Management Prompt     ")
	fmt.Println("\t\t     Algoritma Pemrograman     ")
	fmt.Println(" ------------------------------------------------------------- ")
	fmt.Println(" ------------------------------------------------------------- ")
	fmt.Println("")
}

func setup_league(clubs Clubs, league *League) {
	league.classement = clubs
	league.n_classement = len(clubs)
}

func print_classement(league League) {

	sort.SliceStable(league.classement[:], func(i, j int) bool {
		return league.classement[i].Pts > league.classement[j].Pts
	})
	sort.SliceStable(league.classement[:], func(i, j int) bool {
		return league.classement[i].GD > league.classement[j].GD
	})
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Club", "MP", "W", "D", "L", "GF", "GA", "GD", "Pts"})
	for i := 0; i < league.n_classement; i++ {
		club := league.classement[i]
		t.AppendRow(table.Row{i + 1, club.name, club.MP, club.W, club.D, club.L, club.GF, club.GA, club.GD, club.Pts})
	}
	fmt.Println("")
	fmt.Println("===========================================================")
	fmt.Println("\t\t KLASEMEN ENGLISH PRO LEAGUE")
	fmt.Println("===========================================================")
	t.Render()
}

func print_matches(league League) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Week", "Match"})

	for i := 0; i < league.n_matches; i++ {

		match1 := league.matches[i]

		match1_val := match1.home + " VS " + match1.away

		if match1.status {
			match1_val = match1.home + " " + strconv.Itoa(match1.home_goals) + " - " + strconv.Itoa(match1.away_goals) + " " + match1.away
		}

		t.AppendRow(table.Row{i + 1, match1_val})
	}
	fmt.Println("")
	fmt.Println("=======================================================================")
	fmt.Println("\t\t\tJadwal Pertandingan")
	fmt.Println("=======================================================================")
	t.Render()
}

func check_club(league League, name string) int {
	for i := 0; i < league.n_classement; i++ {
		if league.classement[i].name == name {
			return i
		}
	}
	return -1
}

func check_winner(home_goals, away_goals int) int {

	if home_goals > away_goals {
		return 1
	} else if home_goals < away_goals {
		return 2
	} else {
		return 0
	}

}

func calculate_points(league *League) {
	for i := 0; i < league.n_classement; i++ {
		league.classement[i].MP = 0
		league.classement[i].W = 0
		league.classement[i].D = 0
		league.classement[i].L = 0
		league.classement[i].GF = 0
		league.classement[i].GA = 0
	}
	for i := 0; i < league.n_matches; i++ {
		match := league.matches[i]
		home_idx := check_club(*league, match.home)
		away_idx := check_club(*league, match.away)

		if match.status {
			if home_idx != -1 && away_idx != -1 {
				// increase MP after match was played
				league.classement[home_idx].MP++
				league.classement[away_idx].MP++
			}
			result := check_winner(match.home_goals, match.away_goals)
			// If the result equals 0, it means the match ended in a draw. If the result equals 1, it means home won the match. Otherwise, if the result equals 2, it means away won the match.
			if result == 1 {
				league.classement[home_idx].W++
				league.classement[away_idx].L++
			} else if result == 2 {
				league.classement[away_idx].W++
				league.classement[home_idx].L++
			} else {
				league.classement[home_idx].D++
				league.classement[away_idx].D++
			}
		}
	}

	for i := 0; i < league.n_matches; i++ {
		match := league.matches[i]
		for j := 0; j < league.n_classement; j++ {
			club := league.classement[j]
			if club.name == match.home {
				league.classement[j].GF += match.home_goals
				league.classement[j].GA += match.away_goals
			}
			if club.name == match.away {
				league.classement[j].GF += match.away_goals
				league.classement[j].GA += match.home_goals
			}

		}

	}

	for i := 0; i < league.n_classement; i++ {
		club := league.classement[i]
		league.classement[i].Pts = club.W*3 + club.D
		league.classement[i].GD = club.GF - club.GA
	}

}

func add_matches(league *League) {
	print_matches(*league)
	var weeks []string
	for i := 0; i < league.n_matches; i++ {
		if !(league.matches[i].status) {
			weeks = append(weeks, "Week "+strconv.Itoa(i+1))
		}
	}
	weeks = append(weeks, "Cancel")

	selected := ""
	week_prompt := &survey.Select{
		Message: "Select Week",
		Options: weeks,
	}
	survey.AskOne(week_prompt, &selected)
	if selected == "Cancel" {
		menu(league)
		return
	}

	var weekNumber int
	_, err := fmt.Sscanf(selected, "Week %d", &weekNumber)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var matches []string
	for i := 0; i < league.n_matches; i++ {
		if league.matches[i].week == weekNumber && !league.matches[i].status {
			home := league.matches[i].home
			away := league.matches[i].away
			matches = append(matches, home+" VS "+away)
		}
	}

	matches = append(matches, "Cancel")

	res := ""
	match_prompt := &survey.Select{
		Message: "Select Match",
		Options: matches,
	}
	survey.AskOne(match_prompt, &res)

	if res == "Cancel" {
		menu(league)
		return
	}

	var home, away string
	home = strings.Split(res, " VS ")[0]
	away = strings.Split(res, " VS ")[1]

	var home_goals, away_goals int
	home_goals_prompt := &survey.Input{
		Message: "Goals for " + home + ": ",
		Default: "0",
	}
	away_goals_prompt := &survey.Input{
		Message: "Goals for " + away + ": ",
		Default: "0",
	}

	survey.AskOne(home_goals_prompt, &home_goals)
	survey.AskOne(away_goals_prompt, &away_goals)

	for i := 0; i < league.n_matches; i++ {
		match := league.matches[i]
		if match.home == home && match.away == away {
			league.matches[i].home_goals = home_goals
			league.matches[i].away_goals = away_goals
			league.matches[i].status = true
			break
		}
	}

	calculate_points(league)

}

func reset_match(league *League) {
	var weeks []string
	for i := 0; i < league.n_matches; i++ {
		if league.matches[i].status {
			weeks = append(weeks, "Week "+strconv.Itoa(i+1))
		}
	}
	// if len(weeks) < 1 {

	// }

	selected := ""
	if len(weeks) > 0 {
		week_prompt := &survey.Select{
			Message: "Select Week",
			Options: weeks,
		}
		survey.AskOne(week_prompt, &selected)
	}

	var weekNumber int
	_, err := fmt.Sscanf(selected, "Week %d", &weekNumber)
	if err != nil {
		fmt.Println("========================================================================")
		fmt.Println("\t\tBelum ada pertandingan yang berlangsung")
		fmt.Println("========================================================================")
		return
	}

	var matches []string
	for i := 0; i < league.n_matches; i++ {
		if league.matches[i].week == weekNumber && league.matches[i].status {
			home := league.matches[i].home
			away := league.matches[i].away
			matches = append(matches, home+" VS "+away)
		}
	}

	res := ""
	match_prompt := &survey.Select{
		Message: "Select Match",
		Options: matches,
	}
	survey.AskOne(match_prompt, &res)

	var home, away string
	home = strings.Split(res, " VS ")[0]
	away = strings.Split(res, " VS ")[1]
	for i := 0; i < league.n_matches; i++ {
		if league.matches[i].home == home && league.matches[i].away == away {
			league.matches[i].home_goals = 0
			league.matches[i].away_goals = 0
			league.matches[i].status = false
		}
	}
	calculate_points(league)

}

func edit_match(league *League) {
	var weeks []string
	for i := 0; i < league.n_matches; i++ {
		if league.matches[i].status {
			weeks = append(weeks, "Week "+strconv.Itoa(i+1))
		}
	}
	// if len(weeks) < 1 {

	// }

	selected := ""
	if len(weeks) > 0 {
		week_prompt := &survey.Select{
			Message: "Select Week",
			Options: weeks,
		}
		survey.AskOne(week_prompt, &selected)
	}

	var weekNumber int
	_, err := fmt.Sscanf(selected, "Week %d", &weekNumber)
	if err != nil {
		// fmt.Println("Error:", err)
		fmt.Println("========================================================================")
		fmt.Println("\t\tBelum ada pertandingan yang berlangsung")
		fmt.Println("========================================================================")
		return
	}

	var matches []string
	for i := 0; i < league.n_matches; i++ {
		if league.matches[i].week == weekNumber && league.matches[i].status {
			home := league.matches[i].home
			away := league.matches[i].away
			matches = append(matches, home+" VS "+away)
		}
	}

	res := ""
	match_prompt := &survey.Select{
		Message: "Select Match",
		Options: matches,
	}
	survey.AskOne(match_prompt, &res)

	var home, away string
	home = strings.Split(res, " VS ")[0]
	away = strings.Split(res, " VS ")[1]

	var home_goals, away_goals int
	home_goals_prompt := &survey.Input{
		Message: "Goals for " + home + ": ",
		Default: "0",
	}
	away_goals_prompt := &survey.Input{
		Message: "Goals for " + away + ": ",
		Default: "0",
	}

	survey.AskOne(home_goals_prompt, &home_goals)
	survey.AskOne(away_goals_prompt, &away_goals)

	for i := 0; i < league.n_matches; i++ {
		if league.matches[i].home == home && league.matches[i].away == away {
			league.matches[i].home_goals = home_goals
			league.matches[i].away_goals = away_goals
		}
	}

	calculate_points(league)

}

func generateMatches(clubs Clubs) Matches {
	var matches Matches
	teams := clubs
	numTeams := len(teams)
	numWeeks := 0
	// matchesPerWeek := 2

	for i := 0; i < numTeams; i++ {
		for j := 0; j < numTeams; j++ {
			if teams[i].name != teams[j].name {
				matches[numWeeks] = Match{
					home: teams[i].name,
					away: teams[j].name,
					week: numWeeks + 1,
				}
				numWeeks++
			}
		}
	}
	// for week := numWeeks; week < 2*numWeeks; week++ {
	// 	for match := 0; match < matchesPerWeek; match++ {
	// 		away := (week + match) % (numTeams - 1)
	// 		home := (numTeams - 1 - match + week) % (numTeams - 1)

	// 		if match == 0 {
	// 			home = numTeams - 1
	// 		}

	// 		matches[week*matchesPerWeek+match] = Match{
	// 			home: teams[home].name,
	// 			away: teams[away].name,
	// 			week: week + 1,
	// 		}
	// 	}
	// }

	return matches
}

func menu(league *League) {

	var opsi string = ""

	for opsi != "Keluar" {

		header()

		menu_prompt := &survey.Select{
			Message: "Pilihan Anda",
			Options: []string{"1. Lihat klasemen", "2. Lihat Pertandingan", "3. Tambah Pertandingan", "4. Reset Pertandingan", "5. Edit Pertandingan", "Keluar"},
		}

		survey.AskOne(menu_prompt, &opsi)

		if opsi == "1. Lihat klasemen" {
			print_classement(*league)
		} else if opsi == "2. Lihat Pertandingan" {
			print_matches(*league)

		} else if opsi == "3. Tambah Pertandingan" {
			add_matches(league)

			print_classement(*league)
		} else if opsi == "4. Reset Pertandingan" {
			print_matches(*league)
			reset_match(league)

		} else if opsi == "5. Edit Pertandingan" {
			print_matches(*league)
			edit_match(league)

		}
	}

	fmt.Println("Terima Kasih :)")
}

func main() {
	var epl League
	var mu, mci, ars, newc Club
	mu.name = "Man United FC"
	mci.name = "Man City FC"
	ars.name = "Arsenal FC"
	newc.name = "Newcastle FC"
	var clubs = Clubs{ars, mci, mu}

	setup_league(clubs, &epl)
	matches := generateMatches(epl.classement)

	epl.matches = matches
	epl.n_matches = len(matches)

	menu(&epl)

}
