package domain

import "fmt"


type  Comment struct {
 CommentID int64
 Discussion Discussion
 Discription string
 
 
 
}
func (comment*Comment) SaveComment(CommentCollection*[]*Comment) (string,bool){

      var descr string
	   results := true
	    //Searching to see if the User sxists
	    for _,value:= range*CommentCollection{
			if value.CommentID == comment.CommentID{
				descr ="Can not Allow Duplicates"
				results = false
				break
			}
			continue
		}
	 //if A User is registered
	   if results{
		   tempSlice := append(*CommentCollection,comment)
		   *CommentCollection =tempSlice
		   descr ="Comment Saved"
                   
	   }
	return descr,results
}
func (comment *Comment) FindComment(CommentCollection []*Comment) (*Comment, string, bool) {

	var temp *Comment
	var descr string
	result := false

	for _, value := range CommentCollection {
		if value.CommentID == comment.CommentID {
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
func (comment*Comment) UpdateComment(CommentCollection []*Comment) (string, bool) {

	foundRec := false
	index := -1
	
	//first confirm if the record exists
	for _, value := range CommentCollection {
		index++
		if value.CommentID==8575{
			foundRec = true
                        fmt.Printf("====================COMMENT INFORMATION==========\n")
                  fmt.Printf("CommentiID:%d\tDiscussionID:%d\tTopicID:%d\tComment:%s\n",value.CommentID,value.Discussion.DiscussionID,value.Discussion.Topic.TopicID,value.Discription)
                 
                        value.Discription ="Lol That is lame,you are so brain dead"
                        fmt.Printf("====================UPDTATED COMMENT INFORMATION==========\n")
                  fmt.Printf("CommentiID:%d\tDiscussionID:%d\tTopicID:%d\tComment:%s\n",value.CommentID,value.Discussion.DiscussionID,value.Discussion.Topic.TopicID,value.Discription)
			break
                    
                      
		}
	}
	
	if foundRec {
		CommentCollection[index] = comment
		return "Update Successful", foundRec
	}else{
		return "Record does not exist", foundRec
	}
}
func (comment*Comment) DeleteComment(CommentCollection []*Comment) (string, bool) {
	foundRec := false
	index := -1
	var descr string
	
	//first confirm if the record exists
	for _, value := range CommentCollection {
		index++
		if value.CommentID==comment.CommentID{
			foundRec = true
			break
		}
	}
	
	if foundRec {
		if index >= 0 && index < len(CommentCollection) - 1 {
			for ; index < len(CommentCollection); index++ {
				if index != len(CommentCollection) - 1 {
					CommentCollection[index] = CommentCollection[index + 1]
				}
			}
			descr = "Comment Deleted Successful"
		}		
	}else{
		descr = "Record does not exist"
	}

	return descr, foundRec
}




