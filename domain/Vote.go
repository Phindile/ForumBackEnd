package domain

 
type  Vote struct {
 VoteID int64
 VoteState string
 Comment Comment

 
}
func (vote*Vote) SaveVote(VoteCollection*[]*Vote) (string,bool){

      var descr string
	   results := true
	    //Searching to see if the User sxists
	    for _,value:= range*VoteCollection{
			if value.VoteID == vote.VoteID{
				descr ="Can not Allow Duplicates"
				results = false
				break
			}
			continue
		}
	 //if A User is registered
	   if results{
		   tempSlice := append(*VoteCollection,vote)
		   *VoteCollection =tempSlice
		   descr ="Thank you For Voting Saved"
	   }
	return descr,results
}
func (vote*Vote) FindVote(VoteCollection []*Vote) (*Vote, string, bool) {

	var temp *Vote
	var descr string
	result := false

	for _, value := range VoteCollection {
		if value.VoteID == vote.VoteID {
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
func (vote*Vote) UpdateVote(VoteCollection []*Vote) (string, bool) {

	foundRec := false
	index := -1
	
	//first confirm if the record exists
	for _, value := range VoteCollection {
		index++
		if value.VoteID==vote.VoteID {
			foundRec = true
			break
		}
	}
	
	if foundRec {
		VoteCollection[index] = vote
		return "Update Successful", foundRec
	}else{
		return "Record does not exist", foundRec
	}
}
func (vote*Vote) DeleteVote(VoteCollection []*Vote) (string, bool) {
	foundRec := false
	index := -1
	var descr string
	
	//first confirm if the record exists
	for _, value := range VoteCollection {
		index++
		if value.VoteID==vote.VoteID{
			foundRec = true
			break
		}
	}
	
	if foundRec {
		if index >= 0 && index < len(VoteCollection) - 1 {
			for ; index < len(VoteCollection); index++ {
				if index != len(VoteCollection) - 1 {
					VoteCollection[index] = VoteCollection[index + 1]
				}
			}
			descr = "Vote Deleted Successful"
		}		
	}else{
		descr = "Record does not exist"
	}

	return descr, foundRec
}




