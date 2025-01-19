Iterator Design Pattern - Playlist Example
This project demonstrates the Iterator Design Pattern using a music playlist to traverse songs sequentially without exposing internal details.

Components
Song: Represents an individual song.
PlayList: Manages a collection of songs with methods:
AddSong(song Song): Adds songs to the playlist.
CreateIterator(): Returns an iterator for traversal.
Iterator (Interface): Defines methods:
hasNext() bool: Checks for more elements.
next() *Song: Retrieves the next element.
PlayListIterator: Implements Iterator for the playlist.
