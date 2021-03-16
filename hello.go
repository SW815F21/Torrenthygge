package main

import "log"
import "os"
import "github.com/anacrolix/torrent"
import "strconv"

func main() {
	cf := torrent.NewDefaultClientConfig()
	cf.ListenPort, _ = strconv.Atoi(os.Getenv("TORRENT_CLIENT_PORT"))
	c, _ := torrent.NewClient(cf)
	defer c.Close()
	t, _ := c.AddTorrentFromFile("denmark-latest.osm.pbf.torrent")
	//Kan også laves som c.AddMagnetLink("bedstemagnetlink");
	<-t.GotInfo()
	t.DownloadAll()
	c.WaitAll()
	log.Println("ladies and gentlemen, we got him")
	for true {
		c.WriteStatus(log.Writer())
	}
}
