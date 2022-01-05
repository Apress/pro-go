package fmt

import "strconv"

func ToCurrency(amount float64) string {
    return "$" + strconv.FormatFloat(amount, 'f', 2, 64)
}
