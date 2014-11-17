package RepositoryTest
import ("fmt"
	"testing"
	"domain"
)



func buildTestReply() (reply0,reply1,reply2 *domain.Reply){

		
               


                createdDate0 :=domain.Date{"09","September",2014}
                createdDate1 :=domain.Date{"10","October",2014}
                createdDate2 :=domain.Date{"17","November",2014}
        
             
	topic0 := domain.Topic{8555, "How to use google maps","down this link",createdDate0}
	topic1 := domain.Topic{7647, "How to use  maps on golang","down this link",createdDate1}
	topic2 := domain.Topic{8464, "How to love","go to hell",createdDate2}

	discuss0 := domain.Discussion{8075,topic0,"Revome you second function"}
	discuss1 := domain.Discussion{57475,topic1,"Revome you second function"}
	discuss2 := domain.Discussion{8596,topic2,"Revome you Third function"}

	comm0 := domain.Comment{8575,discuss0,"check me"}
	comm1 := domain.Comment{8545,discuss1,"good"}
	comm2 := domain.Comment{596,discuss2,"you ugly"}

        reply0 = &domain.Reply{8677,discuss0,comm0,"check you on what"}
	reply1 = &domain.Reply{988,discuss1,comm1,"no its too bad"}
	reply2 = &domain.Reply{5464,discuss2,comm2,"as if you pretty"}


	return
}
func TestSaveReply(t *testing.T){
	var (
		reply0 *domain.Reply
		reply1 *domain.Reply
		reply2*domain.Reply
	)

	reply0,reply1,reply2= buildTestReply()

	replySlice := []*domain.Reply{reply0,reply1}
       

	fmt.Println(len(replySlice))

	descr, result := reply2.SaveReply(&replySlice)

	fmt.Println(len(replySlice))
	fmt.Println(descr)


	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}
func TestFindReply(t *testing.T){
	var (
		reply0 *domain.Reply
		reply1 *domain.Reply
		reply2*domain.Reply
	)

	reply0,reply1,reply2= buildTestReply()

	replySlice := []*domain.Reply{reply0,reply1,reply2}

	_,descr,result := reply2.FindReply(replySlice)

	fmt.Println(descr)


	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}
func TestUpdateReply(t *testing.T){
	var (
		reply0 *domain.Reply
		reply1 *domain.Reply
		reply2*domain.Reply
	)

	reply0,reply1,reply2= buildTestReply()

	replySlice := []*domain.Reply{reply0,reply1}
        reply2.ReplyID =8575

	fmt.Println(len(replySlice))

	descr, result := reply2.UpdateReply(replySlice)

	fmt.Println(len(replySlice))
	fmt.Println(descr)


	if !result {
		t.Error("Expected true,but got ", result, ". \nError Description: ", descr)
	}
}	

func TestDeleteReply(t *testing.T){		

	var (
		reply0 *domain.Reply
		reply1 *domain.Reply
		reply2 *domain.Reply
	)

	reply0,reply1,reply2 = buildTestReply()		

	replySlice := []*domain.Reply {reply0,reply1,reply2}	

	descr, result := replySlice[1].DeleteReply(replySlice)
	
	fmt.Println(descr)

	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)		
	}	
}







