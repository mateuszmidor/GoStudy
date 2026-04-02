package main

import (
	"fmt"
	"maps"
	"math/rand"
	"slices"
	"sync"
)

type cowboy struct {
	life int
}

func main() {
	cowboysMap := make(map[string]cowboy) // maps cowboy name -> actual cowboy
	var cowboysMutex sync.Mutex           // sync access to cowboysMap

	// 1. initialize cowboys
	for i := range 10 {
		name := fmt.Sprintf("cowboy-%d", i)
		cowboysMap[name] = cowboy{life: 10}
	}

	// 2. spawn cowboys
	var wg sync.WaitGroup
	wg.Add(len(cowboysMap))
	for shooterName := range cowboysMap {
		// spawn single cowboy of name <shooterName>
		go func() {
			defer wg.Done()

			// cowboy keeps shooting until dead or wins
			for {
				// technical: make sure no race condition
				cowboysMutex.Lock()

				// check if still alive
				_, ok := cowboysMap[shooterName]
				if !ok {
					fmt.Println(shooterName, "died!")
					cowboysMutex.Unlock()
					return
				}

				// check if already winner
				if len(cowboysMap) == 1 {
					fmt.Println(shooterName, "wins!")
					cowboysMutex.Unlock()
					return
				}

				// game is on, shoot random guy
				cowboyNames := slices.Collect(maps.Keys(cowboysMap))
				deleteMe := func(name string) bool { return name == shooterName }
				cowboyNamesExceptMe := slices.DeleteFunc(cowboyNames, deleteMe)
				victimNumber := rand.Intn(len(cowboyNamesExceptMe))
				victimName := cowboyNamesExceptMe[victimNumber]
				victim := cowboysMap[victimName]

				injuredVictim := cowboy{life: victim.life - 1}
				fmt.Println(shooterName, "shoots at", victimName, "whose life drops to", injuredVictim.life)
				if injuredVictim.life == 0 {
					delete(cowboysMap, victimName)
				} else {
					cowboysMap[victimName] = injuredVictim
				}

				// let others shoot
				cowboysMutex.Unlock()
			}
		}()
	}

	// 3. wait last man standing
	wg.Wait()
}
