package main

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/creack/pty"
)

const MaxBufferSize = 16

func main() {

	a := app.New()
	w := a.NewWindow("germ")
	w.Resize(fyne.NewSize(900, 325))

	ui := widget.NewTextGrid()

	os.Setenv("TERM", "dumb")

	cmd := exec.Command("/bin/bash")

	ptmx, err := pty.Start(cmd)
	if err != nil {
		fyne.LogError("Failed to start PTY", err)
		os.Exit(1)
	}

	defer func() {
		_ = ptmx.Close()
		if cmd.Process != nil {
			_ = cmd.Process.Kill()
		}
	}()

	w.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {

		switch e.Name {

		case fyne.KeyEnter, fyne.KeyReturn:
			_, _ = ptmx.Write([]byte{'\r'})

		case fyne.KeyBackspace:
			_, _ = ptmx.Write([]byte{127})

		case fyne.KeyTab:
			_, _ = ptmx.Write([]byte{'\t'})
		}
	})

	w.Canvas().SetOnTypedRune(func(r rune) {
		_, _ = ptmx.Write([]byte(string(r)))
	})

	var (
		buffer [][]rune
		mu     sync.RWMutex
	)

	reader := bufio.NewReader(ptmx)

	go func() {

		line := []rune{}

		mu.Lock()
		buffer = append(buffer, line)
		mu.Unlock()

		for {

			r, _, err := reader.ReadRune()

			if err != nil {

				if err == io.EOF {
					return
				}

				fyne.LogError("PTY read", err)
				return
			}

			mu.Lock()

			line = append(line, r)
			buffer[len(buffer)-1] = line

			if r == '\n' {

				if len(buffer) >= MaxBufferSize {
					buffer = buffer[1:]
				}

				line = []rune{}
				buffer = append(buffer, line)
			}

			mu.Unlock()
		}

	}()

	go func() {

		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {

			mu.RLock()

			var text []rune

			for _, line := range buffer {
				text = append(text, line...)
			}

			s := string(text)

			mu.RUnlock()

			fyne.Do(func() {
				ui.SetText(s)
			})
		}

	}()

	w.SetContent(container.NewScroll(ui))

	w.ShowAndRun()
}
