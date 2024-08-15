package models

import "encoding/json"

type Profile struct {
	ID          int    `json:"id"`
	ID1C        string `json:"id_1c"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Profiles    string `json:"profiles"`
	Description string `json:"description"`
	Department  string `json:"department"`
	MinScore    int    `json:"min_score"`
	DegreeID    int    `json:"degree_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	InstituteID int    `json:"institute_id"`
	Keywords    string `json:"keywords"`
	Slug        string `json:"slug"`
	LogoFile    string `json:"logo_file"`
	LogoURL     string `json:"logo_url"`
}

type ProfileResponse []Profile

func (p *ProfileResponse) Parse(data []byte) error {
	return json.Unmarshal(data, p)
}
