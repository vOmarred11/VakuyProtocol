package scoreboard

// Arraylist is a cleaner version of the default SetDisplayObjective
type Arraylist struct {
	// Title is the title of the Arraylist
	Title string
	// Lines is an array of all lines in the scoreboard
	Lines []struct {
		// Id is the entry id of the line
		Id int32
		// Text is the line text
		Text string
	}
}

// SetDisplayObjectiveArraylist defines the location of the scoreboard
type SetDisplayObjectiveArraylist struct {
	// Location si the location on the screen
	Location struct {
		X float64
		Y float64
	}
}

func (a *Arraylist) ArraylistTitle() string {
	return a.Title
}
func (a *Arraylist) ArraylistLinesText() string {
	return a.Lines[].Text
}
func (a *Arraylist) ArraylistLinesId() int32 {
	return a.Lines[].Id
}

func (s *SetDisplayObjectiveArraylist) SetDisplayObjectiveArraylist(loc SetDisplayObjectiveArraylist) SetDisplayObjectiveArraylist {
	location := loc.Location
	location.X = 850
	location.Y = 1150
	return loc
}
