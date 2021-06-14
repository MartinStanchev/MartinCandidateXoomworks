module MartinCandidate

go 1.16

replace MartinCandidate/translator => ./Plugin/

replace MartinCandidate/controller => ./Controller/

require (
	MartinCandidate/controller v0.0.0-00010101000000-000000000000
	MartinCandidate/translator v0.0.0-00010101000000-000000000000 // indirect
)
