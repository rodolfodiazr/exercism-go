package twelve

import (
    "fmt"
    "strings"
)

func Verse(i int) string {
    ordinalNumbers := []string{
        "first", "second", "third", "fourth", 
        "fifth", "sixth", "seventh", "eighth", 
        "ninth", "tenth", "eleventh", "twelfth",
    }

    numbers := []string{
        "a", "two", "three", "four",
        "five", "six", "seven", "eight",
        "nine", "ten", "eleven", "twelve",
    }

    gifts := []string{
        "Partridge in a Pear Tree", "Turtle Doves", "French Hens", "Calling Birds",
        "Gold Rings", "Geese-a-Laying", "Swans-a-Swimming", "Maids-a-Milking",
        "Ladies Dancing", "Lords-a-Leaping", "Pipers Piping", "Drummers Drumming",
    }

    out := []string{}
    for j := i; j >= 1; j-- {
        number := numbers[j - 1]
        if i > 1 && j == 1 {
            number = "and a"
        }
        out = append(out, fmt.Sprintf(" %v %v", number, gifts[j - 1]))
    }
    
    return fmt.Sprintf(
        "On the %s day of Christmas my true love gave to me:%s.", 
        ordinalNumbers[i - 1],
        strings.Join(out, ","),
    )
}

func Song() string {
    var out string
	for i := 1; i <= 12; i++ {
        out += Verse(i)
        if i != 12 {
            out +="\n"
        }
    }
    return out
}
