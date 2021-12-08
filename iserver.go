package zinxFace

// IServer define a server interface
type IServer interface {

	// Start set up the server
	Start()

	// Stop stop the server
	Stop()

	// Run run the server
	Run()
}