package domain

import "fmt"
 
type  User struct {
      UserID int64
      DisplayName string
      Password  Password
      Avatar Avatar
      Demographic Demographic
      Contacts  Contact
      Comment Comment
     

}
func (u*User) SaveUser(UserCollection*[]*User) (string,bool){

     var descr string
	result := true

	//search for not existing record
	for _, value := range *UserCollection {
		if value.UserID == u.UserID {
                     
			descr = "Could not add the item, Duplicates are not allowed!"
			result = false			
			break		
		}
		continue
	}

	if result {
		//Clearly no duplicates were encounted, add the record				
		tempSlice := append(*UserCollection, u)	
		*UserCollection = tempSlice
		descr = "Successfuly Added!"	     
	}
	
	return descr, result
}
func (u *User) Find(userCollection []*User) (*User, string, bool) {

	var temp *User
	var descr string
	result := false

	for _, value := range userCollection {
		if value.UserID ==7873 {
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
func (u *User) Update(userCollection []*User) (string, bool) {

	foundRec := false
	index := -1
	
	//first confirm if the record exists
       
	for _, value := range userCollection {
		index++
		if value.UserID ==7873 {
			foundRec = true
                         value.DisplayName = "Chrisney"
                      fmt.Printf("====USER UPDATED INFO===\n") 
 fmt.Printf("UserID:%d\tDisplayName:%s\tPassword:%s\tPicture:%s\tContacts:%s\n",value.UserID,value.DisplayName,value.Password.Password,value.Avatar.AvatarLocation,value.Contacts.HomePhone)
                       break
                       

		}
	}
	
	if foundRec {
		userCollection[index] = u
		return "Update Successful", foundRec
	}else{
		return "Record does not exist", foundRec
	}
}


   


func (u *User) Delete(userCollection []*User) (string, bool) {
	foundRec := false
	index := -1
	var descr string
	
	//first confirm if the record exists
	for _, value := range userCollection {
		index++
		if value.UserID == u.UserID{
			foundRec = true
			break
		}
	}
	
	if foundRec {
		if index >= 0 && index < len(userCollection) - 1 {
			for ; index < len(userCollection); index++ {
				if index != len(userCollection) - 1 {
					userCollection[index] = userCollection[index + 1]
				}
			}
			descr = "Deletion Successful"
		}		
	}else{
		descr = "Record does not exist"
	}

	return descr, foundRec
}





