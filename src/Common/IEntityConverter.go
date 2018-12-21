package Common

/**
Convert app entity to database entity and back
*/
type IEntityConverter interface {
	/**
	Convert app entity to database
	*/
	ToDatabaseEntity(entity interface{}) (interface{}, error)

	/**
	Convert database entity to app
	*/
	ToAppEntity(entity interface{}) (interface{}, error)
}
