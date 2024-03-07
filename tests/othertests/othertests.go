package othertests

import (
	"RegistrationToGames/bot/forall"
)

func TestFromIntToStrDate() (correct bool) {
	date := forall.FromIntToStrDate(20250212)
	if date == "12-02-2025" {
		correct = true
	}
	return correct
}

func TestFromIntToStrTime() (correct bool) {
	time := forall.FromIntToStrTime(1200)
	if time == "12:00" {
		correct = true
	}
	return correct
}
