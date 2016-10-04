package models

import (
    "errors"
    "reflect"
    "fmt"
    // "github.com/revel/revel"
)
func SetField(obj interface{}, name string, value interface{}) error {
    structValue := reflect.ValueOf(obj).Elem()
    structFieldValue := structValue.FieldByName(name)

    if !structFieldValue.IsValid() {
        return fmt.Errorf("No such field: %s in obj", name)
    }

    if !structFieldValue.CanSet() {
        return fmt.Errorf("Cannot set %s field value", name)
    }

    structFieldType := structFieldValue.Type()
    val := reflect.ValueOf(value)
    if structFieldType != val.Type() {
        return errors.New("Provided value type didn't match obj field type")
    }

    structFieldValue.Set(val)
    return nil
}

func (s SensorData) FillStruct(m map[string]interface{}) error {
    for k, v := range m {
        err := SetField(s, k, v)
        if err != nil {
            return err
        }
    }
    return nil
}

type SensorData struct {
    name                string
    location            string
    serial_number       int
    datetime            string
    temperate           float64
    pressure            float64
    humidity            float64
    water_level         float64
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
