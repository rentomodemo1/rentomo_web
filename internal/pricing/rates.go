package pricing

type Rates struct{ Daily, Weekly, Annual int }

func Lookup() Rates { return Rates{Daily: 50, Weekly: 300, Annual: 14000} }
