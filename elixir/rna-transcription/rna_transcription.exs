defmodule RNATranscription do
  @doc """
  Transcribes a character list representing DNA nucleotides to RNA

  ## Examples

  iex> RNATranscription.to_rna('ACTG')
  'UGAC'
  """
  @spec to_rna([char]) :: [char]
  def to_rna(dna) do
    dna
    |> Enum.map(&nucleotide_compliment/1)
  end

  @spec nucleotide_compliment(char) :: char
  def nucleotide_compliment(?G), do: ?C
  def nucleotide_compliment(?C), do: ?G
  def nucleotide_compliment(?T), do: ?A
  def nucleotide_compliment(?A), do: ?U
end
