package main
import "fmt"

// function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Printf("You got %v gallons of fuel left.", fuel)
}

// function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int
  switch planet {
    case "Venus":
      fuel = 300000
    case "Mercury":
      fuel = 500000
    case "Mars":
      fuel = 700000
    default:
      fuel = 0
  }
  return fuel
}

// function greetPlanet() here
func greetPlanet(planet string) {
  fmt.Println("Welcome to the planet", planet)
}

// function cantFly() here
func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

// function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining, fuelCost int
  fuelRemaining = fuel
  fuelCost = calculateFuel("Mercury")
  if fuelRemaining >= fuelCost {
    greetPlanet("Mercury")
    fuelRemaining -= fuelCost
  } else {
    cantFly()
  }
  return fuelRemaining
}

func main() {
  // `planetChoice` and `fuel` variables
  var fuel int
  fuel = 1000000
  planetChoice := "Venus"
  fuel = flyToPlanet(planetChoice, fuel)
  // And then liftoff!
  fuelGauge(fuel)
}