package RepositoryTest
import ("fmt"
	"testing"
	"domain"
)



func buildTestDiscuss() (discuss0,discuss1,discuss2 *domain.Discussion){

             
                createdDate0 :=domain.Date{"09","September",2014}
                createdDate1 :=domain.Date{"10","October",2014}
                createdDate2 :=domain.Date{"17","November",2014}
        
             
	topic0 := domain.Topic{8555, "How to use google maps","down this link",createdDate0}
	topic1 := domain.Topic{7647, "How to use  maps on golang","down this link",createdDate1}
	topic2 := domain.Topic{8464, "How to love","go to hell",createdDate2}

	discuss0 = &domain.Discussion{8575,topic0,"Revome you second function"}
	discuss1 = &domain.Discussion{57475,topic1,"Revome you second function"}
	discuss2 = &domain.Discussion{8596,topic2,"Revome you Third function"}

	return
}
func TestSaveDiscussion(t *testing.T){
	var (
		discuss0 *domain.Discussion
		discuss1 *domain.Discussion
		discuss2*domain.Discussion
	)

	discuss0,discuss1,discuss2= buildTestDiscuss()

	discussionSlice := []*domain.Discussion{discuss0,discuss1}
       

	fmt.Println(len(discussionSlice))

	descr, result := discuss2.SaveDiscussion(&discussionSlice)

	fmt.Println(len(discussionSlice))
	fmt.Println(descr)


	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}
func TestFindDiscussion(t *testing.T){
	var (
		discuss0 *domain.Discussion
		discuss1 *domain.Discussion
		discuss2*domain.Discussion
	)

	discuss0,discuss1,discuss2= buildTestDiscuss()

	discussionSlice := []*domain.Discussion{discuss0,discuss1,discuss2}

	_,descr,result := discuss2.FindDiscussion(discussionSlice)

	fmt.Println(descr)


	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}
func TestUpdateDiscussion(t *testing.T){
	var (
		discuss0 *domain.Discussion
		discuss1 *domain.Discussion
		discuss2*domain.Discussion
	)

	discuss0,discuss1,discuss2= buildTestDiscuss()

	discussionSlice := []*domain.Discussion{discuss0,discuss1}
        discuss2.DiscussionID =8575

	fmt.Println(len(discussionSlice))

	descr, result := discuss2.UpdateDiscussion(discussionSlice)

	fmt.Println(len(discussionSlice))
	fmt.Println(descr)


	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}	

func TestDeleteDiscussion(t *testing.T){		

	var (
		discuss0 *domain.Discussion
		discuss1 *domain.Discussion
		discuss2 *domain.Discussion
	)

	discuss0,discuss1,discuss2 = buildTestDiscuss()		

	discussSlice := []*domain.Discussion {discuss0,discuss1,discuss2}	

	descr, result := discussSlice[1].DeleteDiscussion(discussSlice)
	
	fmt.Println(descr)

	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)		
	}	
}







