package domain

import "fmt"

type  Discussion struct {
 DiscussionID int64
 Topic Topic
 Discription string
 
}
func (discuss*Discussion) SaveDiscussion(DiscussionCollection*[]*Discussion) (string,bool){

      var descr string
	   results := true
	    //Searching to see if the User sxists
	    for _,value:= range*DiscussionCollection{
			if value.DiscussionID == discuss.DiscussionID{
				descr ="Can not Allow Duplicates"
				results = false
				break
			}
			continue
		}
	 //if A User is registered
	   if results{
		   tempSlice := append(*DiscussionCollection,discuss)
		   *DiscussionCollection =tempSlice
		   descr ="Discussion Saved"
	   }
	return descr,results
}
func (discuss *Discussion) FindDiscussion(DiscussionCollection []*Discussion) (*Discussion, string, bool) {

	var temp *Discussion
	var descr string
	result := false

	for _, value := range DiscussionCollection {
		if value.DiscussionID == discuss.DiscussionID {
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
func (discuss*Discussion) UpdateDiscussion(DiscussionCollection []*Discussion) (string, bool) {

	foundRec := false
	index := -1
	
	//first confirm if the record exists
	for _, value := range DiscussionCollection {
		index++
		if value.DiscussionID==8575 {
			foundRec = true
                        fmt.Printf("===============DISCUSSION INFORMATION===================\n")
                        fmt.Printf("\n")
                        fmt.Printf("DiscusssionID:%d\tTopicID:%d\tDescription:%s\t\n",value.DiscussionID,value.Topic.TopicID,value.Discription)
                        value.Discription ="Please follow my Tutorial on www.prramed@gmail.com"
                        fmt.Printf("===============UPDATED DISCUSSION INFORMATION===================\n")
                        fmt.Printf("\n")
                        fmt.Printf("DiscusssionID:%d\tTopicID:%d\tDescription:%s\t\n",value.DiscussionID,value.Topic.TopicID,value.Discription)
                        
			break
		}
	}
	
	if foundRec {
		DiscussionCollection[index] = discuss
		return "Update Successful", foundRec
	}else{
		return "Record does not exist", foundRec
	}
}
func (discuss*Discussion) DeleteDiscussion(DiscussionCollection []*Discussion) (string, bool) {
	foundRec := false
	index := -1
	var descr string
	
	//first confirm if the record exists
	for _, value := range DiscussionCollection {
		index++
		if value.DiscussionID==discuss.DiscussionID{
			foundRec = true
			break
		}
	}
	
	if foundRec {
		if index >= 0 && index < len(DiscussionCollection) - 1 {
			for ; index < len(DiscussionCollection); index++ {
				if index != len(DiscussionCollection) - 1 {
					DiscussionCollection[index] = DiscussionCollection[index + 1]
				}
			}
			descr = "Discussion Deleted Successful"
		}		
	}else{
		descr = "Record does not exist"
	}

	return descr, foundRec
}




