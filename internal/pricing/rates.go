package pricing

type Rates struct{ Daily, Weekly int }

func Lookup() Rates { return Rates{Daily: 50, Weekly: 300} }
