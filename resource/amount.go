package resource

import "math/big"

type int64Amount struct {
	value int64
	scale Scale
}

type Scale int32

type infDecAmount struct {
	*Dec
}

type Dec struct {
	unscaled big.Int
	scale Scale
}




