package Percolation

import "math"

type PercolationStats struct {
	*Percolation

	mean     float64
	standDev float64
}

func (p *PercolationStats) PercolationStat(n, c int) {
	sum := float64(0)
	ratioSlice := []float64{}
	total := float64(n * n)

	p.init(n)

	for i := 0; i < c; i++ {
		for {
			row := random(1, n+1)
			col := random(1, n+1)

			p.open(row, col)

			if p.percolates() {
				openSiteNum := float64(p.openSiteNum)
				ratio := openSiteNum / total

				ratioSlice = append(ratioSlice, ratio)
				sum += ratio

				break
			}
		}
	}

	p.mean = sum / float64(c)

	tmp := float64(0)
	for _, v := range ratioSlice {
		tmp += math.Pow(v-p.mean, 2)
	}
	p.standDev = math.Sqrt(tmp / float64(c-1))
}

func (p *PercolationStats) StandDev() float64 {
	return p.standDev
}

func (p *PercolationStats) Mean() float64 {
	return p.mean
}
