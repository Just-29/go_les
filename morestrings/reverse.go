// Пакет morestrings реализует дополнительные функции для работы со строками в кодировке UTF-8,
// помимо тех, что предоставляются в стандартном пакете "strings".
package morestrings

// Функция ReverseRunes возвращает строку-аргумент, перевернутую рунами слева направо.
func ReverseRunes(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}