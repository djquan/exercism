defmodule Words do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t()) :: map
  def count(sentence) do
    sentence
    |> String.downcase
    |> String.replace(~r/[^\w1-9\s\_\-]/u, "")
    |> String.split([" ", "_"], trim: true)
    |> Enum.reduce(%{}, fn(word, acc) -> Map.update(acc, word, 1, &(&1 + 1)) end)
  end
end
