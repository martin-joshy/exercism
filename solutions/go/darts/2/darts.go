package darts

var conditions = []struct{
    maxRadiusSquared int
    points int
}{
    {1, 10},
    {25, 5},
    {100, 1},
}

// Score takes in the x and y axis of the darts on the 
// board and returns points for the particular axis
func Score(x, y float64) int {
    
	radiusSquared := x*x + y*y // Pythagoras theorem in concentric circles

    for _, condition := range conditions{
        if radiusSquared <= float64(condition.maxRadiusSquared){
            return condition.points
        } 
    }
    return 0
}
