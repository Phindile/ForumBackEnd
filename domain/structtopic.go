package domain

import "fmt"

type  Topic struct {
	TopicID int64
	Title string
	Description string
    DatePosted Date
	
}



func (top*Topic) SaveTopic(TopicCollection*[]*Topic) (string,bool){

	var descr string
	results := true
	//Searching to see if the User sxists
	for _,value:= range*TopicCollection{
                   
                  if value.TopicID ==top.TopicID{
			descr ="Can not allow Duplicates"
			results = false
			break 
                                    
		}
		continue
	}
	//if A User is registered
	if results{
		tempSlice := append(*TopicCollection,top)
		*TopicCollection =tempSlice
		descr ="Thank you for Positing you Topic"
	}
	return descr,results
}
func (top*Topic) Find(TopicCollection []*Topic) (*Topic, string, bool) {

	var temp *Topic
	var descr string
	result := false

	for _, value := range TopicCollection {
		if value.TopicID == top.TopicID {
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
func (top *Topic) Update(TopicCollection []*Topic) (string, bool) {

	foundRec := false
	index := -1
	
	//first confirm if the record exists
	for _, value := range TopicCollection {
		index++
		if value.TopicID==7647{
                        foundRec = true
                        value.Title ="How to use For loops in golang"
                         fmt.Printf("Title:%s\tDescription:%s\tDatePosted:\t%s-%s-%d\n",value.Title,value.Description,value.DatePosted.Day,value.DatePosted.Month,value.DatePosted.Year)
                         break
			
		}
	}
	
	if foundRec {


		TopicCollection[index] = top
		return "Update Successful", foundRec
	}else{
		return "Record does not exist", foundRec
	}
}
func (top *Topic) Delete(TopicCollection []*Topic) (string, bool) {
	foundRec := false
	index := -1
	var descr string
	
	//first confirm if the record exists
	for _, value := range TopicCollection {
		index++
		if value.TopicID == top.TopicID{
			foundRec = true
			break
		}
	}
	
	if foundRec {
		if index >= 0 && index < len(TopicCollection) - 1 {
			for ; index < len(TopicCollection); index++ {
				if index != len(TopicCollection) - 1 {
					TopicCollection[index] = TopicCollection[index + 1]
				}
			}
			descr = "Deletion Successful"
		}		
	}else{
		descr = "Record does not exist"
	}

	return descr, foundRec
}


