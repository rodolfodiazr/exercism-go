package twelve

import (
    "fmt"
    "strings"
)

var ordinalNumbers = []string {
    "first", "second", "third", "fourth", 
    "fifth", "sixth", "seventh", "eighth", 
    "ninth", "tenth", "eleventh", "twelfth",
}

var numbers = []string {
    "a", "two", "three", "four",
    "five", "six", "seven", "eight",
    "nine", "ten", "eleven", "twelve",
}

var gifts = []string {
    "Partridge in a Pear Tree", "Turtle Doves", "French Hens", "Calling Birds",
    "Gold Rings", "Geese-a-Laying", "Swans-a-Swimming", "Maids-a-Milking",
    "Ladies Dancing", "Lords-a-Leaping", "Pipers Piping", "Drummers Drumming",
}

func Verse(i int) string {
    list := []string{}
    for j := i; j >= 1; j-- {
        number := numbers[j-1]
        if i > 1 && j == 1 {
            number = "and a"
        }
        list = append(list, fmt.Sprintf(" %s %s", number, gifts[j-1]))
    }
    
    return fmt.Sprintf(
        "On the %s day of Christmas my true love gave to me:%s.", 
        ordinalNumbers[i-1],
        strings.Join(list, ","),
    )
}

func Song() string {
    var verses []string
	for i := 1; i <= 12; i++ {
        verses = append(verses, Verse(i))
    }
    return strings.Join(verses, "\n")
}
