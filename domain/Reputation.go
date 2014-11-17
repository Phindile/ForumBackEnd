package domain

 
type  Reputation struct {
 ReputationID int64
 Owner User
 Rate int64

 
}
func (rep*Reputation ) SaveReputation (ReputationCollection*[]*Reputation ) (string,bool){

      var descr string
	   results := true
	    //Searching to see if the User sxists
	    for _,value:= range*ReputationCollection{
			if value.ReputationID == rep.ReputationID{
				descr ="Can not Allow Duplicates"
				results = false
				break
			}
			continue
		}
	 //if A User is registered
	   if results{
		   tempSlice := append(*ReputationCollection,rep)
		   *ReputationCollection =tempSlice
		   descr ="Thank you For Voting Saved"
	   }
	return descr,results
}
func (rep*Reputation) FindReputation(ReputationCollection []*Reputation) (*Reputation, string, bool) {

	var temp *Reputation
	var descr string
	result := false

	for _, value := range ReputationCollection {
		if value.ReputationID == rep.ReputationID {
			temp = value
			descr = "Found"
			result = true
			break
		}else{	
			continue
		}
	}
	
	if !result {
		descr = "Not Found"
	}		

	return temp, descr, result
}
func (rep*Reputation ) UpdateReputation (ReputationCollection []*Reputation) (string, bool) {

	foundRec := false
	index := -1
	
	//first confirm if the record exists
	for _, value := range ReputationCollection {
		index++
		if value.ReputationID==76658 {
			foundRec = true
                        value.Rate =100
			break
		}
	}
	
	if foundRec {
		ReputationCollection[index] = rep
		return "Update Successful", foundRec
	}else{
		return "Record does not exist", foundRec
	}
}
func (rep*Reputation) DeleteReputation(ReputationCollection []*Reputation) (string, bool) {
	foundRec := false
	index := -1
	var descr string
	
	//first confirm if the record exists
	for _, value := range ReputationCollection {
		index++
		if value.ReputationID==rep.ReputationID{
			foundRec = true
			break
		}
	}
	
	if foundRec {
		if index >= 0 && index < len(ReputationCollection) - 1 {
			for ; index < len(ReputationCollection); index++ {
				if index != len(ReputationCollection) - 1 {
					ReputationCollection[index] = ReputationCollection[index + 1]
				}
			}
			descr = "Reputation Deleted Successful"
		}		
	}else{
		descr = "Record does not exist"
	}

	return descr, foundRec
}




