package expectedruns

type Stat struct {
	Id            int // Binary representation of the runners on base
	ScoringChance float64
	ExpectedRuns  float64
}

type Out struct {
	FriendlyName string
	Stats        []Stat
}

var Stats = []Out{
	{
		FriendlyName: "No Outs",
		Stats: []Stat{
			{
				Id:            0,
				ScoringChance: 27.1,
				ExpectedRuns:  0.491,
			},
			{
				Id:            1,
				ScoringChance: 40.2,
				ExpectedRuns:  0.84,
			},
			{
				Id:            2,
				ScoringChance: 60.2,
				ExpectedRuns:  1.084,
			},
			{
				Id:            3,
				ScoringChance: 60.3,
				ExpectedRuns:  1.433,
			},
			{
				Id:            4,
				ScoringChance: 85.1,
				ExpectedRuns:  1.418,
			},
			{
				Id:            5,
				ScoringChance: 85.4,
				ExpectedRuns:  1.775,
			},
			{
				Id:            6,
				ScoringChance: 84.2,
				ExpectedRuns:  1.936,
			},
			{
				Id:            7,
				ScoringChance: 85.9,
				ExpectedRuns:  2.278,
			},
		},
	}, {
		FriendlyName: "1 Out",
		Stats: []Stat{
			{
				Id:            0,
				ScoringChance: 15.8,
				ExpectedRuns:  0.263,
			},
			{
				Id:            1,
				ScoringChance: 25,
				ExpectedRuns:  0.493,
			},
			{
				Id:            2,
				ScoringChance: 38.9,
				ExpectedRuns:  0.663,
			},
			{
				Id:            3,
				ScoringChance: 39.6,
				ExpectedRuns:  0.877,
			},
			{
				Id:            4,
				ScoringChance: 65.2,
				ExpectedRuns:  0.939,
			},
			{
				Id:            5,
				ScoringChance: 64.2,
				ExpectedRuns:  1.114,
			},
			{
				Id:            6,
				ScoringChance: 67,
				ExpectedRuns:  1.367,
			},
			{
				Id:            7,
				ScoringChance: 65.1,
				ExpectedRuns:  1.524,
			},
		},
	},
	{
		FriendlyName: "2 Outs",
		Stats: []Stat{
			{
				Id:            0,
				ScoringChance: 6.8,
				ExpectedRuns:  0.101,
			},
			{
				Id:            1,
				ScoringChance: 11.8,
				ExpectedRuns:  0.216,
			},
			{
				Id:            2,
				ScoringChance: 21.5,
				ExpectedRuns:  0.321,
			},
			{
				Id:            3,
				ScoringChance: 22.4,
				ExpectedRuns:  0.436,
			},
			{
				Id:            4,
				ScoringChance: 24.2,
				ExpectedRuns:  0.337,
			},
			{
				Id:            5,
				ScoringChance: 28.9,
				ExpectedRuns:  0.498,
			},
			{
				Id:            6,
				ScoringChance: 26,
				ExpectedRuns:  0.577,
			},
			{
				Id:            7,
				ScoringChance: 33,
				ExpectedRuns:  0.786,
			},
		},
	},
}
