package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"io"
	"net"
	"os"
	"time"
)

type transWin struct {
	progress *widget.ProgressBar
	btnCot   *widget.Button
	win      fyne.Window
	temFlag  int
}

func sendFile(u *user, fileName string, s *Server) {
	t := &transWin{
		progress: widget.NewProgressBar(),
		win:      s.win.App.NewWindow("file transferring"),
		temFlag:  0,
	}
	t.btnCot = widget.NewButton("start", cotinue(t))

	go sendf(u, fileName, s, t)
	t.win.SetContent(container.NewVBox(t.progress, t.btnCot))
	t.win.Resize(fyne.NewSize(600, 450))
	t.win.Show()

}

func sendf(u *user, fileName string, s *Server, t *transWin) {

	conn := u.conn
	fmt.Println(u.name, " is recving")

	if _, ok := u.fileList[fileName]; !ok {
		u.fileList[fileName] = 0
	}

	ofile, err := os.Open(fileName)
	fileSize, _ := ofile.Seek(0, os.SEEK_END)
	ofile.Seek(int64(u.fileList[fileName]), os.SEEK_SET)

	fmt.Println("filesize", fileSize, ", filename", fileName)
	defer ofile.Close()

	if err != nil {
		dialog.ShowError(err, s.win.Win)
		fmt.Println("failed to open the file : ", fileName, err)
		return
	}

	buf := make([]byte, 4096)
	for {
		t.progress.SetValue(float64(u.fileList[fileName]) / float64(fileSize))
		for ; t.temFlag == 2; time.Sleep(time.Millisecond * 10) {
		}
		n, err := ofile.Read(buf)
		if err != nil {
			if err == io.EOF {
				dialog.ShowInformation("congratulations", "Send Success", t.win)
				fmt.Println("send success")
				t.btnCot.SetText("close")
				t.temFlag = 3
				u.fileList[fileName] = 0
			} else {
				dialog.ShowError(err, t.win)
				fmt.Println("file read err: ", err)
			}

			return
		}

		_, err = conn.Write(buf[:n])
		u.fileList[fileName] += n

	}

}

func cotinue(t *transWin) func() {
	return func() {
		if t.temFlag == 0 {
			t.btnCot.SetText("transferring")
			t.temFlag = 1
		} else if t.temFlag == 1 {
			t.btnCot.SetText("stopping")
			t.temFlag = 2
		} else if t.temFlag == 2 {
			t.btnCot.SetText("continue")
			t.temFlag = 1
		} else if t.temFlag == 3 {
			t.win.Close()
		}
	}
}

func recvFile(conn net.Conn, fileName string) {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("file create err: ", err)
		return
	}

	buf := make([]byte, 4096)
	for {

		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				//dialog.ShowInformation("congratulation","","",nil)
				fmt.Println(fileName, "accept success")
			} else {
				fmt.Println("recv error: ", err)
			}
			return
		}

		_, err = file.Write(buf[:n])

		if err != nil {

			return
		}

	}

}
