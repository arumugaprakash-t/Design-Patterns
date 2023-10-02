package main

import "fmt"

type Observer interface {
	update(float64, float64, float64)
	getID() string
}

type Subject interface {
	registerObserver(observer Observer)
	deregisterObserver(observer Observer)
	notifyAll()
}

//WetherData struct is the subject here, it needs to implement all the subject inferface methods
type WetherData struct {
	Temperature float64
	Humidity    float64
	Pressure    float64
	Observers   []Observer
}

func (data *WetherData) registerObserver(observer Observer) {
	data.Observers = append(data.Observers, observer)
}

func (data *WetherData) deregisterObserver(observer Observer) {
	data.Observers = removeFromSlice(data.Observers, observer)
}

func (data *WetherData) notifyAll() {
	for _, observer := range data.Observers {
		observer.update(data.Temperature, data.Humidity, data.Pressure)
	}
}

func newWeatherData(temparature float64, humiditiy float64, pressure float64) *WetherData {
	return &WetherData{
		Temperature: temparature,
		Humidity:    humiditiy,
		Pressure:    pressure,
	}
}

func removeFromSlice(data []Observer, observer Observer) []Observer {

	for i, val := range data {
		if val.getID() == observer.getID() {
			data = append(data[:i], data[i+1:]...)
			break
		}
	}
	return data
}

func (w *WetherData) updateWeatherData(temparature float64, humiditiy float64, pressure float64) {
	w.Temperature = temparature
	w.Humidity = humiditiy
	w.Pressure = pressure
	w.notifyAll()
}

// To display the weather data
type Display interface {
	display()
}

// Digital display is our Observer and it implements the Display interface as well
type DigitalDisplay struct {
	ID          string
	Temperature float64
	Pressure    float64
	WetherData  *WetherData
}

func (d *DigitalDisplay) newDigitalDisplay(id string, weatherData *WetherData) {
	d.ID = id
	d.WetherData = weatherData
	d.WetherData.registerObserver(d)
}

func (d *DigitalDisplay) display() {
	fmt.Printf("Weather Data Updated :===: Temperature: %f\nPressure: %f\n", d.Temperature, d.Pressure)
}

func (d *DigitalDisplay) update(temparature float64, _ float64, pressure float64) {
	d.Temperature = temparature
	d.Pressure = pressure
	d.display()
}

func (d *DigitalDisplay) getID() string {
	return d.ID
}

func main() {
	weatherData := newWeatherData(99.9, 100, 98)

	digitalDisplay := DigitalDisplay{}
	digitalDisplay.newDigitalDisplay("display-1", weatherData)

	weatherData.updateWeatherData(100, 101, 102)
	weatherData.updateWeatherData(103, 104, 103)
}
