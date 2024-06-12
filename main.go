//Trackers are central servers that introduce peers to each other.
//A .torrent file describes the contents of a torrentable file and information for connecting to a tracker. It’s all we need in order to kickstart the process of downloading a torrent.

//It is using Bencode for the .torrent file.

import (
  "github.com/jackpal/bencode-go"
)

type bencodeInfo struct {
    Pieces      string `bencode:"pieces"`
    PieceLength int    `bencode:"piece length"`
    Length      int    `bencode:"length"`
    Name        string `bencode:"name"`
}

type bencodeTorrent struct {
   Announce string      `bencode:"announce"`
  Info     bencodeInfo `bencode:"info"`
}

//Open parses a torrent file
func Open(r io.Reader) (*bencodeTorrent, error) {
  bto := bencodeTorrent{}
  err := bencode.Umarshal(r, &bto)
  if err != nil {
    return nil, err
  }

  return &bto, nil
}
