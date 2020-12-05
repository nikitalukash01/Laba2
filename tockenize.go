package calc
import(
	"regexp"
	"strings"
)


func tockenize(ex string) []Tocken {
	var tockenBuff []Tocken
	exBuff := strings.Split(ex, "")
	for _, value := range exBuff {
		matched, _ := regexp.MatchString(`[0-9]+`, value)
		switch matched {
		case true:
			var tocken Tocken = Tocken{value, Number}
			tockenBuff = append(tockenBuff, tocken)
		case false:
			var tocken Tocken = Tocken{value, Operator}
			tockenBuff = append(tockenBuff, tocken)
			break
		}
	}
	return tockenBuff
}
