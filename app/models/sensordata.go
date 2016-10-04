package models

import (
    // "github.com/revel/revel"
)
type SensorData struct {
    Name                string  `json:"name"`
    Location            string  `json:"location"`
    Serial_Number       int     `json:"serial_number"`
    DateTime            string  `json:"datetime"`
    Temperate           float64 `json:"temperature"`
    Pressure            float64 `json:"pressure"`
    Humidity            float64 `json:"humidity"`
    Water_Level         float64 `json:"water_level"`
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
