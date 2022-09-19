package models

// TLocalization contains an object of the supported localization in Yearn
type TLocalization struct {
	En TLocalizationDetails `json:"en"`
	Fr TLocalizationDetails `json:"fr"`
	Es TLocalizationDetails `json:"es"`
	De TLocalizationDetails `json:"de"`
	Pt TLocalizationDetails `json:"pt"`
	El TLocalizationDetails `json:"el"`
	Tr TLocalizationDetails `json:"tr"`
	Vi TLocalizationDetails `json:"vi"`
	Zh TLocalizationDetails `json:"zh"`
	Hi TLocalizationDetails `json:"hi"`
	Ja TLocalizationDetails `json:"ja"`
	Id TLocalizationDetails `json:"id"`
	Ru TLocalizationDetails `json:"ru"`
}

// TLocalizationDetails contains the localization details for a specific language
type TLocalizationDetails struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
}
