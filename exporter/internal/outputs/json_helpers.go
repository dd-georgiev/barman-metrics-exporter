package outputs

// A few types to improve readability. Those are related to helping parse JSON ouput from barman
type jsonKeyValuePair map[string]string
type jsonKeyInterfacePair map[string]interface{}
type jsonKeyObjectPair map[string]jsonKeyValuePair
