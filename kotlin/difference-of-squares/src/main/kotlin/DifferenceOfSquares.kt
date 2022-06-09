class Squares(private val n: Int) {

    fun sumOfSquares(): Int {
        return (1..n).fold(0) { sum, element -> sum + (element * element)}
    }

    fun squareOfSum(): Int {
        val sum = (1..n).fold(0) { sum, element -> sum + element}

        return sum * sum
    }

    fun difference(): Int {
        return squareOfSum() - sumOfSquares()
    }
}
