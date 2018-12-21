package Common

/**
Convert struct to map and back
*/
type IHydrator interface {
	/**
	Convert data to struct
	*/
	Create(data map[string]interface{}) (interface{}, error)

	/**
	Map data to object
	*/
	Hydrate(data map[string]interface{}, object interface{}) (interface{}, error)

	/**
	Convert struct to map
	*/
	Extract(object interface{}) (map[string]interface{}, error)
}
