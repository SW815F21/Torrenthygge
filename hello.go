package main

import "fmt"

import "github.com/anacrolix/torrent"

func main() {
	c, _ := torrent.NewClient(nil)
	defer c.Close()
	t, _ := c.AddTorrentFromFile("ly.torrent");
	//Kan også laves som c.AddMagnetLink("bedstemagnetlink");
	<-t.GotInfo()
	t.DownloadAll()
	c.WaitAll()
	fmt.Println("ladies and gentlemen, we got him")
}
