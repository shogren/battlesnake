package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type Game struct {
	ID      string `json:"id"`
	Timeout int32  `json:"timeout"`
}

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Battlesnake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int32   `json:"health"`
	Body   []Coord `json:"body"`
	Head   Coord   `json:"head"`
	Length int32   `json:"length"`
	Shout  string  `json:"shout"`
}

type Board struct {
	Height int           `json:"height"`
	Width  int           `json:"width"`
	Food   []Coord       `json:"food"`
	Snakes []Battlesnake `json:"snakes"`
}

type BattlesnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

type GameRequest struct {
	Game  Game        `json:"game"`
	Turn  int         `json:"turn"`
	Board Board       `json:"board"`
	You   Battlesnake `json:"you"`
}

type MoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout,omitempty"`
}

// HandleIndex is called when your Battlesnake is created and refreshed
// by play.battlesnake.com. BattlesnakeInfoResponse contains information about
// your Battlesnake, including what it should look like on the game board.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	response := BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "shogibear", // Your Battlesnake username
		Color:      "#45b9e3", // Personalize Color
		Head:       "gamer", // Personalize Head
		Tail:       "bolt", // Personalize Tail
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

// HandleStart is called at the start of each game your Battlesnake is playing.
// The GameRequest object contains information about the game that's about to start.
func HandleStart(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	// Nothing to respond with here
	fmt.Print("START\n")
}

// HandleMove is called for each turn of each game.
// Valid responses are "up", "down", "left", or "right".
func HandleMove(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}


Turn := request.Turn
fmt.Println(Turn)

/* trying to get names to print alongside the turn numbers
for _, players := range request.Board.Snakes {
  PlayerNames := players.Name
  fmt.Println(PlayerNames)
}
*/


  // Choose a random direction to move in
	PossibleMoves := []string{"up", "right", "down", "left"}

  //or start by just choosing one like "right"
  //move := "up"

  //move in a circle
  //move := possibleMoves[request.Turn % 4]

  head := request.You.Head
  body := request.You.Body

  enemy := request.Board.Snakes


  BoardHeight := request.Board.Height -1
  BoardWidth := request.Board.Width -1

//Avoid walls

if head.Y == BoardHeight {
    PossibleMoves = Remove(PossibleMoves, "up")
    }

  if head.X == BoardWidth {
    PossibleMoves = Remove(PossibleMoves, "right")
    }

  if head.Y == 0 {
    PossibleMoves = Remove(PossibleMoves, "down")
    }

  if head.X == 0 {
    PossibleMoves = Remove(PossibleMoves, "left")
    }

//Avoid Self

  for _, pos := range body {

    if head.X -1 == pos.X && head.Y == pos.Y {
      PossibleMoves = Remove(PossibleMoves, "left")
    }

    if head.X +1 == pos.X && head.Y == pos.Y {
      PossibleMoves = Remove(PossibleMoves, "right")
    }

    if head.Y +1 == pos.Y && head.X == pos.X {
      PossibleMoves = Remove(PossibleMoves, "up")
    }

    if head.Y -1 == pos.Y && head.X == pos.X {
      PossibleMoves = Remove(PossibleMoves, "down")
    }

  }

  //avoid enemy Snakes
  for _, snake := range enemy {

    for _, pos := range snake.Body{

        if head.X -1 == pos.X && head.Y == pos.Y {
        PossibleMoves = Remove(PossibleMoves, "left")
      }
        if head.X +1 == pos.X && head.Y == pos.Y {
        PossibleMoves = Remove(PossibleMoves, "right")
      }
        if head.Y +1 == pos.Y && head.X == pos.X {
        PossibleMoves = Remove(PossibleMoves, "up")
      }
        if head.Y -1 == pos.Y && head.X == pos.X {
        PossibleMoves = Remove(PossibleMoves, "down")
      }
//don't hit head straight on
    if head.X -2 == snake.Head.X && head.Y == snake.Head.Y {
      PossibleMoves = Remove(PossibleMoves, "left")
            fmt.Println("avoided head-on")
    }
    if head.X +2 == snake.Head.X && head.Y == snake.Head.Y{
      PossibleMoves = Remove(PossibleMoves, "right")
            fmt.Println("avoided head-on")
    }
    if head.Y -2 == snake.Head.Y && head.X == snake.Head.X{
      PossibleMoves = Remove(PossibleMoves, "down")
            fmt.Println("avoided head-on")
    }
    if head.Y +2 == snake.Head.Y && head.X == snake.Head.X{
      PossibleMoves = Remove(PossibleMoves, "up")
            fmt.Println("avoided head-on")
    }
//avoid diagonals
    //NW diag
    if head.X -1 == snake.Head.X && head.Y +1 == snake.Head.Y {
      PossibleMoves = Remove(PossibleMoves, "left")
      PossibleMoves = Remove(PossibleMoves, "up")
            fmt.Println("avoided diagonal")
    }
    //NE diag
    if head.X +1 == snake.Head.X && head.Y +1 == snake.Head.Y {
      PossibleMoves = Remove(PossibleMoves, "right")
      PossibleMoves = Remove(PossibleMoves, "up")
            fmt.Println("avoided diagonal")
    }
    //SW diag
    if head.X -1 == snake.Head.X && head.Y -1 == snake.Head.Y {
      PossibleMoves = Remove(PossibleMoves, "left")
      PossibleMoves = Remove(PossibleMoves, "down")
            fmt.Println("avoided diagonal")
    }
    //SE diag
    if head.X +1 == snake.Head.X && head.Y -1 == snake.Head.Y {
      PossibleMoves = Remove(PossibleMoves, "left")
      PossibleMoves = Remove(PossibleMoves, "down")
            fmt.Println("avoided diagonal")
  }
}
  }


  move := PossibleMoves[rand.Intn(len(PossibleMoves))]

  

	response := MoveResponse{
		Move: move,
	}

  fmt.Println(PossibleMoves)
	fmt.Printf("MOVE: %s\n", response.Move)
  
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

func Remove(s []string, elem string) []string {
	for k, v := range s {
		if v == elem {
			return append(s[:k:k], s[k+1:]...)
		}
	}
	return s
}

// HandleEnd is called when a game your Battlesnake was playing has ended.
// It's purely for informational purposes, no response required.
func HandleEnd(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	// Nothing to respond with here
	fmt.Print("END\n")
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/start", HandleStart)
	http.HandleFunc("/move", HandleMove)
	http.HandleFunc("/end", HandleEnd)

	fmt.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
