package main

func calculateNextState(p golParams, world [][]byte) [][]byte {
	y := 0
	x := 0
	sum := 0
	newWorld := make([][]byte, p.imageHeight)
	for i := range newWorld {
		newWorld[i] = make([]byte, p.imageWidth)
	}

	for y = 0; y < p.imageHeight; y++ {
		for x = 0; x < p.imageWidth; x++ {
			sum = int(world[(y+p.imageHeight-1)%p.imageHeight][(x+p.imageWidth-1)%p.imageWidth]/255 +
				world[(y+p.imageHeight-1)%p.imageHeight][(x+p.imageWidth)%p.imageWidth]/255 +
				world[(y+p.imageHeight-1)%p.imageHeight][(x+p.imageWidth+1)%p.imageWidth]/255 +
				world[(y+p.imageHeight)%p.imageHeight][(x+p.imageWidth-1)%p.imageWidth]/255 +
				world[(y+p.imageHeight)%p.imageHeight][(x+p.imageWidth+1)%p.imageWidth]/255 +
				world[(y+p.imageHeight+1)%p.imageHeight][(x+p.imageWidth-1)%p.imageWidth]/255 +
				world[(y+p.imageHeight+1)%p.imageHeight][(x+p.imageWidth)%p.imageWidth]/255 +
				world[(y+p.imageHeight+1)%p.imageHeight][(x+p.imageWidth+1)%p.imageWidth]/255)

			if (world[y][x]) == 255 {
				if sum < 2 {
					newWorld[y][x] = 0
				} else if sum == 2 || sum == 3 {
					newWorld[y][x] = 255
				} else {
					newWorld[y][x] = 0
				}
			} else {
				if sum == 3 {
					newWorld[y][x] = 255
				} else {
					newWorld[y][x] = 0
				}
			}
		}
	}
	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	aliveCells := []cell{} // Initialize as an empty slice

	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			if world[y][x] == 255 {
				// Append to the aliveCells slice
				aliveCells = append(aliveCells, cell{y: y, x: x})
			}
		}
	}
	return aliveCells // Return the populated slice
}
