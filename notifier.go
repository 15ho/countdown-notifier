package main

import (
	"os"
	"os/exec"
)

const cmd = "terminal-notifier"

func checkTerminalNotifier() {
	_, err := exec.
		Command(cmd, "-version").
		Output()
	if err != nil {
		stderr(`Need install terminal-notifier!!!
If you have brew. you can: 'brew install terminal-notifier'
Install doc: https://github.com/julienXX/terminal-notifier`)
		os.Exit(2)
	}
}

func notice() {
	_ = exec.Command(cmd,
		"-sound", "default",
		"-title", "Countdown-Notifier",
		"-message", "Finished",
		"-open", "https://github.com/15ho/countdown-notifier").
		Run()
}
