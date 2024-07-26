package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"¡Ganaste!",
	"¡Felicidades!",
	"¡Excelente!",
	"¡Enhorabuena!",
	"¡Felicidades!",
}

var loseMessages = []string{
	"¡Perdiste!",
	"¡Mal!",
	"¡No has aprendido!",
	"¡Pierdes!",
	"¡Lo siento!",
}
var drawMessages = []string{
	"¡Empate!",
	"Iguales!",
	"Igualados!",
	"jajajaj iguales!",
	"Empatados!",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {

	computerValue := rand.Intn(3)
	var computerChoice, roundResult string

	var ComputerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoice = "La computadora eligio Piedra"
		ComputerChoiceInt = ROCK
	case PAPER:
		computerChoice = "La computadora eligio Papel"
		ComputerChoiceInt = PAPER
	case SCISSORS:
		computerChoice = "La computadora eligio Tijera"
		ComputerChoiceInt = SCISSORS
	}

	messageInt := rand.Intn(5)

	var message string

	if playerValue == ComputerChoiceInt {
		message = drawMessages[messageInt]
		roundResult = "Empate"
	} else if playerValue == (ComputerChoiceInt+1)%3 {
		PlayerScore++
		roundResult = "El jugador gana!!!"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "La computadora gana"
		message = loseMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: ComputerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
