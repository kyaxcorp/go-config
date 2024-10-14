package events

import (
	"errors"

	"github.com/kyaxcorp/go-helper/conv"
	"github.com/kyaxcorp/go-helper/function"
)

var onLoadedAsyncCallbacks = make(map[string]func())

func OnLoadedAsync(callback func()) (string, error) {
	// Register the callback
	if function.IsCallable(callback) {
		// Generating a new unique id
		currentVal := conv.UInt64ToStr(callbackId.Inc(1))
		// Setting the callback
		onLoadedAsyncCallbacks[currentVal] = callback
		return currentVal, nil
	}
	return "", errors.New("invalid config callback on loaded async")
}

// CallOnLoadedAsync -> when everything is loaded, this function is being called
func CallOnLoadedAsync() {
	//log.Println("Calling CallOnLoadedAsync")
	for _, callback := range onLoadedAsyncCallbacks {
		go callback()
	}
}
