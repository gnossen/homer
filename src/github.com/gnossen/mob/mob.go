package mob

type MobClass struct {
	Name         string `json:name`
	Nick         string `json:nick`
	Strength     string `json:"strength"`
	Dexterity    string `json:"dexterity"`
	Constitution string `json:"constitution"`
	Intelligence string `json:"intelligence"`
	Wisdom       string `json:"wisdom"`
	Charisma     string `json:"charisma"`
	MaxHP        string `json:"max-hp"`
	MaxMana      string `json:"max-mana"`
	ArmorClass   string `json:"armor-class"`
}

type Mob struct {
	Name         string `json:name`
	Nick         string `json:nick`
	Class        string `json:class`
	Strength     int    `json:"strength"`
	Dexterity    int    `json:"dexterity"`
	Constitution int    `json:"constitution"`
	Intelligence int    `json:"intelligence"`
	Wisdom       int    `json:"wisdom"`
	Charisma     int    `json:"charisma"`
	MaxHP        int    `json:"max-hp"`
	MaxMana      int    `json:"max-mana"`
	ArmorClass   int    `json:"armor-class"`
}

type ActiveMob struct {
	Mob
	HP   int `json:"hp"`
	Mana int `json:"mana"`
}
