// Binding for JS to Go
// This package is responsible for the communication between the JS and Go code.
package bindings

import (
	"context"
	"embed"
	"log"
	"strings"
	"sync"

	"runtime"

	"github.com/PiterWeb/RemoteController/src/devices/audio"
	"github.com/PiterWeb/RemoteController/src/devices/gamepad"
	"github.com/PiterWeb/RemoteController/src/devices/keyboard"
	net "github.com/PiterWeb/RemoteController/src/net/webrtc"
	"github.com/PiterWeb/RemoteController/src/net/webrtc/streaming_signal"
	"github.com/PiterWeb/RemoteController/src/onfinish"
	"github.com/PiterWeb/RemoteController/src/oninit"
	"github.com/pion/webrtc/v4"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

var triggerEnd chan struct{} = make(chan struct{})
var pidAudioChan chan uint32 = make(chan uint32)

var openPeer bool = false
var openPeerMutex sync.Mutex

// App struct
type App struct {
	ctx    context.Context
	assets embed.FS
}

// NewApp creates a new App application struct
func NewApp(assets embed.FS) *App {
	return &App{
		ctx:    context.Background(),
		assets: assets,
	}
}

// Startup is called at application Startup
func (a *App) Startup(ctx context.Context) {
	// Perform your setup here

	go func() {

		if err := oninit.Execute(a.assets); err != nil {
			log.Println(err)
		}

	}()
	
	go func () {
		if err := streaming_signal.InitWhipServer(streaming_signal.WhipConfig); err != nil {
			log.Println(err)
		}
	}()

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
	if err := onfinish.Execute(); err != nil {
		log.Printf("Error onfinish: %s", err.Error())
	}
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

	answerResponse := make(chan string)

	go net.InitHost(a.ctx, ICEServers, offerEncoded, answerResponse, triggerEnd, pidAudioChan)

	response := <-answerResponse

	if strings.Contains(response, "ERROR") {
		openPeerMutex.Lock()
		defer openPeerMutex.Unlock()
		openPeer = false
		log.Println("Error on WebRTC host connection")
	}

	return response

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

func (a *App) ToogleWhip() {
	streaming_signal.WhipConfig.Enabled.Toogle()
}

func (a *App) IsWhipEnabled() bool {
	return streaming_signal.WhipConfig.Enabled.IsEnabled()
}

func (a *App) GetCurrentOS() string {
	return strings.ToUpper(runtime.GOOS)
}

func (a *App) LogPrintln(info string) {
	log.Println(info)
}

func (a *App) SetAudioPid(pid uint32) {
	pidAudioChan <- pid
}

func (a *App) GetAudioProcess() []audio.AudioProcess {
	return audio.GetAudioProcess()
}
