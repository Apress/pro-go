package params

import (
    "reflect"
    "encoding/json"
    "io"
    "strings"
)

func populateStructFromForm(structVal reflect.Value, 
        formVals map[string][]string) (err error) {
    for i := 0; i < structVal.Elem().Type().NumField(); i++ {
        field := structVal.Elem().Type().Field(i)
        for key, vals := range formVals {
            if strings.EqualFold(key, field.Name) && len(vals) > 0 {
                valField := structVal.Elem().Field(i)
                if (valField.CanSet()) {
                    valToSet, convErr := parseValueToType(valField.Type(), vals[0])
                    if (convErr == nil) {
                        valField.Set(valToSet)
                    } else {
                        err = convErr
                    }
                }
            }
        }
    }
    return
}

func populateStructFromJSON(structVal reflect.Value, 
        reader io.ReadCloser) (err error) {
    return json.NewDecoder(reader).Decode(structVal.Interface())
}
