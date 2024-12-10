package lenconv

import "fmt"

type Metr float64
type Centimeter float64
type Milimeter float64

func (m Metr) String() string       { return fmt.Sprintf("%g ,(m)", m) }
func (m Centimeter) String() string { return fmt.Sprintf("%g ,(m)", m) }
func (m Milimeter) String() string  { return fmt.Sprintf("%g ,(m)", m) }
