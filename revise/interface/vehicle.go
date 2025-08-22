package main

import (
	"fmt"
)

// Vehicle Management System
type Vehicle interface {
	Start() string
	Stop() string
	GetInfo() string
}

type FuelVehicle interface {
	Vehicle
	Refuel(amount float64) string
	GetFuelLevel() float64
}

type ElectricVehicle interface {
	Vehicle
	Charge(percentage float64) string
	GetBatteryLevel() float64
}

type Maintainable interface {
	PerformMaintenance() string
	GetMaintenanceStatus() string
}

// Car implementation
type Car struct {
	Brand       string
	Model       string
	FuelLevel   float64
	IsRunning   bool
	LastService string
}

func NewCar(brand, model string) *Car {
	return &Car{
		Brand:       brand,
		Model:       model,
		FuelLevel:   50.0,
		IsRunning:   false,
		LastService: "2024-01-01",
	}
}

func (c *Car) Start() string {
	if c.IsRunning {
		return fmt.Sprintf("%s %s is already running", c.Brand, c.Model)
	}
	c.IsRunning = true
	return fmt.Sprintf("%s %s engine started", c.Brand, c.Model)
}

func (c *Car) Stop() string {
	if !c.IsRunning {
		return fmt.Sprintf("%s %s is already stopped", c.Brand, c.Model)
	}
	c.IsRunning = false
	return fmt.Sprintf("%s %s engine stopped", c.Brand, c.Model)
}

func (c *Car) GetInfo() string {
	status := "stopped"
	if c.IsRunning {
		status = "running"
	}
	return fmt.Sprintf("%s %s - Status: %s, Fuel: %.1f%%", c.Brand, c.Model, status, c.FuelLevel)
}

func (c *Car) Refuel(amount float64) string {
	c.FuelLevel += amount
	if c.FuelLevel > 100 {
		c.FuelLevel = 100
	}
	return fmt.Sprintf("%s %s refueled. Current level: %.1f%%", c.Brand, c.Model, c.FuelLevel)
}

func (c *Car) GetFuelLevel() float64 {
	return c.FuelLevel
}

func (c *Car) PerformMaintenance() string {
	c.LastService = "2024-08-22"
	return fmt.Sprintf("%s %s maintenance completed", c.Brand, c.Model)
}

func (c *Car) GetMaintenanceStatus() string {
	return fmt.Sprintf("%s %s last serviced on: %s", c.Brand, c.Model, c.LastService)
}

// Electric Car implementation
type ElectricCar struct {
	Brand        string
	Model        string
	BatteryLevel float64
	IsRunning    bool
	LastService  string
}

func NewElectricCar(brand, model string) *ElectricCar {
	return &ElectricCar{
		Brand:        brand,
		Model:        model,
		BatteryLevel: 80.0,
		IsRunning:    false,
		LastService:  "2024-01-01",
	}
}

func (e *ElectricCar) Start() string {
	if e.IsRunning {
		return fmt.Sprintf("%s %s is already running", e.Brand, e.Model)
	}
	e.IsRunning = true
	return fmt.Sprintf("%s %s motor started silently", e.Brand, e.Model)
}

func (e *ElectricCar) Stop() string {
	if !e.IsRunning {
		return fmt.Sprintf("%s %s is already stopped", e.Brand, e.Model)
	}
	e.IsRunning = false
	return fmt.Sprintf("%s %s motor stopped", e.Brand, e.Model)
}

func (e *ElectricCar) GetInfo() string {
	status := "stopped"
	if e.IsRunning {
		status = "running"
	}
	return fmt.Sprintf("%s %s - Status: %s, Battery: %.1f%%", e.Brand, e.Model, status, e.BatteryLevel)
}

func (e *ElectricCar) Charge(percentage float64) string {
	e.BatteryLevel += percentage
	if e.BatteryLevel > 100 {
		e.BatteryLevel = 100
	}
	return fmt.Sprintf("%s %s charged. Current level: %.1f%%", e.Brand, e.Model, e.BatteryLevel)
}

func (e *ElectricCar) GetBatteryLevel() float64 {
	return e.BatteryLevel
}

func (e *ElectricCar) PerformMaintenance() string {
	e.LastService = "2024-08-22"
	return fmt.Sprintf("%s %s software update and maintenance completed", e.Brand, e.Model)
}

func (e *ElectricCar) GetMaintenanceStatus() string {
	return fmt.Sprintf("%s %s last serviced on: %s", e.Brand, e.Model, e.LastService)
}

// Motorcycle implementation
type Motorcycle struct {
	Brand     string
	Model     string
	FuelLevel float64
	IsRunning bool
}

func NewMotorcycle(brand, model string) *Motorcycle {
	return &Motorcycle{
		Brand:     brand,
		Model:     model,
		FuelLevel: 70.0,
		IsRunning: false,
	}
}

func (m *Motorcycle) Start() string {
	if m.IsRunning {
		return fmt.Sprintf("%s %s is already running", m.Brand, m.Model)
	}
	m.IsRunning = true
	return fmt.Sprintf("%s %s roars to life", m.Brand, m.Model)
}

func (m *Motorcycle) Stop() string {
	if !m.IsRunning {
		return fmt.Sprintf("%s %s is already stopped", m.Brand, m.Model)
	}
	m.IsRunning = false
	return fmt.Sprintf("%s %s engine stopped", m.Brand, m.Model)
}

func (m *Motorcycle) GetInfo() string {
	status := "stopped"
	if m.IsRunning {
		status = "running"
	}
	return fmt.Sprintf("%s %s - Status: %s, Fuel: %.1f%%", m.Brand, m.Model, status, m.FuelLevel)
}

func (m *Motorcycle) Refuel(amount float64) string {
	m.FuelLevel += amount
	if m.FuelLevel > 100 {
		m.FuelLevel = 100
	}
	return fmt.Sprintf("%s %s refueled. Current level: %.1f%%", m.Brand, m.Model, m.FuelLevel)
}

func (m *Motorcycle) GetFuelLevel() float64 {
	return m.FuelLevel
}

// Vehicle Fleet Manager
type Fleet struct {
	vehicles []Vehicle
}

func NewFleet() *Fleet {
	return &Fleet{vehicles: make([]Vehicle, 0)}
}

func (f *Fleet) AddVehicle(vehicle Vehicle) {
	f.vehicles = append(f.vehicles, vehicle)
	fmt.Printf("Added vehicle: %s\n", vehicle.GetInfo())
}

func (f *Fleet) StartAllVehicles() {
	fmt.Println("\n=== Starting All Vehicles ===")
	for _, vehicle := range f.vehicles {
		fmt.Println(vehicle.Start())
	}
}

func (f *Fleet) StopAllVehicles() {
	fmt.Println("\n=== Stopping All Vehicles ===")
	for _, vehicle := range f.vehicles {
		fmt.Println(vehicle.Stop())
	}
}

func (f *Fleet) RefuelAllFuelVehicles() {
	fmt.Println("\n=== Refueling Fuel Vehicles ===")
	for _, vehicle := range f.vehicles {
		if fuelVehicle, ok := vehicle.(FuelVehicle); ok {
			fmt.Println(fuelVehicle.Refuel(20.0))
		}
	}
}

func (f *Fleet) ChargeAllElectricVehicles() {
	fmt.Println("\n=== Charging Electric Vehicles ===")
	for _, vehicle := range f.vehicles {
		if electricVehicle, ok := vehicle.(ElectricVehicle); ok {
			fmt.Println(electricVehicle.Charge(15.0))
		}
	}
}

func (f *Fleet) PerformMaintenanceAll() {
	fmt.Println("\n=== Performing Maintenance ===")
	for _, vehicle := range f.vehicles {
		if maintainable, ok := vehicle.(Maintainable); ok {
			fmt.Println(maintainable.PerformMaintenance())
		}
	}
}

func (f *Fleet) GetFleetStatus() {
	fmt.Println("\n=== Fleet Status ===")
	for _, vehicle := range f.vehicles {
		fmt.Println(vehicle.GetInfo())
	}
}

func vehicleExample() {
	// Create fleet
	fleet := NewFleet()

	// Create vehicles
	car := NewCar("Toyota", "Camry")
	electricCar := NewElectricCar("Tesla", "Model 3")
	motorcycle := NewMotorcycle("Harley", "Davidson")

	// Add to fleet
	fleet.AddVehicle(car)
	fleet.AddVehicle(electricCar)
	fleet.AddVehicle(motorcycle)

	// Demonstrate operations
	fleet.GetFleetStatus()
	fleet.StartAllVehicles()
	fleet.RefuelAllFuelVehicles()
	fleet.ChargeAllElectricVehicles()
	fleet.PerformMaintenanceAll()
	fleet.GetFleetStatus()
	fleet.StopAllVehicles()
}
