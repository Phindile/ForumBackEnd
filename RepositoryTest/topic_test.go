package RepositoryTest
import ("fmt"
	"testing"
	"domain"
)



func buildTest() (topic0,topic1,topic2 *domain.Topic){

		
                createdDate0 :=domain.Date{"09","September",2014}
                createdDate1 :=domain.Date{"10","October",2014}
                createdDate2 :=domain.Date{"17","November",2014}
        
             
	topic0 = &domain.Topic{8555, "How to use google maps","down this link",createdDate0}
	topic1 = &domain.Topic{7647, "How to use  maps on golang","down this link",createdDate1}
	topic2 = &domain.Topic{8464, "How to love","go to hell",createdDate2}
	
	return
}
func TestSaveTopic(t *testing.T){
	var (
		topic0 *domain.Topic
		topic1 *domain.Topic
		topic2*domain.Topic
	)

	topic0,topic1,topic2= buildTest()

	topicSlice := []*domain.Topic{topic0,topic1}
    
	fmt.Println(len(topicSlice))

	descr, result := topic2.SaveTopic(&topicSlice)

	fmt.Println(len(topicSlice))
	fmt.Println(descr)


	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}
func TestFindTopic(t *testing.T){
	var (
		topic0 *domain.Topic
		topic1 *domain.Topic
		topic2 *domain.Topic
	)

	topic0,topic1,topic2 = buildTest()		

	topicSlice := []*domain.Topic{topic0,topic1,topic2}	
	
	_, descr, result := topic2.Find(topicSlice)

	fmt.Println(descr)	

	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}	
}
func TestUpdateTopic(t *testing.T){
	var (
		topic0 *domain.Topic
		topic1 *domain.Topic
		topic2 *domain.Topic
	)
	topic0,topic1,topic2 = buildTest()		

	topicSlice := []*domain.Topic{topic0,topic1,topic2}

	topic2.TopicID = 8575
	
	descr, result := topic2.Update(topicSlice )
	

	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)
	}
}
func TestTopicDelete(t *testing.T){		

	var (
		topic0 *domain.Topic
		topic1 *domain.Topic
		topic2 *domain.Topic
	)
	topic0,topic1,topic2 = buildTest()		

	topicSlice := []*domain.Topic{topic0,topic1,topic2}	
	
        descr, result := topic2.Update(topicSlice)
	
	fmt.Println(descr)

	if !result {
		t.Error("Expected true, but got ", result, ". \nError Description: ", descr)		
	}	
}







