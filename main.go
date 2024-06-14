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

type TorrentFile struct {
  Announce string
  infoHash [20]byte
  PieceHashes [][20]byte
  PieceLength int
  Length int
  Name string
}

func (bto bencodeTorrent) toTorrentFile() (TorrentFile, error){

}

func (t *TorrentFile) buildTrackerURL(peerID [20]byte, port uint16)(string, error){
  base, err := url.Parse(t.Announce)

  if err != nil {
    return "", err
  }

  params := url.Values{
    "info_hash":  []string{string(t.InfoHash[:])},
    "peer_id":    []string{string(peerID[:])},
    "port":       []string{strconv.Itoa(int(Port))},
    "uploaded":   []string{"0"},
    "downloaded": []string{"0"},
    "compact":    []string{"1"},
    "left":       []string{strconv.Itoa(t.Length)},
  }
  base.RawQuery = params.Encode()
  return base.String(), nil
}

type Peer struct {
  IP net.IP
  Port uint16
}

//Unmarshal parses peer IP addres and ports from a buffer
func Unmarshal(peersByn []byte) ([]Peer, error) {
  const peerSize = 6
  numPeers := len(peersBin) / peerSize
  if len(peersBin)%peerSize != 0 {
    err := fmt.Errorf("Received malformed peers")
    return nil, err
  }

  peers := make([]Peer, numPeers)
  for i:=0; i < numPeers; i++ {
    offset := i * peerSize
    peers[i].IP = net.IP(peersBin[offset: offset+4])
    peers[i].Port = binary.BigEndian.Uint16(peersBin[offset+4 : offset+6])
  }

  return peers, nil
}
