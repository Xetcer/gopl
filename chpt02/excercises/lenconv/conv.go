package lenconv

func MToCentim(m Metr) Centimeter { return Centimeter(m * 100) }
func MToMilim(m Metr) Milimeter   { return Milimeter(m * 1000) }

func CentimToMeter(m Centimeter) Metr { return Metr(m / 100) }
func MilimToMeter(m Milimeter) Metr   { return Metr(m / 1000) }
