package main

import (
	"fmt"
	"maps"
	"math/rand"
	"slices"
	"sync"
	"time"
)

type cowboy struct {
	life int
}

func main() {
	var cowboysGang = make(map[string]cowboy) // maps the cowboy names to actual cowboys
	var winnerName string                     // the last cowboy alive
	var cowboysMutex sync.Mutex               // to sync access to cowboysMap
	var wg sync.WaitGroup                     // to wait till all cowboys are done, either dead or winning

	// 1. initialize all cowboys
	for i := range 10 {
		name := fmt.Sprintf("Cowboy%d", i)
		cowboysGang[name] = cowboy{life: 10}
	}

	// 2. spawn all cowboys
	wg.Add(len(cowboysGang))
	for cowboyName := range cowboysGang {
		// spawn a single cowboy of name <shooterName>
		go func() {
			isWinner := spawnCowboy(cowboyName, cowboysGang, &cowboysMutex)
			if isWinner {
				winnerName = cowboyName
			} else {
				fmt.Println(cowboyName, "died!")
			}
			wg.Done()
		}()
	}

	// 3. wait till last man standing
	wg.Wait()

	// 4. announce the winner
	fmt.Println(winnerName, "wins!")
}

func spawnCowboy(shooterName string, cowboysMap map[string]cowboy, cowboysMutex *sync.Mutex) (isWinner bool) {
	// cowboy keeps shooting until either winner or dead
	for {
		// new shooting attempt; make sure no race condition while reading/updating cowboysMap
		cowboysMutex.Lock()

		// check if cowboy still alive
		_, ok := cowboysMap[shooterName]
		if !ok {
			cowboysMutex.Unlock()
			return false // died in the shooting
		}

		// check if cowboy already a winner
		if len(cowboysMap) == 1 {
			cowboysMutex.Unlock()
			return true // won the shooting
		}

		// game is on, shoot random guy
		// 1. pick a target
		cowboyNames := slices.Collect(maps.Keys(cowboysMap))
		deleteMe := func(name string) bool { return name == shooterName }
		cowboyNamesExceptMe := slices.DeleteFunc(cowboyNames, deleteMe)
		victimNumber := rand.Intn(len(cowboyNamesExceptMe))
		victimName := cowboyNamesExceptMe[victimNumber]
		victim := cowboysMap[victimName]
		// 2. shoot at the target
		injuredVictim := cowboy{life: victim.life - 1}
		fmt.Println(shooterName, "shoots at", victimName, "whose life drops to", injuredVictim.life)
		if injuredVictim.life == 0 {
			delete(cowboysMap, victimName) // kill the victim
		} else {
			cowboysMap[victimName] = injuredVictim // injure the victim
		}

		// let others shoot
		cowboysMutex.Unlock()
		time.Sleep(50 * time.Millisecond)
	}
}
