// Binding for JS to Go
// This package is responsible for the communication between the JS and Go code.
package bindings

import (
	"context"
	"log"
	"strings"
	"sync"

	"runtime"

	"github.com/PiterWeb/RemoteController/src/devices/gamepad"
	"github.com/PiterWeb/RemoteController/src/devices/keyboard"
	net "github.com/PiterWeb/RemoteController/src/net/webrtc"
	"github.com/pion/webrtc/v3"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

var triggerEnd chan struct{} = make(chan struct{})

var openPeer bool = false
var openPeerMutex sync.Mutex

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called at application Startup
func (a *App) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// BeforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) BeforeClose(ctx context.Context) (prevent bool) {

	openPeerMutex.Lock()
	defer openPeerMutex.Unlock()

	if !openPeer {
		prevent = false
		return prevent
	}

	// Show a dialog to confirm the user wants to quit
	option, err := wailsRuntime.MessageDialog(ctx, wailsRuntime.MessageDialogOptions{
		Type:          wailsRuntime.QuestionDialog,
		Title:         "Quit",
		Message:       "Are you sure you want to quit?",
		Buttons:       []string{"Yes", "No"},
		DefaultButton: "No",
		CancelButton:  "No",
	})

	if err != nil {
		return a.BeforeClose(ctx)
	}

	if option == "Yes" {
		prevent = false
		return prevent
	}

	prevent = true
	return prevent
}

// Shutdown is called at application termination
func (a *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
	a.TryClosePeerConnection()
	close(triggerEnd)
}

func (a *App) NotifyCreateClient() {

	openPeerMutex.Lock()
	defer openPeerMutex.Unlock()

	openPeer = true
	println("NotifyCreateClient")
}

func (a *App) NotifyCloseClient() {

	openPeerMutex.Lock()
	defer openPeerMutex.Unlock()

	openPeer = false
	println("NotifyCloseClient")
}

// Create a Host Peer, it receives the offer encoded and returns the encoded answer response
func (a *App) TryCreateHost(ICEServers []webrtc.ICEServer, offerEncoded string) (value string) {

	openPeerMutex.Lock()
	defer openPeerMutex.Unlock()

	if openPeer {
		triggerEnd <- struct{}{}
	}

	openPeer = true

	defer func() {

		if err := recover(); err != nil {

			log.Println(err)

			openPeerMutex.Lock()
			defer openPeerMutex.Unlock()
			openPeer = false
			value = "ERROR"
		}

	}()

	answerResponse := make(chan string)

	go net.InitHost(a.ctx, ICEServers, offerEncoded, answerResponse, triggerEnd)

	return <-answerResponse

}

// Closes the peer connection and returns a boolean indication if a connection existed and was closed or not
func (a *App) TryClosePeerConnection() bool {

	openPeerMutex.Lock()
	defer openPeerMutex.Unlock()

	if !openPeer {
		return false
	}

	triggerEnd <- struct{}{}

	openPeer = false

	return true

}

func (a *App) ToogleGamepad() {
	gamepad.GamepadEnabled.Toogle()
}

func (a *App) IsGamepadEnabled() bool {
	return gamepad.GamepadEnabled.IsEnabled()
}

func (a *App) ToogleKeyboard() {
	keyboard.KeyboardEnabled.Toogle()
}

func (a *App) IsKeyboardEnabled() bool {
	return keyboard.KeyboardEnabled.IsEnabled()
}

func (a *App) GetCurrentOS() string {
	return strings.ToUpper(runtime.GOOS)
}

func (a *App) LogPrintln(info string) {
	log.Println(info)
}
