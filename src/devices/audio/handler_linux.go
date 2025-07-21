package audio

import (
	"context"
	"io"
	"log"
	"math"
	"os/exec"
	"strings"

	"github.com/pion/webrtc/v3"
)

func HandleAudio(ctx context.Context, track *webrtc.TrackLocalStaticSample) error {
	return nil
}

func GetAudioProcess() []AudioProcess {

	procs := []AudioProcess{
		{
			Name: "None",
			Pid:  0,
		},
	}

	cmdExistsPipeWire := exec.Command("command", "-v", "pw-cli")

	if outputExists, err := cmdExistsPipeWire.Output(); err != nil {
		log.Println("Error - Pipe Wire is not installed")
		return procs
	} else if len(outputExists) == 0 {
		log.Println("Error - Pipe Wire is not installed")
		return procs
	}

	cmdPipeWireList := exec.Command("pw-cli", "ls", "Node")

	pipeWireListStdOutPipe, err := cmdPipeWireList.StdoutPipe()

	if err != nil {
		log.Println("Error - unexpected error on pw-cli stdin pipe")
		return procs
	}

	cmdGrepFilterApplications := exec.Command("grep", "application.name = ")

	grepFilterStdInPipe, err := cmdGrepFilterApplications.StdinPipe()

	if err != nil {
		log.Println("Error - unexpected error on grep stdin pipe")
		return procs
	}

	go io.Copy(grepFilterStdInPipe, pipeWireListStdOutPipe)

	if err := cmdPipeWireList.Run(); err != nil {
		log.Printf("Error - Pipe Wire command `%s` is not working\n", strings.Join(cmdPipeWireList.Args, " "))
		return procs
	}

	pipeWireListRaw, err := cmdGrepFilterApplications.Output()

	if err != nil {
		log.Printf("Error - grep command `%s` is not working\n", strings.Join(cmdGrepFilterApplications.Args, " "))
		return procs
	}

	// Replace " with void, then split the string in \n to get every application
	pipeWireList := strings.Split(strings.ReplaceAll(string(pipeWireListRaw), "\"", ""), "\n")

	for i, appNamePipeWire := range pipeWireList {

		if (i + 1) > math.MaxUint32 {
			return procs
		}

		procs = append(procs, AudioProcess{
			Name: appNamePipeWire,
			Pid:  uint32(i + 1), // It is not a real Pid, it is only array index
		})
	}

	return procs
}
