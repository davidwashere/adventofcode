package day19

import (
	"aoc/util"
	"fmt"
	"math"
)

// type Stash struct {
// 	ore      int
// 	clay     int
// 	obsidian int
// 	geode    int

// 	oreRobots      int
// 	clayRobots     int
// 	obsidianRobots int
// 	geodeRobots    int
// }

// type BuildResult struct {
// 	oreRobots      int
// 	clayRobots     int
// 	obsidianRobots int
// 	geodeRobots    int
// }

// func (b *BuildResult) Apply(stash *Stash) {
// 	stash.oreRobots += b.oreRobots
// 	stash.clayRobots += b.clayRobots
// 	stash.obsidianRobots += b.obsidianRobots
// 	stash.geodeRobots += b.geodeRobots
// }

// type Cost struct {
// 	ore      int
// 	clay     int
// 	obsidian int
// }

// type RobotBuilder interface {
// 	Build(stash *Stash) *BuildResult
// }

// type defaultRobotBuilder struct {
// 	childBuilder   RobotBuilder
// 	cost           Cost
// 	onBuildSuccess func(*BuildResult)
// }

// func (b *defaultRobotBuilder) Build(stash *Stash) *BuildResult {
// 	cost := b.cost

// 	// if can't build, try child builder
// 	if stash.ore < cost.ore || stash.clay < cost.clay || stash.obsidian < cost.obsidian {
// 		// can't build, if no child builder return
// 		if b.childBuilder == nil {
// 			return nil
// 		}

// 		// otherwise pass on to child builder
// 		return b.childBuilder.Build(stash)
// 	}

// 	// can build it, lets do it
// 	stash.ore -= cost.ore
// 	stash.clay -= cost.clay
// 	stash.obsidian -= cost.obsidian

// 	result := new(BuildResult)
// 	b.onBuildSuccess(result)

// 	return result
// }

// var blueprints []RobotBuilder

// func load(inputFile string) {
// 	data, _ := util.ReadFileToStringSlice(inputFile)

// 	for _, line := range data {
// 		ints := util.ParseInts(line)
// 		oreOre := ints[1]
// 		clayOre := ints[2]
// 		obsidianOre := ints[3]
// 		obsidianClay := ints[4]
// 		geodeOre := ints[5]
// 		geodeObsidian := ints[6]

// 		fmt.Println(oreOre, clayOre, obsidianOre, obsidianClay, geodeOre, geodeObsidian)

// 		var cost Cost
// 		// ore builder
// 		cost = Cost{ore: oreOre}
// 		oreBuilder := &defaultRobotBuilder{cost: cost}
// 		oreBuilder.onBuildSuccess = func(r *BuildResult) { r.oreRobots += 1 }

// 		// clay builder
// 		cost = Cost{ore: clayOre}
// 		clayBuilder := &defaultRobotBuilder{cost: cost}
// 		clayBuilder.onBuildSuccess = func(r *BuildResult) { r.clayRobots += 1 }

// 		// obsidian builder
// 		cost = Cost{ore: obsidianOre, clay: obsidianClay}
// 		obsidianBuilder := &defaultRobotBuilder{cost: cost}
// 		obsidianBuilder.onBuildSuccess = func(r *BuildResult) { r.obsidianRobots += 1 }

// 		// build geode builder
// 		cost = Cost{ore: geodeOre, obsidian: geodeObsidian}
// 		geodeBuilder := &defaultRobotBuilder{cost: cost}
// 		geodeBuilder.onBuildSuccess = func(r *BuildResult) { r.geodeRobots += 1 }

// 		geodeBuilder.childBuilder = obsidianBuilder
// 		obsidianBuilder.childBuilder = clayBuilder
// 		clayBuilder.childBuilder = oreBuilder

// 		blueprints = append(blueprints, geodeBuilder)
// 	}
// }

// func robotsCollect(stash *Stash) {
// 	stash.ore += stash.oreRobots
// 	stash.clay += stash.clayRobots
// 	stash.obsidian += stash.obsidianRobots
// 	stash.geode += stash.geodeRobots
// }

// func part1(inputFile string) int {
// 	load(inputFile)
// 	maxMins := 24

// 	stashes := []*Stash{}
// 	for _, blueprint := range blueprints {
// 		stash := new(Stash)
// 		stash.oreRobots = 1
// 		stashes = append(stashes, stash)

// 		for min := 1; min <= maxMins; min++ {
// 			result := blueprint.Build(stash)

// 			robotsCollect(stash)

// 			if result != nil {
// 				result.Apply(stash)
// 			}
// 		}
// 	}

// 	return 0
// }

const (
	OreRobot = iota
	ClayRobot
	ObsidianRobot
	GeodeRobot
)

type Stash struct {
	ore      int
	clay     int
	obsidian int
	geode    int

	robots [4]int
}

func (s Stash) String() string {
	return fmt.Sprintf("%v ore %v clay %v ob %v geo - robots: %v", s.ore, s.clay, s.obsidian, s.geode, s.robots)
}

func (s Stash) BuildRobot(cost Cost, i int) Stash {
	robots := s.robots
	robots[i]++
	stash := Stash{
		ore:      s.ore - cost.ore,
		clay:     s.clay - cost.clay,
		obsidian: s.obsidian - cost.obsidian,
		geode:    s.geode,

		robots: robots,
	}

	return stash
}

func (s Stash) RobotsDoWork() Stash {
	updStash := Stash{}

	updStash.ore = s.ore + s.robots[OreRobot]
	updStash.clay = s.clay + s.robots[ClayRobot]
	updStash.obsidian = s.obsidian + s.robots[ObsidianRobot]
	updStash.geode = s.geode + s.robots[GeodeRobot]

	robots := s.robots
	updStash.robots = robots

	return updStash
}

func NewStash() Stash {
	stash := Stash{}
	stash.robots = [4]int{}
	stash.robots[OreRobot] = 1

	return stash
}

type Cost struct {
	ore      int
	clay     int
	obsidian int
}

type Blueprint []Cost

var blueprints []Blueprint

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	for _, line := range data {
		ints := util.ParseInts(line)
		oreOre := ints[1]
		clayOre := ints[2]
		obsidianOre := ints[3]
		obsidianClay := ints[4]
		geodeOre := ints[5]
		geodeObsidian := ints[6]

		blueprint := Blueprint{}
		blueprint = append(blueprint, Cost{ore: oreOre})                            // ore builder
		blueprint = append(blueprint, Cost{ore: clayOre})                           // clay builder
		blueprint = append(blueprint, Cost{ore: obsidianOre, clay: obsidianClay})   // obsidian builder
		blueprint = append(blueprint, Cost{ore: geodeOre, obsidian: geodeObsidian}) // build geode builder

		blueprints = append(blueprints, blueprint)
	}
}

func canAfford(stash Stash, cost Cost) bool {
	if stash.ore < cost.ore || stash.clay < cost.clay || stash.obsidian < cost.obsidian {
		return false
	}

	return true
}

func part1(inputFile string) int {
	load(inputFile)
	maxMins := 24

	var recur func(mins int, stash Stash) int

	result := 0
	for i, blueprint := range blueprints {
		memo := map[string]int{}

		maxOrePerRound := util.MaxAll(blueprint[0].ore, blueprint[1].ore, blueprint[2].ore, blueprint[3].ore)
		maxClayPerRound := util.MaxAll(blueprint[0].clay, blueprint[1].clay, blueprint[2].clay, blueprint[3].clay)
		maxObsidianPerRound := util.MaxAll(blueprint[0].obsidian, blueprint[1].obsidian, blueprint[2].obsidian, blueprint[3].obsidian)

		recur = func(mins int, stash Stash) int {
			if mins == 0 {
				return stash.geode
			}

			oreRobots := stash.robots[OreRobot]
			clayRobots := stash.robots[ClayRobot]
			obsidianRobots := stash.robots[ObsidianRobot]
			geodeRobots := stash.robots[GeodeRobot]

			// If we have more robots than we could possibly use in a given iteration, for cache purposes
			// use a number robots equal to the max necessary
			if oreRobots >= maxOrePerRound {
				oreRobots = maxOrePerRound
			}

			if clayRobots >= maxClayPerRound {
				clayRobots = maxClayPerRound
			}

			if obsidianRobots >= maxObsidianPerRound {
				obsidianRobots = maxObsidianPerRound
			}

			ore := stash.ore
			totalPossibleOreNeeded := (mins * maxOrePerRound)
			guaranteedOreWillProduce := oreRobots * (mins - 1)
			maxOreToHold := totalPossibleOreNeeded - guaranteedOreWillProduce
			if ore >= maxOreToHold {
				ore = maxOreToHold
			}

			clay := stash.clay
			totalPossibleClayNeeded := (mins * maxClayPerRound)
			guaranteedClayWillProduce := clayRobots * (mins - 1)
			maxClayToHold := totalPossibleClayNeeded - guaranteedClayWillProduce
			if clay > maxClayToHold {
				clay = maxClayToHold
			}

			obsidian := stash.obsidian
			totalPossibleObsidianNeeded := (mins * maxObsidianPerRound)
			guaranteedObsidianWillProduce := obsidianRobots * (mins - 1)
			maxObsidianToHold := totalPossibleObsidianNeeded - guaranteedObsidianWillProduce
			if obsidian > maxObsidianToHold {
				obsidian = maxObsidianToHold
			}

			key := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v", ore, clay, obsidian, stash.geode, oreRobots, clayRobots, obsidianRobots, geodeRobots, mins)
			if v, ok := memo[key]; ok {
				return v
			}

			maxGeode := math.MinInt
			// Build Robots
			for j, cost := range blueprint {
				if canAfford(stash, cost) {
					updStash := stash.RobotsDoWork()
					updStash = updStash.BuildRobot(cost, j)

					g := recur(mins-1, updStash)
					maxGeode = util.Max(maxGeode, g)
				}
			}

			// Don't build Robot
			updStash := stash.RobotsDoWork()
			g := recur(mins-1, updStash)
			maxGeode = util.Max(maxGeode, g)

			memo[key] = maxGeode
			return maxGeode
		}

		geodes := recur(maxMins, NewStash())

		quality := (i + 1) * geodes
		fmt.Printf("[%v] geodes %v - quality: %v\n", i+1, geodes, quality)

		result += quality
	}

	return result
}

func part2(inputFile string) int {
	load(inputFile)
	maxMins := 32

	var recur func(mins int, stash Stash) int

	max := 3
	if len(blueprints) < max {
		max = len(blueprints)
	}

	result := 1
	for i := 0; i < max; i++ {
		blueprint := blueprints[i]
		memo := map[string]int{}

		maxOrePerRound := util.MaxAll(blueprint[0].ore, blueprint[1].ore, blueprint[2].ore, blueprint[3].ore)
		maxClayPerRound := util.MaxAll(blueprint[0].clay, blueprint[1].clay, blueprint[2].clay, blueprint[3].clay)
		maxObsidianPerRound := util.MaxAll(blueprint[0].obsidian, blueprint[1].obsidian, blueprint[2].obsidian, blueprint[3].obsidian)

		recur = func(mins int, stash Stash) int {
			if mins == 0 {
				return stash.geode
			}

			oreRobots := stash.robots[OreRobot]
			clayRobots := stash.robots[ClayRobot]
			obsidianRobots := stash.robots[ObsidianRobot]
			geodeRobots := stash.robots[GeodeRobot]

			// If we have more robots than we could possibly use in a given iteration, for cache purposes
			// use a number robots equal to the max necessary
			if oreRobots >= maxOrePerRound {
				oreRobots = maxOrePerRound
			}

			if clayRobots >= maxClayPerRound {
				clayRobots = maxClayPerRound
			}

			if obsidianRobots >= maxObsidianPerRound {
				obsidianRobots = maxObsidianPerRound
			}

			ore := stash.ore
			totalPossibleOreNeeded := (mins * maxOrePerRound)
			guaranteedOreWillProduce := oreRobots * (mins - 1)
			maxOreToHold := totalPossibleOreNeeded - guaranteedOreWillProduce
			if ore >= maxOreToHold {
				ore = maxOreToHold
			}

			clay := stash.clay
			totalPossibleClayNeeded := (mins * maxClayPerRound)
			guaranteedClayWillProduce := clayRobots * (mins - 1)
			maxClayToHold := totalPossibleClayNeeded - guaranteedClayWillProduce
			if clay > maxClayToHold {
				clay = maxClayToHold
			}

			obsidian := stash.obsidian
			totalPossibleObsidianNeeded := (mins * maxObsidianPerRound)
			guaranteedObsidianWillProduce := obsidianRobots * (mins - 1)
			maxObsidianToHold := totalPossibleObsidianNeeded - guaranteedObsidianWillProduce
			if obsidian > maxObsidianToHold {
				obsidian = maxObsidianToHold
			}

			key := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v", ore, clay, obsidian, stash.geode, oreRobots, clayRobots, obsidianRobots, geodeRobots, mins)
			if v, ok := memo[key]; ok {
				return v
			}

			maxGeode := math.MinInt
			// Build Robots
			for j, cost := range blueprint {
				if canAfford(stash, cost) {
					updStash := stash.RobotsDoWork()
					updStash = updStash.BuildRobot(cost, j)

					g := recur(mins-1, updStash)
					maxGeode = util.Max(maxGeode, g)
				}
			}

			// Don't build Robot
			updStash := stash.RobotsDoWork()
			g := recur(mins-1, updStash)
			maxGeode = util.Max(maxGeode, g)

			memo[key] = maxGeode
			return maxGeode
		}

		geodes := recur(maxMins, NewStash())

		// quality := (i + 1) * geodes
		// fmt.Printf("[%v] geodes %v - quality: %v\n", i+1, geodes, quality)
		fmt.Printf("[%v] geodes %v\n", i+1, geodes)

		result *= geodes
	}

	return result
}
