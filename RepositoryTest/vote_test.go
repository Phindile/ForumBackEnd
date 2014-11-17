package RepositoryTest
import ("fmt"
	"testing"
	"domain"
)



func buildTestVote() (vote0,vote1,vote2 *domain.Vote){

		


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

	vote0 = &domain.Vote{8575,"yes",comm0}
	vote1 = &domain.Vote{8576,"no",comm1}
	vote2 = &domain.Vote{885,"yes",comm2}

	return
}
func TestSaveVote(t *testing.T){
	var (
		vote0 *domain.Vote
		vote1 *domain.Vote
		vote2*domain.Vote
	)

	vote0,vote1,vote2= buildTestVote()

	voteSlice := []*domain.Vote{vote0,vote1}
       

	fmt.Println(len(voteSlice))

	descr, result := vote2.SaveVote(&voteSlice)

	fmt.Println(len(voteSlice))
	fmt.Println(descr)


	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}
func TestFindVote(t *testing.T){
	var (
		vote0 *domain.Vote
		vote1 *domain.Vote
		vote2*domain.Vote
	)

	vote0,vote1,vote2= buildTestVote()

	voteSlice := []*domain.Vote{vote0,vote1,vote2}

	_,descr,result := vote2.FindVote(voteSlice)

	fmt.Println(descr)


	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}
func TestUpdateVote(t *testing.T){
	var (
		vote0 *domain.Vote
		vote1 *domain.Vote
		vote2*domain.Vote
	)

	vote0,vote1,vote2= buildTestVote()

	voteSlice := []*domain.Vote{vote0,vote1}
        vote2.VoteID =8575

	fmt.Println(len(voteSlice))

	descr, result := vote2.UpdateVote(voteSlice)

	fmt.Println(len(voteSlice))
	fmt.Println(descr)


	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}	

func TestDeleteVote(t *testing.T){		

	var (
		vote0 *domain.Vote
		vote1 *domain.Vote
		vote2 *domain.Vote
	)

	vote0,vote1,vote2 = buildTestVote()		

	voteSlice := []*domain.Vote {vote0,vote1,vote2}	

	descr, result := voteSlice[1].DeleteVote(voteSlice)
	
	fmt.Println(descr)

	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)		
	}	
}







