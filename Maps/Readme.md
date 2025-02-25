# Maps
Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.

The zero value of a map is nil.

We can create a map by using a literal or by using the make() function:

ages := make(map[string]int)
ages["John"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 // overwrites 24

ages = map[string]int{
  "John": 37,
  "Mary": 21,
}

# Mutations
Insert an Element
m[key] = elem

Get an Element
elem = m[key]

Delete an Element
delete(m, key)

Check If a Key Exists
elem, ok := m[key]

# Note on Passing Maps
Like slices, maps are also passed by reference into functions. This means that when a map is passed into a function we write, we can make changes to the original, we don't have a copy.

# Key Types
Any type can be used as the value in a map, but keys are more restrictive.

As mentioned earlier, map keys may be of any type that is comparable. The language spec defines this precisely, but in short, comparable types are boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types. Notably absent from the list are slices, maps, and functions; these types cannot be compared using ==, and may not be used as map keys.

It's obvious that strings, ints, and other basic types should be available as map keys, but perhaps unexpected are struct keys. Struct can be used to key data by multiple dimensions. For example, this map of maps could be used to tally web page hits by country:

hits := make(map[string]map[string]int)

This is a map of string to (map of string to int). Each key of the outer map is the path to a web page with its own inner map. Each inner map key is a two-letter country code. This expression retrieves the number of times an Australian has loaded the documentation page:

n := hits["/doc/"]["au"]

Unfortunately, this approach becomes unwieldy when adding data, as for any given outer key you must check if the inner map exists, and create it if needed:

func add(m map[string]map[string]int, path, country string) {
    mm, ok := m[path]
    if !ok {
        mm = make(map[string]int)
        m[path] = mm
    }
    mm[country]++
}
add(hits, "/doc/", "au")

On the other hand, a design that uses a single map with a struct key does away with all that complexity:

type Key struct {
    Path, Country string
}
hits := make(map[Key]int)

When a Vietnamese person visits the home page, incrementing (and possibly creating) the appropriate counter is a one-liner:

hits[Key{"/", "vn"}]++

And it’s similarly straightforward to see how many Swiss people have read the spec:

n := hits[Key{"/ref/spec", "ch"}]

# Count Instances
Remember that you can check if a key is already present in a map by using the second return value from the index operation.

You can combine an if statement with an assignment operation to use the variables inside the if block:

names := map[string]int{}
missingNames := []string{}

if _, ok := names["Denna"]; !ok {
    // if the key doesn't exist yet,
    // append the name to the missingNames slice
    missingNames = append(missingNames, "Denna")
}

# Map Literals
Maps can be constructed using the usual composite literal syntax with colon-separated key-value pairs, so it's easy to build them during initialization.

var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}

# Missing Keys
An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map. For instance, if the map contains integers, looking up a non-existent key will return 0. A set can be implemented as a map with value type bool. Set the map entry to true to put the value in the set, and then test it by simple indexing.

attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // will be false if person is not in the map
    fmt.Println(person, "was at the meeting")
}

Sometimes you need to distinguish a missing entry from a zero value. Is there an entry for "UTC" or is that 0 because it's not in the map at all? You can discriminate with a form of multiple assignment.

var seconds int
var ok bool
seconds, ok = timeZone[tz]

For obvious reasons, this is called the “comma ok” idiom. In this example, if tz is present, seconds will be set appropriately and ok will be true; if not, seconds will be set to zero and ok will be false. Here's a function that puts it together with a nice error report:

func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}

# Deleting Map Entries
To delete a map entry, use the delete built-in function, whose arguments are the map and the key to be deleted. It's safe to do this even if the key is already absent from the map.

delete(timeZone, "PDT")  // Now on Standard Time

# Nested
Maps can contain maps, creating a nested structure. For example:

map[string]map[string]int
map[rune]map[string]int
map[int]map[string]map[string]int