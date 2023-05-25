package uploadImage

type Images []Image

type Image struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
