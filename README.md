go-diff
=======

A (soon to be) collection of helper functions to determine the deltas of various Go "objects"

## Usage

```
	import "github.com/jsimnz/go-delta"
    
    ..
    
	type someStruct struct {
    	Field1 string
        Field2 string
        Field3 int
    }
    
    ...
    
	val1 := someStruct{Field1: "Hello"}
    val2 := someStruct{Field1: "World"}
    
    diff, err := delta.Struct(val1, val2) // returns the delta (difference) between the 2 structs as a map
    
    // returns: 
    //	map[string]interface{}{
    //	   "Field1": "World",
    //    }
    // 
    // Which tells us that the second struct differs from the first by the field 'Field1',
    // and the new value is "World"
```

### More delta functions to come...

## Documentation
GoDoc: http://godoc.org/github.com/jsimnz/go-delta

## Todo
- More tests!

## Author
 **John-Alan Simmons**
 
 - [http://github.com/jsimnz](http://github.com/jsimnz)
 - [simmons.johnalan@gmail.com](mailto:simmons.johnalan@gmail.com)
 - [@iamjsimnz](http://twitter.com/iamjsimnz)