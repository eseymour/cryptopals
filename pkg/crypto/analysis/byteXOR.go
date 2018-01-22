package analysis

import (
	"math"

	"gonum.org/v1/gonum/stat"

	"github.com/eseymour/cryptopals/pkg/crypto/xor"
)

// English language letter distribution generated by using 1-gram data collected
// by Peter Norvig at http://norvig.com/mayzner.html
var englishDistribution = [26]float64{
	0.08040605150731232, 0.014846487698236691, 0.03343773688810862,
	0.038169582783224694, 0.12492062534197067, 0.02403123383775965,
	0.018693758446142437, 0.05053301406632262, 0.07569277540220806,
	0.0015877372404490017, 0.00540513489633885, 0.0406898604777074,
	0.02511760599410516, 0.07233629221774214, 0.076406929387264,
	0.02135891018410595, 0.0012046892068254826, 0.06279420706085999,
	0.06512766562623011, 0.0927556484289489, 0.02729701843405106,
	0.01053251617932296, 0.01675664190911723, 0.002348568874531159,
	0.016649800974448417, 0.0008995069366664322,
}

// BreakXOREncryptByteKey finds byte key that is most likely to be the actual
// key to decode the ciphertext using character frequency analysis from English
// text. Lower chiSquare indicate greater confidence that the key is correct
func BreakXOREncryptByteKey(ciphertext []byte) (key byte, chiSquare float64) {
	score := math.Inf(0)

	plaintext := make([]byte, len(ciphertext))
	for candidate := 0; candidate <= math.MaxUint8; candidate++ {
		xor.EncryptByteKey(plaintext, ciphertext, byte(candidate))

		var candidateFrequencies [26]float64
		count := 0
		for _, c := range plaintext {
			switch {
			case 'a' <= c && c <= 'z':
				c -= 'a'
			case 'A' <= c && c <= 'Z':
				c -= 'A'
			default:
				continue
			}
			candidateFrequencies[c]++
			count++
		}

		var expectedFrequencies [26]float64
		for i, p := range englishDistribution {
			expectedFrequencies[i] = p * float64(count)
		}

		// Scores are ChiSquare and multiplied by the ratio of plaintext length and
		// number of alphabet characters found
		candidateChiSquare := stat.ChiSquare(candidateFrequencies[:], expectedFrequencies[:])
		candidateScore := candidateChiSquare * float64(len(plaintext)) / float64(count)

		if candidateScore < score {
			score = candidateScore
			chiSquare = candidateChiSquare
			key = byte(candidate)
		}
	}

	return key, chiSquare
}
