package darts

// Score takes in the x and y axis of the darts on the 
// board and returns points for the particular axis
func Score(x, y float64) int {
    conditions := []struct{
        maxRadiusSquared float64
        points int
    }{{1.0, 10}, {25.0, 5}, {100.0, 1}}
	radiusSquared := x*x + y*y //pythagoras theorem in concentric circles

    for _, condition := range conditions{
        if radiusSquared <= condition.maxRadiusSquared{
            return condition.points
        } 
    }
    return 0
}
