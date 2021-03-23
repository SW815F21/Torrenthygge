package main

import "log"
import "fmt"
import "os"
import "github.com/anacrolix/torrent"
import "strconv"
import "time"

func main() {
	if os.Args[1] == "" {
		fmt.Println("Please specify a torrent file as parameter")
		os.Exit(1)
	}
	torrFile := os.Args[1]
	cf := torrent.NewDefaultClientConfig()
	cf.ListenPort, _ = strconv.Atoi(os.Getenv("TORRENT_CLIENT_PORT"))
	c, _ := torrent.NewClient(cf)
	defer c.Close()
	t, _ := c.AddTorrentFromFile(torrFile)
	torrentBar(t, false)
	//Kan også laves som c.AddMagnetLink("bedstemagnetlink");
	<-t.GotInfo()
	t.DownloadAll()
	c.WaitAll()
	log.Println("ladies and gentlemen, we got him")
}

func torrentBar(t *torrent.Torrent, pieceStates bool) {
	go func() {
		if t.Info() == nil {
			fmt.Printf("Getting info for %q\n", t.Name())
			<-t.GotInfo()
		}
		var lastLine string
		for {
			var completedPieces, partialPieces int
			psrs := t.PieceStateRuns()
			for _, r := range psrs {
				if r.Complete {
					completedPieces += r.Length
				}
				if r.Partial {
					partialPieces += r.Length
				}
			}
			line := fmt.Sprintf(
				"Downloading %q: %d/%d, %d/%d pieces completed (%d partial)\n",
				t.Name(),
				uint64(t.BytesCompleted()),
				uint64(t.Length()),
				completedPieces,
				t.NumPieces(),
				partialPieces,
			)

			if line != lastLine {
				lastLine = line
				os.Stdout.WriteString(line)
			}
			if pieceStates {
				fmt.Println(psrs)
			}
			time.Sleep(time.Second)
		}
	}()
}
