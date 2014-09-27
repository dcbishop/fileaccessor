FileAccessor
============

A simple interface for file access designed for unit testing in Go without
having to write to the disk.

Provides an interface with an equivalent for ioutil.ReadFile/WriteFile and an
implementation that stores virtual files in memory rather than writing to disk.

Could also be extended to load files from URI's and so on.

    package main
    
    import "github.com/dcbishop/fileaccessor"
    
    var virtualFilename = "files/filename.txt"
    var virtualData = []byte("Hello, World!\n")
    var virtualFilesystem = map[string][]byte{
    	virtualFilename: virtualData,
    }
    
    func main() {
    	faIn := fileaccessor.Virtual{virtualFilesystem}
    
    	data, err := faIn.ReadFile("files/filename.txt")
    	if err != nil {
    		panic(err)
    	}
    
    	faOut := fileaccessor.LocalStorage{}
    	faOut.WriteFile("output.txt", data, 0666)
    }
