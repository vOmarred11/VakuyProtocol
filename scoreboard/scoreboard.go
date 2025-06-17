package scoreboard

// Scoreboard is a cleaner version of the default SetDisplayObjective
type Scoreboard struct {
	// Title is the title of the Scoreboard
	Title string
	// Lines is an array of all lines in the scoreboard
	Lines []struct {
		// Id is the entry id of the line
		Id int32
		// Text is the line text
		Text string
	}
}

// SetDisplayObjectiveScoreboard defines the location of the scoreboard
type SetDisplayObjectiveScoreboard struct {
	// Location si the location on the screen
	Location struct {
		X float64
		Y float64
	}
}

func (s *Scoreboard) ScoreboardTitle() string {
	return s.Title
}
func (s *Scoreboard) ScoreboardLinesId() int32 {
	return s.Lines[].Id
}
func (s *Scoreboard) ScoreboardLinesText() string {
	return s.Lines[].Text
}

func (s *SetDisplayObjectiveScoreboard) SetDisplayObjectiveScoreboard(loc SetDisplayObjectiveScoreboard) SetDisplayObjectiveScoreboard {
	location := loc.Location
	location.X = 850
	location.Y = 1150
	return loc
}
