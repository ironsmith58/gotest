package pasta

type Pasta int8

const (
    Elbow Pasta = iota + 1
    Macaroni
    SewerPipe
    Spagatti
    Rigatoni
)



//go:generate stringer -type=Pasta
