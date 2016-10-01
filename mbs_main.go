
type loan struct {
	loanAccNo int
	loanAmt int
	loanRate int
	loanTerm int
	propertyNo string
	borrowerName string
	borrowerBSN string
	borrowerCreditRating int
}

type tranche struct {
	trancheId int
	trancheRating int
	trancheRate int	
	loans []loan //array of loans

}


type trancheRatingList struct {
	list [3]string //harcoded list of values
}




	