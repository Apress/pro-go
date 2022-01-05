package store

type Crew struct {
    Captain, FirstOfficer string
}

type RentalBoat struct {
    *Boat
    IncludeCrew bool
    *Crew
}

func NewRentalBoat(name string, price float64, capacity int, 
        motorized, crewed bool, captain, firstOfficer string) *RentalBoat {
    return &RentalBoat{NewBoat(name, price, capacity, motorized), crewed, 
        &Crew{captain, firstOfficer}}
}
