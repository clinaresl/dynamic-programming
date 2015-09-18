/* 
  maximum_subarray.go
  Description: Solution to the Maximum Subarray Problem
  ----------------------------------------------------------------------------- 

  Started on  <Thu Sep 17 15:48:54 2015 Carlos Linares Lopez>
  Last update <viernes, 18 septiembre 2015 16:28:40 Carlos Linares Lopez (clinares)>
  -----------------------------------------------------------------------------

  $Id::                                                                      $
  $Date::                                                                    $
  $Revision::                                                                $
  -----------------------------------------------------------------------------

  Made by Carlos Linares Lopez
  Login   <clinares@atlas>
*/

package main

// imports
// ----------------------------------------------------------------------------
import (
	"bufio"			        // buffered input/output
	"bytes"
	"flag"				// arg parsing
	"fmt"				// formatted input/output
	"io"
	"log"				// logging services
	"os"				// operating system services
	
)


// global variables
// ----------------------------------------------------------------------------
var EXIT_SUCCESS int = 0		// exit with success
var helpRequired bool			// has --help been given?

// functions
// ----------------------------------------------------------------------------

// initializes the command-line parser
func init () {

	// Flag to provide additional help
	flag.BoolVar (&helpRequired, "help", false, "if given, additional information is provided")
}

// shows additional information about the problem and the input/output format
// followed
func showHelp (signal int) {

	fmt.Println (`
 Given an array X = {x1, x2, ..., xN} of N elements, find the maximum possible
 sum of a contiguous array.

 Empty subarrays/subsequences should not be considered. 

 #Input Format

 First line of the input has an integer T. T cases follow.

 Each test case begins with an integer N. In the next line, N integers follow
 representing the elements of array X.  

 #Output Format

 One integer denoting the maximum contiguous subarray. At least one integer
 should be selected and put into the subarray (this may be required in cases
 where all elements are negative).

 #Example

 In the following, ">" means input and "<" is the output of the program and are
 used as decorators to illustrate the example:

 >1
 >6
 >2 -3 1 7 -4 1
 <8

 #Usage

 Run the program without arguments to enter the cases
`)
	os.Exit (signal)
}

// this function returns in value the maximum sum over all items in a contiguous
// sub-array in sequence
func getContiguousMaxSubArray (sequence []int) (int) {

	// if the sequence is empty, return zero
	if len (sequence) == 0 {
		return 0
	}
	
	// incumbent contains the best solution found so far whereas current
	// contains the sum of the items currently under consideration. They are
	// both initialized to the first item in sequence
	current := sequence[0]
	incumbent := sequence[0]

	// and now process all items in the sequence from the second until the
	// last one
	for i := 1 ; i < len (sequence) ; i++ {

		// if adding the next item to the current serie improves the
		// sum, ...
		if current + sequence[i] >= sequence[i] {

			// ... then add this item
			current += sequence[i]
		} else {

			// and start a new serie from this point onwards
			current = sequence[i]
		}

		// at every step, this loop computes the maximum sum up to
		// location i in current. Compare it to incumbent so that the
		// largest sum is stored as the current solution
		if incumbent < current {
			incumbent = current
		}		
	}

	// and return the best solution found so far
	return incumbent
}

func main () {

	var text string
	var nbcases int

	// first, parse the flags and, in case that --help has been given, show
	// the help banner and exit
	flag.Parse ()
	if helpRequired {
		showHelp (EXIT_SUCCESS)
	}
	
	// create a reader to read from stdin
	reader := bufio.NewReader (os.Stdin)

	// read the number of cases
	text, _ = reader.ReadString ('\n')
	fmt.Sscanf (text, "%d", &nbcases)

	// Now, for each test case
	for icase := 0 ; icase < nbcases ; icase++ {

		var nbitems int
		
		// read the number of items in this test case
		text, _ = reader.ReadString ('\n')
		fmt.Sscanf (text, "%d", &nbitems)

		// and now read the entire sequence from a single line and
		// decode it as a sequence of ints. Be aware, lines can be
		// impressively long so that it might be necessary to
		// concatenate them
		var slice []byte
		for ;; {

			line, err := reader.ReadSlice ('\n')
			if err != nil && err != bufio.ErrBufferFull && err != io.EOF {
				log.Fatal (" Error while reading the slice!")
			}
			slice = append (slice, line...)
			if err == nil {
				break
			}
		}
		
		buf := bytes.NewBuffer (slice)
		sequence := make ([]int, nbitems)
		for iitem := 0 ; iitem < nbitems ; iitem++ {
			_, err := fmt.Fscan (buf, &sequence[iitem])
			if err != nil && err != io.EOF {
				log.Fatal (" Fatal error while scanning the sequence!")
			}
		}

		// and now, compute the solution to this specific case and show
		// it on stdout
		fmt.Println (getContiguousMaxSubArray (sequence))
	}
}


/* Local Variables: */
/* mode:go */
/* fill-column:80 */
/* End: */
