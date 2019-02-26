package resource

type Quantity struct {
	i int64Amount
	d infDecAmount
	s string

	Format
}

type Format string
