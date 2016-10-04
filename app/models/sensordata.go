package models

import (
    // "github.com/revel/revel"
)

type SensorData struct {
    name                string  `json:"name"`
    location            string  `json:"location"`
    serial_number       int     `json:"serial_number"`
    datetime            string  `json:"datetime"`
    temperate           float64 `json:"temperature"`
    pressure            float64 `json:"pressure"`
    humidity            float64 `json:"humidity"`
    water_level         float64 `json:"water_level"`
}

// func (sensor_data *SensorData) Validate(v *revel.Validation) {
//     v.Check(sensor_data.Name,
//             revel.Required{},
//             revel.MaxSize{100},
//     )

//     v.MaxSize(sensor_data.location, 200)

        
//     v.Check(hotel.Country,
//             revel.Required{},
//             revel.MaxSize{40},
//             revel.MinSize{2},
//     )
// }
