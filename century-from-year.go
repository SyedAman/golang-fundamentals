// determines the century of a given year
// 2000 is the 20th century, 1999 is the 20th century, and 2001 
func getCenturyFromYear(year int) int {
 if (year % 100 > 0) {
   return year / 100 + 1
 } else {
   return year / 100
 }
}
