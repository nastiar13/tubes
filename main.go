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

const NMAX = 20
const MAX_MATCHES = 380

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
	t.AppendHeader(table.Row{"Week", "Match 1", "Match 2", "Match 3", "Match 4", "Match 5", "Match 6", "Match 7", "Match 8", "Match 9", "Match 10"})

	for i := 0; i < league.n_matches; i += 10 {

		match1 := league.matches[i]
		match2 := league.matches[i+1]
		match3 := league.matches[i+2]
		match4 := league.matches[i+3]
		match5 := league.matches[i+4]
		match6 := league.matches[i+5]
		match7 := league.matches[i+6]
		match8 := league.matches[i+7]
		match9 := league.matches[i+8]
		match10 := league.matches[i+9]
		match1_val := match1.home + " VS " + match1.away
		match2_val := match2.home + " VS " + match2.away
		match3_val := match3.home + " VS " + match3.away
		match4_val := match4.home + " VS " + match4.away
		match5_val := match5.home + " VS " + match5.away
		match6_val := match6.home + " VS " + match6.away
		match7_val := match7.home + " VS " + match7.away
		match8_val := match8.home + " VS " + match8.away
		match9_val := match9.home + " VS " + match9.away
		match10_val := match10.home + " VS " + match10.away
		if match1.status {
			match1_val = match1.home + " " + strconv.Itoa(match1.home_goals) + " - " + strconv.Itoa(match1.away_goals) + " " + match1.away
		}
		if match2.status {
			match2_val = match2.home + " " + strconv.Itoa(match2.home_goals) + " - " + strconv.Itoa(match2.away_goals) + " " + match2.away
		}
		if match3.status {
			match3_val = match3.home + " " + strconv.Itoa(match3.home_goals) + " - " + strconv.Itoa(match3.away_goals) + " " + match3.away
		}
		if match4.status {
			match4_val = match4.home + " " + strconv.Itoa(match4.home_goals) + " - " + strconv.Itoa(match4.away_goals) + " " + match4.away
		}
		if match5.status {
			match5_val = match5.home + " " + strconv.Itoa(match5.home_goals) + " - " + strconv.Itoa(match5.away_goals) + " " + match5.away
		}
		if match6.status {
			match6_val = match6.home + " " + strconv.Itoa(match6.home_goals) + " - " + strconv.Itoa(match6.away_goals) + " " + match6.away
		}
		if match7.status {
			match7_val = match7.home + " " + strconv.Itoa(match7.home_goals) + " - " + strconv.Itoa(match7.away_goals) + " " + match7.away
		}
		if match8.status {
			match8_val = match8.home + " " + strconv.Itoa(match8.home_goals) + " - " + strconv.Itoa(match8.away_goals) + " " + match8.away
		}
		if match9.status {
			match9_val = match9.home + " " + strconv.Itoa(match9.home_goals) + " - " + strconv.Itoa(match9.away_goals) + " " + match9.away
		}
		if match10.status {
			match10_val = match10.home + " " + strconv.Itoa(match10.home_goals) + " - " + strconv.Itoa(match10.away_goals) + " " + match10.away
		}
		t.AppendRow(table.Row{i/10 + 1, match1_val, match2_val, match3_val, match4_val, match5_val, match6_val, match7_val, match8_val, match9_val, match10_val})
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
	for i := 0; i < league.n_matches; i += 10 {
		m := league.matches
		if !(m[i].status && m[i+1].status && m[i+2].status && m[i+3].status && m[i+4].status && m[i+5].status && m[i+6].status && m[i+7].status && m[i+8].status && m[i+9].status) {
			weeks = append(weeks, "Week "+strconv.Itoa(i/10+1))
		}
	}

	selected := ""
	week_prompt := &survey.Select{
		Message: "Select Week",
		Options: weeks,
	}
	survey.AskOne(week_prompt, &selected)

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
	for i := 0; i <= league.n_matches-10; i += 10 {
		m := league.matches
		if m[i].status || m[i+1].status || m[i+2].status || m[i+3].status || m[i+4].status || m[i+5].status || m[i+6].status || m[i+7].status || m[i+8].status || m[i+9].status {
			weeks = append(weeks, "Week "+strconv.Itoa(i/10+1))
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
		fmt.Println("\n\n========================================================================")
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
	for i := 0; i <= league.n_matches-10; i += 10 {
		m := league.matches
		if m[i].status || m[i+1].status || m[i+2].status || m[i+3].status || m[i+4].status || m[i+5].status || m[i+6].status || m[i+7].status || m[i+8].status || m[i+9].status {
			weeks = append(weeks, "Week "+strconv.Itoa(i/10+1))
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
		fmt.Println("\n\n========================================================================")
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
	numWeeks := (numTeams - 1)
	matchesPerWeek := numTeams / 2

	for week := 0; week < numWeeks; week++ {

		for match := 0; match < matchesPerWeek; match++ {
			home := (week + match) % (numTeams - 1)
			away := (numTeams - 1 - match + week) % (numTeams - 1)

			if match == 0 {
				away = numTeams - 1
			}

			matches[week*matchesPerWeek+match] = Match{
				home: teams[home].name,
				away: teams[away].name,
				week: week + 1,
			}
			// fmt.Println("Week", week+1, "home:", home, "away:", away)
		}
	}
	for week := numWeeks; week < 2*numWeeks; week++ {
		for match := 0; match < matchesPerWeek; match++ {
			away := (week + match) % (numTeams - 1)
			home := (numTeams - 1 - match + week) % (numTeams - 1)

			if match == 0 {
				home = numTeams - 1
			}

			matches[week*matchesPerWeek+match] = Match{
				home: teams[home].name,
				away: teams[away].name,
				week: week + 1,
			}
			// fmt.Println("Week", week+1, "home:", home, "away:", away)
		}
	}

	return matches
}

func menu(league *League, clubs Clubs) {

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
	var clubs Clubs

	var list_clubs = []string{"ARS", "AVL", "BRE", "BHA", "BUR", "CHE", "CRY", "EVE", "LEE", "LEI", "LIV", "MCI", "MUN", "NEW", "NOR", "SOU", "TOT", "WAT", "WHU", "WOL"}
	for i := 0; i < len(list_clubs); i++ {
		var club Club
		club.name = list_clubs[i]
		clubs[i] = club

	}

	setup_league(clubs, &epl)
	matches := generateMatches(epl.classement)
	epl.matches = matches
	epl.n_matches = len(matches)

	menu(&epl, clubs)

}
