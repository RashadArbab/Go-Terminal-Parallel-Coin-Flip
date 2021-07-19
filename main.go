package main 

import ("math/rand"
	"fmt"
	"time"
	"sync"
	"runtime"
)




func main () {

	var batchSize int = 0 
	var repetitions int = 0
	var addSet bool = false 
	var numThreads = 0 
	var limMem = false
	greeting(&batchSize, &repetitions , &addSet)
	fmt.Println("")
	if (addSet == true){
		additionalSettings(&numThreads , &limMem)
		runtime.GOMAXPROCS(numThreads)
	}



	fmt.Println("Please wait while we flip some coins...")
	var wg sync.WaitGroup
	var Oavg  float32
	Oavg = 0 
	start := time.Now()
	for i:= 0 ; i <=(repetitions -1) ; i++ {
		wg.Add(1)
		//fmt.Println("i: " , i)
		//fmt.Println("iteration: " , i) 
		go repeatFlip(&wg , &Oavg , batchSize)
		if limMem && i%500 == 0{
			runtime.GC() 
			fmt.Println("taking out the trash")
		}

	}
	wg.Wait() 
	elapsed := time.Since(start)
	fmt.Println("Time: " , elapsed)
	fmt.Println("The ratio of heads to tails is " , Oavg/float32(repetitions))

}

func greeting (batchSize *int , repetitions *int , addSet *bool) {
	fmt.Println("***************** Welcome to Flip *****************")
	fmt.Println("~The objective of this program is to determine how random is random.")
	fmt.Println("~We determine the randomness of rand ints in golang by")
	fmt.Println("~how close we are to a 1:1 ratio of heads to tails.")
	fmt.Println("~The experiment is run in batches and repeated multiple times for accuracy")
	fmt.Println("")
	var bs int = 0
	var rp int = 0
	var as string = ""
	fmt.Println("Be careful this program flips coins in parallel can your computer handle it?")
	fmt.Println("")
	fmt.Println("~How large would you like the batch size to be?")
	fmt.Scanln(&bs)
	fmt.Println("~How many times would you like to repeat the experiment?")
	fmt.Scanln(&rp)
	fmt.Println("~Would you like to tune additional settings? [y/n]")
	fmt.Scanln(&as)
	if (as == "y" || as == "Y"){
		*addSet = true 
	}else if (as == "n" || as == "N") {
		fmt.Println("~Ok lets begin")
		*addSet = false
	}else {
		fmt.Println("~Invalid input please try again? [y/n]")
		fmt.Scanln(&as)
	}


	*batchSize = bs 
	*repetitions = rp
}

func additionalSettings (maxThreads *int , limMem *bool){
	var numThreads int = 0
	fmt.Println("This is for advanced users, beware this may crash your computer.")
	fmt.Println("You have " , runtime.NumCPU() , " threads available.")
	fmt.Println("How many threads would you like to utilize")
	fmt.Scanln(&numThreads)
	*maxThreads = numThreads
	fmt.Println("Would you like to reduce memory use, this will drastically slow performance")
	fmt.Println("however you will be able to run much larger sets without a heap overflow [y/n]")
	memInput(limMem)
	
	
	

}

func memInput (limMem *bool) {
	var limitMem string 
	fmt.Scanln(&limitMem)
	if (limitMem == "y" || limitMem == "Y"){
		*limMem = true
	}else if (limitMem == "n" || limitMem == "N"){
		*limMem = false
	}else {
		memInput(limMem)
	}
}

func repeatFlip (wg *sync.WaitGroup, Oavg *float32 , batchSize int) {
	defer wg.Done()
	var heads int = 0 
	var tails int = 0
	
	var wg2 sync.WaitGroup
	
	for i:= 0 ; i <=(batchSize -1) ; i ++ {
		wg2.Add(1)
		//fmt.Println("j: " , i)
		go flipCoin(&heads , &tails , &wg2)
	}
	wg2.Wait()
	//fmt.Println("Heads: " ,heads)
	//fmt.Println("Tails: " ,tails)
	//fmt.Println("Average: " , float32(heads)/float32(tails))
	newAvg := float32(*Oavg + (float32(heads)/float32(tails)))
	*Oavg = newAvg
}


func  flipCoin (heads *int, tails *int, wg2 *sync.WaitGroup) {
	defer wg2.Done()
	rand.Seed(time.Now().UnixNano())
	v:= rand.Intn(2)  


	if v == 1 {
		h := *heads +1 
		*heads = h
		//fmt.Println("heads")
	}else {
		t := *tails +1 
		*tails = t  
		//fmt.Println("tails") 
	}
	

}
