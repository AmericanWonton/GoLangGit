//////////////////
//Building Errors
/*
Anything that uses type error functions, is now also of type error
*/

///////////////////
//Checking Errors

/*
	Typically, we don't check erros of fmt
*/
n, err := fmt.Println("Hello")
if err != nil {
	fmt.Println(err)
}
fmt.Println(n)
///////////////////////////
f, err := os.Create("names.txt") //This creates a new file from OS package
if err != nil {
	fmt.Println(err)
	return //close program
}
defer f.Close() //This closes the file at the end of the main fucntion
//so you aren't wasting resources

r := strings.NewReader("James Bond")

io.Copy(f, r) //This reads from that r string and copies the result
//to the f file we opened earlier.

f, err := os.Open("names.txt") //Opens an existing file
if err != nil {
	fmt.Println(err)
	return //Ends main function
}

defer f.Close() //Close this when done

bs, err := ioutil.ReadAll(f) //Takes in a reader and returns
//a byteslice and error.

if err != nil {
	fmt.Println(err)
}

fmt.Println(string(bs)) //This converts the byteslice to a string
//and prints it out.

///////////////////////////
//How to print out errors.

defer foo()
_, err := os.Open("no-file.txt") //This file dosen't exist yet.
if err != nil {
	fmt.Println("err happended", err)
	// log.Fatalln(err) //CLOSES THE PROGRAM AFTER WRITING, NO FUNCITON DONE.
	//panic(err)
}

f, err := os.Create("log.txt") //Create a sample log file to write to.
if err != nil {
	fmt.Println(err)
}
defer f.Close()
log.SetOutput(f)

fmt.Println("Check the log.txt file in the directory.")



func foo() {
	fmt.Println("When os.Exit() is called, defered functions don't run.")
}

//Panic
_, err := os.Open("No-file.txt")
if err != nil {
	log.Panic(err)
}

///////////////////////////
//Recover
//Fuck dude, IDK, the lesson sucked.


///////////////////////////
//Errors with info.
//Maybe do this later, IDK...

