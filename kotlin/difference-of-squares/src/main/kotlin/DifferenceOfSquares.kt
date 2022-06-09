class Squares(private val n: Int) {
    private fun Int.squared() = this * this

    fun sumOfSquares(): Int {
        return (1..n).sumOf { it.squared() }
    }

    fun squareOfSum(): Int {
        return (1..n).sum().squared()
    }

    fun difference(): Int {
        return squareOfSum() - sumOfSquares()
    }
}
