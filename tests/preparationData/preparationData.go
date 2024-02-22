package preparationdata

import (
	"registrationtogames/fmtogram/types"
	"registrationtogames/tests/registration"
	"registrationtogames/tests/routine"
	"registrationtogames/tests/welcome"
)

func PreparationToShowRules(counter int, responses chan *types.TelegramResponse) {
	if counter == 0 {
		welcome.JustTrash(responses)
	} else if counter == 1 {
		welcome.JustTrash2(responses)
	} else if counter == 2 {
		welcome.QueryForShowRules(responses)
	}
	routine.DeleteUser(456)
}

func WelcomeToMainMenu(counter int, responses chan *types.TelegramResponse) {
	if counter == 0 {
		welcome.JustTrash(responses)
	} else if counter == 1 {
		welcome.JustTrash2(responses)
	} else if counter == 2 {
		welcome.QueryForWelcomeToMainMenu(responses)
	}
	routine.DeleteUser(456)
}

func PresentationScheduele(counter int, responses chan *types.TelegramResponse) {
	if counter == 0 {
		registration.JustTrash(responses)
	} else if counter == 1 {
		registration.JustTrash2(responses)
	} else if counter == 2 {
		registration.QueryForPresentationScheduele(responses)
	}
	routine.DeleteUser(456)
}

func ChooseGame(counter int, responses chan *types.TelegramResponse) {
	if counter == 0 {
		registration.JustTrash(responses)
	} else if counter == 1 {
		registration.JustTrash2(responses)
	} else if counter == 2 {
		registration.QueryForChooseGame(responses)
	}
	routine.DeleteUser(456)
}

func ChooseSeats(counter int, responses chan *types.TelegramResponse) {
	if counter == 0 {
		registration.JustTrash(responses)
	} else if counter == 1 {
		registration.JustTrash2(responses)
	} else if counter == 2 {
		registration.QueryForChooseSeats(responses)
	}
	routine.DeleteUser(456)
}

func ChoosePayment(counter int, responses chan *types.TelegramResponse) {
	if counter == 0 {
		welcome.JustTrash(responses)
	} else if counter == 1 {
		registration.JustTrash2(responses)
	} else if counter == 2 {
		registration.QueryForChoosePayment(responses)
	}
	routine.DeleteUser(456)
}

func BestWishes(counter int, responses chan *types.TelegramResponse) {
	if counter == 0 {
		registration.JustTrash(responses)
	} else if counter == 1 {
		registration.JustTrash2(responses)
	} else if counter == 2 {
		registration.QueryForBestWishes(responses)
	}
	routine.DeleteUser(456)
}
