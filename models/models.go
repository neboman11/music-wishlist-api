package models

type Album struct {
	Artist string `json:"artist"`
	Album  string `json:"album"`
}

type Want struct {
	Artist        string `json:"Artist"`
	Album         string `json:"Album"`
	Year          int    `json:"Year"`
	CoverArt_Link string `json:"CoverArt_Link"`
}

type MusicBrainzResponse struct {
	Releases []struct {
		Id string `json:"id"`
	} `json:"releases"`
}

type CoverResponse struct {
	CoverArt_Link string `json:"cover"`
}

type DeleteAlbumRequest struct {
	Albums []Album `json:"albums"`
}
