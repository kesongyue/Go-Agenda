package entity

type Meeting struct {
  Sponsor string
  Participator []string
  Start, End Date
  Title string
}

func GetSponsor(m Meeting) string {
	return m.Sponsor
}
func GetParticipator(m Meeting) []string{
	return m.Participator
}
func GetStart(m Meeting) Date{
	return m.Start
}
func GetEnd(m Meeting) Date{
	return m.End
}
func GetTitle(m Meeting) string {
	return m.Title
}