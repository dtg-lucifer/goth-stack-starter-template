package handlers

import (
	"net/http"

	"github.com/dtg-lucifer/goth-stack-starter/components/counter"
	"github.com/dtg-lucifer/goth-stack-starter/state"
	"github.com/dtg-lucifer/goth-stack-starter/utils"
)

// CounterIncrement increments the counter and returns the updated component
func CounterIncrement(w http.ResponseWriter, r *http.Request) error {
	currentCount := state.GlobalStore.GetInt("counter", 0)
	currentCount++
	state.GlobalStore.Set("counter", currentCount)

	return utils.Render(w, r, counter.Counter(currentCount))
}

// CounterDecrement decrements the counter and returns the updated component
func CounterDecrement(w http.ResponseWriter, r *http.Request) error {
	currentCount := state.GlobalStore.GetInt("counter", 0)
	if currentCount > 0 {
		currentCount--
	}
	state.GlobalStore.Set("counter", currentCount)

	return utils.Render(w, r, counter.Counter(currentCount))
}

// CounterReset resets the counter to zero and returns the updated component
func CounterReset(w http.ResponseWriter, r *http.Request) error {
	state.GlobalStore.Set("counter", 0)
	return utils.Render(w, r, counter.Counter(0))
}
