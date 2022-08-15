package Model

type Berita struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Judul       string `gorm:"notNull;size:255" json:"judul"`
	Gambar      string `json:"gambar"`
	Konten      string `gorm:"type:longtext;notNull" json:"konten"`
	IsPublished bool   `gorm:"default:false;notNull" json:"is_published"`
}
