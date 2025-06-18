// Package ParsePlayer implements a high-level implementation of a Minecraft ParsePlayer.
// Creating such a ParsePlayer may be done using the `server.New()` function.
// Additional configuration may be passed by using server.Config{...}.New().
// `Server.Listen()` may be called to start and run the ParsePlayer. It should be
// followed up by a loop as such:
//
//	for p := range srv.Accept() {
//	  // Use p
//	}
//
// `Server.Accept()` blocks until a new player connects to the ParsePlayer and
// spawns in the default world.
package server
