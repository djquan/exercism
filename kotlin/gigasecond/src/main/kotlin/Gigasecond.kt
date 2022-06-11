import java.time.LocalDate
import java.time.LocalDateTime
import java.time.temporal.ChronoUnit

class Gigasecond(startDateTime: LocalDateTime) {
    private val oneBillion = 1_000_000_000L
    val date: LocalDateTime = startDateTime.plus(oneBillion, ChronoUnit.SECONDS)

    constructor(startDate: LocalDate): this(startDate.atStartOfDay())
}
