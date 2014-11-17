package RepositoryTest
import ("fmt"
	"testing"
	"domain"
)



func buildTestComment() (comment0,comment1,comment2 *domain.Comment){

		
               


                createdDate0 :=domain.Date{"09","September",2014}
                createdDate1 :=domain.Date{"10","October",2014}
                createdDate2 :=domain.Date{"17","November",2014}
        
             
	topic0 := domain.Topic{8555, "How to use google maps","down this link",createdDate0}
	topic1 := domain.Topic{7647, "How to use  maps on golang","down this link",createdDate1}
	topic2 := domain.Topic{8464, "How to love","go to hell",createdDate2}

	discuss0 := domain.Discussion{8075,topic0,"Revome you second function"}
	discuss1 := domain.Discussion{57475,topic1,"Revome you second function"}
	discuss2 := domain.Discussion{8596,topic2,"Revome you Third function"}

	comment0 = &domain.Comment{8575,discuss0,"check me"}
	comment1 = &domain.Comment{8545,discuss1,"good"}
	comment2 = &domain.Comment{596,discuss2,"you ugly"}

	return
}
func TestSaveComment(t *testing.T){
	var (
		comment0 *domain.Comment
		comment1 *domain.Comment
		comment2*domain.Comment
	)

	comment0,comment1,comment2= buildTestComment()

	commentSlice := []*domain.Comment{comment0,comment1}
       

	fmt.Println(len(commentSlice))

	descr, result := comment2.SaveComment(&commentSlice)

	fmt.Println(len(commentSlice))
	fmt.Println(descr)


	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}
func TestFindComment(t *testing.T){
	var (
		comment0 *domain.Comment
		comment1 *domain.Comment
		comment2*domain.Comment
	)

	comment0,comment1,comment2= buildTestComment()

	commentSlice := []*domain.Comment{comment0,comment1,comment2}

	_,descr,result := comment2.FindComment(commentSlice)

	fmt.Println(descr)


	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}
func TestUpdateComment(t *testing.T){
	var (
		comment0 *domain.Comment
		comment1 *domain.Comment
		comment2*domain.Comment
	)

	comment0,comment1,comment2= buildTestComment()

	commentSlice := []*domain.Comment{comment0,comment1}
        comment2.CommentID =8575

	fmt.Println(len(commentSlice))

	descr, result := comment2.UpdateComment(commentSlice)

	fmt.Println(len(commentSlice))
	fmt.Println(descr)


	if !result {
		t.Error("Expected true,but got ", result, ". \nError Description: ", descr)
	}
}	

func TestDeleteComment(t *testing.T){		

	var (
		comment0 *domain.Comment
		comment1 *domain.Comment
		comment2 *domain.Comment
	)

	comment0,comment1,comment2 = buildTestComment()		

	commentSlice := []*domain.Comment {comment0,comment1,comment2}	

	descr, result := commentSlice[1].DeleteComment(commentSlice)
	
	fmt.Println(descr)

	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)		
	}	
}







