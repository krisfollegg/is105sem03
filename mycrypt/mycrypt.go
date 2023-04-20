package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; ")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
        cipherAlphabet := append([]rune("KSN"), alphabet...)
        kryptertMelding := make([]rune, len(melding))
        for i := 0; i < len(melding); i++ {
                indeks := sokIAlfabetet(melding[i], cipherAlphabet)
                if indeks+chiffer >= len(cipherAlphabet) {
                        kryptertMelding[i] = cipherAlphabet[indeks+chiffer-len(cipherAlphabet)]
                } else {
                        kryptertMelding[i] = cipherAlphabet[indeks+chiffer]
                }
        }
        return kryptertMelding
}

}

func ShiftAlphabet(alphabet []rune, shiftAmount int, shiftDirection bool) []rune {
    shiftedAlphabet := make([]rune, len(alphabet))
    for i := 0; i < len(alphabet); i++ {
        shiftIndex := i
        if shiftDirection {
            shiftIndex = (i + shiftAmount) % len(alphabet)
        } else {
            shiftIndex = (i - shiftAmount + len(alphabet)) % len(alphabet)
        }
        shiftedAlphabet[i] = alphabet[shiftIndex]
    }
    return shiftedAlphabet
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
    for i := 0; i < len(alfabet); i++ {
        if symbol == alfabet[i] {
            return i
        }
    }
    return -1
}

