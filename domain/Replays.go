package domain

type Reply struct {
	ReplyID int64
	Discussion Discussion
	Comment Comment
	Reply string // either its for a comment or discussion (COMM or DISC)
}

func (reply*Reply) SaveReply(ReplyCollection*[]*Reply) (string,bool){

      var descr string
	   results := true
	    //Searching to see if the User sxists
	    for _,value:= range*ReplyCollection{
			if value.ReplyID == reply.ReplyID{
				descr ="Can not Allow Duplicates"
				results = false
				break
			}
			continue
		}
	 //if A User is registered
	   if results{
		   tempSlice := append(*ReplyCollection,reply)
		   *ReplyCollection =tempSlice
		   descr ="Reply Saved"
                   
	   }
	return descr,results
}
func (reply *Reply) FindReply(ReplyCollection []*Reply) (*Reply, string, bool) {

	var temp *Reply
	var descr string
	result := false

	for _, value := range ReplyCollection {
		if value.ReplyID == reply.ReplyID {
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
func (reply*Reply) UpdateReply(ReplyCollection []*Reply) (string, bool) {

	foundRec := false
	index := -1
	
	//first confirm if the record exists
	for _, value := range ReplyCollection {
		index++
		if value.ReplyID==8677{
			foundRec = true
                        value.Reply ="we dont know you"
			break
                    
		}
	}
	
	if foundRec {
		ReplyCollection[index] = reply
		return "Update Successful", foundRec
	}else{
		return "Record does not exist", foundRec
	}
}
func (reply*Reply) DeleteReply(ReplyCollection []*Reply) (string, bool) {
	foundRec := false
	index := -1
	var descr string
	
	//first confirm if the record exists
	for _, value := range ReplyCollection {
		index++
		if value.ReplyID==reply.ReplyID{
			foundRec = true
			break
		}
	}
	
	if foundRec {
		if index >= 0 && index < len(ReplyCollection) - 1 {
			for ; index < len(ReplyCollection); index++ {
				if index != len(ReplyCollection) - 1 {
					ReplyCollection[index] = ReplyCollection[index + 1]
				}
			}
			descr = "Repl Deleted Successful"
		}		
	}else{
		descr = "Record does not exist"
	}

	return descr, foundRec
}


