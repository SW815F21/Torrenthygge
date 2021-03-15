package main

import "log"

import "github.com/anacrolix/torrent"

func main() {
	c, _ := torrent.NewClient(nil)
	defer c.Close()
	t, _ := c.AddTorrentFromFile("denmark-latest.osm.pbf.torrent");
	//Kan også laves som c.AddMagnetLink("bedstemagnetlink");
	<-t.GotInfo()
	t.DownloadAll()
	c.WaitAll()
	log.Println("ladies and gentlemen, we got him")
	for true {
		c.WriteStatus(log.Writer())
	}
}
