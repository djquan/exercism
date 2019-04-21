defmodule Roman do
  @doc """
  Convert the number to a roman number.
  """
  @spec numerals(pos_integer) :: String.t()
  def numerals(number), do: numeral_helper(number, "")

  [
    M: 1000,
    CM: 900,
    D: 500,
    CD: 400,
    C: 100,
    XC: 90,
    L: 50,
    XL: 40,
    X: 10,
    IX: 9,
    V: 5,
    IV: 4,
    I: 1
  ]
  |> Enum.each(fn {numeral, number} ->
    def numeral_helper(n, acc) when n >= unquote(number) do
      numeral_helper(n - unquote(number), acc <> unquote(Atom.to_string(numeral)))
    end
  end)

  def numeral_helper(0, acc), do: acc
end
