package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	testvar := viper.Get("element")
	//[map[one:map[url:http://test nested:123]] map[two:map[url:http://test nested:123]]]
	fmt.Println(testvar)
	elementsMap := testvar.([]interface{})
	for k, vmap := range elementsMap {
		fmt.Print("Key: ", k) // 0, 1
		// map[one:map[url:http://test nested:123]], map[two:map[url:http://test nested:123]]
		fmt.Println(" Value: ", vmap)
		eachElementsMap := vmap.(map[interface{}]interface{})

		for k, vEachValMap := range eachElementsMap {
			//one: map[url:http://test nested:123], two: map[url:http://test nested:123]
			fmt.Printf("%v: %v \n", k, vEachValMap)
			vEachValDataMap := vEachValMap.(map[interface{}]interface{})

			for k, v := range vEachValDataMap {
				//url: http://test
				//nested: 123
				fmt.Printf("%v: %v \n", k, v)
			}
		}
	}
}

// Output:
/*
Key: 0 Value:  map[one:map[url:http://test nested:123]]
one: map[url:http://test nested:123]
url: http://test
nested: 123
Key: 1 Value:  map[two:map[url:http://test nested:123]]
two: map[url:http://test nested:123]
url: http://test
nested: 123
*/
