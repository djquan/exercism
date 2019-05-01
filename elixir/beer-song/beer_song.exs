defmodule BeerSong do
  @doc """
  Get a single verse of the beer song
  """
  @spec verse(integer) :: String.t()
  def verse(0) do
    """
    No more bottles of beer on the wall, no more bottles of beer.
    Go to the store and buy some more, 99 bottles of beer on the wall.
    """
  end

  def verse(number) do
    descriptor = if number == 1, do: "it", else: "one"

    """
    #{bottle(number)} of beer on the wall, #{bottle(number)} of beer.
    Take #{descriptor} down and pass it around, #{bottle(number - 1)} of beer on the wall.
    """
  end

  @doc """
  Get the entire beer song for a given range of numbers of bottles.
  """
  @spec lyrics(Range.t()) :: String.t()
  def lyrics(range) do
    range
    |> Enum.map(fn n -> verse(n) end)
    |> Enum.join("\n")
  end

  @doc """
  Get the entire beer song
  """
  @spec lyrics() :: String.t()
  def lyrics, do: lyrics(99..0)

  defp bottle(0), do: "no more bottles"
  defp bottle(1), do: "1 bottle"
  defp bottle(number), do: "#{number} bottles"
end
