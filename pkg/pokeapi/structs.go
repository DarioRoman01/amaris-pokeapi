package pokeapi

type PokemonForm struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	Order        int          `json:"order"`
	FormOrder    int          `json:"form_order"`
	IsDefault    bool         `json:"is_default"`
	IsBattleOnly bool         `json:"is_battle_only"`
	IsMega       bool         `json:"is_mega"`
	FormName     string       `json:"form_name"`
	Pokemon      interface{}  `json:"pokemon"`
	Sprites      Sprites      `json:"sprites"`
	Types        []Type       `json:"types"`
	VersionGroup VersionGroup `json:"version_group"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Sprites struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type Type struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}

type VersionGroup struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
