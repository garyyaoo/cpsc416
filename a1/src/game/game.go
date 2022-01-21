package game

import (
	"fmt"
	"github.com/DistributedClocks/tracing"
	"time"
)

type GameStart struct {
	Seed int8
}

type ClientMove StateMoveMessage

type ServerMoveReceive StateMoveMessage

type GameComplete struct {
	Winner string
}

/** Message structs **/

type StateMoveMessage struct {
	GameState []uint8
	MoveRow   int8
	MoveCount int8
}

func Start(trace *tracing.Trace, seed int8) {
	trace.RecordAction(
		GameStart{seed})

	end := make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		end <- true
	}()

	for <-end {
		fmt.Println("done")
		break
	}
}
