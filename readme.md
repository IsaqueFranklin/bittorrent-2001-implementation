**Writing a BitTorrent client from the ground up in Go.**

BitTorrent is a protocol for downloading and distributing files across the internet. In contrast with the traditional client/server relantioship, in which downloaders connect to a central server (for example: watching a movie on Netflix, or loading a web page) participants in the BitTorrent network, called peers, download pieces of files from each other—this is what makes it a peer-to-peer protocol.

Here we'll implement the original spec from 2001 for peer-to-peer BitTorrent.
