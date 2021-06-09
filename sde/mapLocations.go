package sde

type System struct {
	SystemID        int
	ConstellationID int
	RegionID        int
	Name            string
	Security        float32
	FactionID       *int
}

type Constellation struct {
	ConstellationID int
	RegionID        int
	Name            string
	FactionID       *int
}

type Region struct {
	RegionID       int
	Name           string
	FactionID      *int
	Systems        map[string]System // Not used by default, putting the shell in place as I'm 100% sure I'll need this later
	Constellations map[string]Constellation
}
