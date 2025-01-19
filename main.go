package main

import "fmt"

type Iterator interface {
	hasNext() bool
	next() *Song
}

type PlayListIterator struct {
	songlist []Song
	index    int
}

func (p *PlayListIterator) hasNext() bool {

	return (p.index < len(p.songlist))

}

func (p *PlayListIterator) next() *Song {
	if p.hasNext() {
		songs := p.songlist[p.index]
		p.index++
		return &songs
	}
	return nil
}

type Song struct {
	song string
}

//  func newSong() *Song {
//  	return &Song{}
// }

type PlayList struct {
	songs []Song
}

func (p *PlayList) AddSong(song Song) {

	p.songs = append(p.songs, song)

}

func (p *PlayList) CreateIterator() Iterator {
	return &PlayListIterator{
		songlist: p.songs,
		index:    0,
	}

}
func main() {
	s1 := Song{song: "lily "}
	s2 := Song{song: "Alone"}
	s3 := Song{song: "Deja vu"}
	playList := &PlayList{}
	playList.AddSong(s1)
	playList.AddSong(s2)
	playList.AddSong(s3)
	playList.AddSong(s1)
	// playList.AddSong("lily ")
	// playList.AddSong("alone")
	// playList.AddSong("deja vu")
	iterator := playList.CreateIterator()
	for iterator.hasNext() {
		fmt.Println("NOW PLAYING", iterator.next().song)
	}
}
