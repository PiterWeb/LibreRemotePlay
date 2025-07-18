package audio

import (
	"context"
	"fmt"
	"time"

	"github.com/PiterWeb/RemoteController/src/bin"
	"github.com/amenzhinsky/go-memexec"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
	"github.com/go-ole/go-ole"
	"github.com/moutend/go-wca/pkg/wca"
)

type AudioProcess struct {
	Name string
	Pid  uint32
}

func HandleAudio(ctx context.Context, track *webrtc.TrackLocalStaticSample) error {
	
	appLoopbackExe, err := memexec.New(bin.StdoutPCMApplicationLoopback_exe)
	if err != nil {
		return err
	}
	defer appLoopbackExe.Close()

	pid := ctx.Value("pid").(uint32)

	cmd := appLoopbackExe.Command(fmt.Sprintf("%d", pid))

	stdoutPipe, err := cmd.StdoutPipe()

	if err != nil {
		return  err
	}

	buff := make([]byte, 1024)

	tickerChan := time.NewTicker(time.Millisecond * 20).C

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-tickerChan:
			if n, err := stdoutPipe.Read(buff); err != nil {
				return err
			} else if n == 0 {
				continue
			}
	
			if err = track.WriteSample(media.Sample{Data: buff, Duration: time.Millisecond * 20}); err != nil {
				return err
			}
		}
	}

}

func GetAudioProcess() []AudioProcess {

	procs := []AudioProcess{
		{
			Name: "None",
			Pid: 0,
		},
	}

	if err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return procs
	}
	defer ole.CoUninitialize()

	var mmde *wca.IMMDeviceEnumerator
	if err := wca.CoCreateInstance(wca.CLSID_MMDeviceEnumerator, 0, wca.CLSCTX_ALL, wca.IID_IMMDeviceEnumerator, &mmde); err != nil {
		return procs
	}
	defer mmde.Release()

	var mmd *wca.IMMDevice
	if err := mmde.GetDefaultAudioEndpoint(wca.ERender, wca.EConsole, &mmd); err != nil {
		return procs
	}

	var asm2 *wca.IAudioSessionManager2
	if err := mmd.Activate(wca.IID_IAudioSessionManager2, ole.CLSCTX_ALL, nil, &asm2); err != nil {
		return procs
	}

	defer asm2.Release()

	var ase *wca.IAudioSessionEnumerator
	if err := asm2.GetSessionEnumerator(&ase); err != nil {
		return procs
	}

	defer ase.Release()

	var sessionNumber int
	if err := ase.GetCount(&sessionNumber); err != nil {
		return procs
	}

	fmt.Printf("Sessions: %d\n", sessionNumber)


	for i := range sessionNumber {
		var asc *wca.IAudioSessionControl
		if err := ase.GetSession(i, &asc); err != nil {
			continue
		}

		defer asc.Release()

		var asc2 *wca.IAudioSessionControl2
		if err := asc.PutQueryInterface(wca.IID_IAudioSessionControl2, &asc2); err != nil {
			continue
		}

		defer asc2.Release()

		var name string
		if err := asc.GetDisplayName(&name); err != nil {
			continue
		}

		var pid uint32 = 0
		if err := asc2.GetProcessId(&pid); err != nil {
			continue
		}

		procs = append(procs, AudioProcess{
			Name: name,
			Pid:  pid,
		})

	}

	return procs
}