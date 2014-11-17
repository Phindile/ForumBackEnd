package RepositoryTest
import ("fmt"
	"testing"
	"domain"
)



func buildTestReputation() (rep0,rep1,rep2 *domain.Reputation){

		 createdDate0 :=domain.Date{"09","September",2014}
                createdDate1 :=domain.Date{"10","October",2014}
                createdDate2 :=domain.Date{"17","November",2014}
        
             
	topic0 := domain.Topic{8555, "How to use google maps","down this link",createdDate0}
	topic1 := domain.Topic{7647, "How to use  maps on golang","down this link",createdDate1}
	topic2 := domain.Topic{8464, "How to love","go to hell",createdDate2}

	discuss0 := domain.Discussion{8075,topic0,"Revome you second function"}
	discuss1 := domain.Discussion{57475,topic1,"Revome you second function"}
	discuss2 := domain.Discussion{8596,topic2,"Revome you Third function"}

	comment0 := domain.Comment{8575,discuss0,"check me"}
	comment1 := domain.Comment{8545,discuss1,"good"}
	comment2 := domain.Comment{596,discuss2,"you ugly"}


                user0Avatar := domain.Avatar{"Pics/pizz"}
        user0Demo := domain.Demographic{"Male"}
        user0Email:=domain.Email{"prramed@gmail.com"}
        user0Contact :=domain.Contact{user0Email, "021 555 5545", "076 646 1457", "021 666 8854"}
        user0Password := domain.Password{"chicco_3473"}

        user1Avatar := domain.Avatar{"Pics/izz"}
	user1Demo := domain.Demographic{"Male"}
	user1Email := domain.Email{"lwandilem21@gmail.com"}
	user1Contact := domain.Contact{user1Email, "021 444 5545", "076 111 1457", "021 623 8434"}
	user1Password := domain.Password{"123ewq2"}

	user2Avatar := domain.Avatar{"Pics/pizz"}
	user2Demo := domain.Demographic{"Female"}
	user2Email := domain.Email{"lelethu@gmail.com"}
	user2Contact := domain.Contact{user2Email, "031 565 5456", "045 456 4564", "021 621 8222"}
	user2Password := domain.Password{"cooper@gmail.com"}

        user0 := domain.User{7878, "Pizzy's", user0Password, user0Avatar, user0Demo, user0Contact,comment0}
	user1 := domain.User{7873, "EasyPay", user1Password, user1Avatar, user1Demo, user1Contact,comment1}
	user2 := domain.User{7233, "Epping", user2Password, user2Avatar, user2Demo, user2Contact,comment2}

        rep0 =&domain.Reputation{76658,user0,45}
        rep1 =&domain.Reputation{8885,user1,98}
        rep2 =&domain.Reputation{5554,user2,88}
        

	return
}
func TestSaveReputation(t *testing.T){
	var (
		rep0 *domain.Reputation
		rep1 *domain.Reputation
		rep2*domain.Reputation
	)

	rep0,rep1,rep2= buildTestReputation()

	reputationSlice := []*domain.Reputation{rep0,rep1}
       

	fmt.Println(len(reputationSlice))

	descr, result := rep2.SaveReputation(&reputationSlice)

	fmt.Println(len(reputationSlice))
	fmt.Println(descr)


	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}
func TestFindReputation(t *testing.T){
	var (
		rep0 *domain.Reputation
		rep1 *domain.Reputation
		rep2*domain.Reputation
	)

	rep0,rep1,rep2= buildTestReputation()

	reputationSlice := []*domain.Reputation{rep0,rep1,rep2}

	_,descr,result := rep2.FindReputation(reputationSlice)

	fmt.Println(descr)


	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}
func TestUpdateReputation(t *testing.T){
	var (
		rep0 *domain.Reputation
		rep1 *domain.Reputation
		rep2*domain.Reputation
	)

	rep0,rep1,rep2= buildTestReputation()

	reputationSlice := []*domain.Reputation{rep0,rep1}
        
	fmt.Println(len(reputationSlice))

	descr, result := rep2.UpdateReputation(reputationSlice)

	fmt.Println(len(reputationSlice))
	fmt.Println(descr)


	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}	

func TestDeleteReputation(t *testing.T){		

	var (
		rep0 *domain.Reputation
		rep1 *domain.Reputation
		rep2 *domain.Reputation
	)

	rep0,rep1,rep2 = buildTestReputation()		

	repSlice := []*domain.Reputation {rep0,rep1,rep2}	

	descr, result := repSlice[1].DeleteReputation(repSlice)
	
	fmt.Println(descr)

	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)		
	}	
}







