defmodule Bob do
  def hey(input) do
    cond do
      String.trim(input) == "" ->
        "Fine. Be that way!"

      question?(input) && shouting?(input) ->
        "Calm down, I know what I'm doing!"

      question?(input) ->
        "Sure."

      shouting?(input) ->
        "Whoa, chill out!"

      true ->
        "Whatever."
    end
  end

  defp shouting?(input) do
    String.upcase(input) == input &&
      String.downcase(input) != input
  end

  defp question?(input) do
    String.ends_with?(input, "?")
  end
end
