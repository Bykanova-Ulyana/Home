package plumbing

// сантехника

type Plumbung struct {
	Length   int
	Width    int
	Height   int
	Color    string
	Material string
	Features string
}

type Bath struct {
	Bath Plumbung
}

type Sink struct { // раковина
	Bath Plumbung
}

type Toilet struct {
	Bath Plumbung
}

type Faucet struct { // смеситель
	Bath Plumbung
}

type Shower struct {
	Bath Plumbung
}
