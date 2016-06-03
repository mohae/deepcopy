package deepcopy

import (
	_ "reflect"
	"testing"

	_ "github.com/mohae/customjson"
	. "github.com/smartystreets/goconvey/convey"
)

var stringSliceOrig = []string{
	"element1",
	"element2",
	"element3",
}

var intSliceOrig = []int{
	1,
	2,
	3,
}

var mapStringInterfaceStringSlice = map[string]interface{}{
	"slice1": []string{
		"value1",
		"value2",
		"value3",
	},
	"slice2": []string{
		"val1",
		"val2",
		"val3",
	},
	"slice3": []string{
		"v1",
		"v2",
		"v3",
	},
}

var mapStringInterfaceIntSlice = map[string]interface{}{
	"slice1": []int{
		1,
		2,
		3,
	},
	"slice2": []int{
		11,
		22,
		33,
	},
	"slice3": []int{
		111,
		222,
		333,
	},
}

var mapStringInterfaceBoolSlice = map[string]interface{}{
	"slice1": []bool{
		true,
		true,
		true,
	},
	"slice2": []bool{
		false,
		false,
		false,
	},
}

var mapStringInterfaceMaps = map[string]interface{}{
	"mapIface": map[string]interface{}{
		"mapIface-2": map[string]interface{}{
			"stringSlice": []string{
				"value1",
				"value2",
				"value3",
			},
			"intSlice": []int{
				1,
				2,
				3,
			},
		},
	},
	"mapIface2": map[string]interface{}{
		"stringSlice": []string{
			"val1",
			"val2",
			"val3",
		},
		"intSlice": []int{
			1,
			2,
			3,
		},
	},
	"mapIface3": map[string]interface{}{},
	"slice3": []string{
		"v1",
		"v2",
		"v3",
	},
	"slice4": []float64{
		1.232132,
		123.2322,
	},
}

type testStruct struct {
	TestName string
	Tests    []*Testz
}

type Testz struct {
	Expected   string
	Unexpected []string
	Done       bool
}

var testS = testStruct{
	TestName: "testname",
	Tests: []*Testz{
		&Testz{
			Expected: "expected1",
			Unexpected: []string{
				"oops1",
				"uhoh1",
			},
			Done: true,
		},
		&Testz{
			Expected: "expected2",
			Unexpected: []string{
				"oops2",
				"uhoh2",
			},
			Done: false,
		},
	},
}

//var origMapStringInterface = map[string]interface{}{
//}

func TestStringSlice(t *testing.T) {
	Convey("Given a nil slice", t, func() {
		Convey("Deep copying it", func() {
			copy := StringSlice(nil)
			Convey("Should result in nil", func() {
				So(copy, ShouldBeNil)
			})
		})
	})

	Convey("Given a slice of type string", t, func() {
		Convey("Deep copying it", func() {
			copy := StringSlice(stringSliceOrig)
			Convey("Should result in a copy", func() {
				So(copy, ShouldResemble, stringSliceOrig)
			})
			Convey("And modifying the copy", func() {
				copy[1] = "elemental"
				Convey("Should not change the original", func() {
					So(stringSliceOrig[1], ShouldEqual, "element2")
				})
			})
		})
	})
}

func TestInterfaceStringSlice(t *testing.T) {
	Convey("Given a nil slice", t, func() {
		Convey("Deep copying it", func() {
			copy := InterfaceToStringSlice(nil)
			Convey("Should result in nil", func() {
				So(copy, ShouldBeNil)
			})
		})
	})

	Convey("Given a non-slice value", t, func() {
		Convey("Deep copying it with InterfaceToStringSlice", func() {
			copy := InterfaceToStringSlice(1233)
			Convey("Should result in nil", func() {
				So(copy, ShouldBeNil)
			})
		})
	})

	Convey("Given a slice of type string", t, func() {
		Convey("Deep copying it", func() {
			copy := InterfaceToStringSlice(stringSliceOrig)
			Convey("Should result in a copy", func() {
				So(copy, ShouldResemble, stringSliceOrig)
			})
			Convey("And their addresses should be different", func() {
				So(&copy, ShouldNotPointTo, &stringSliceOrig)
			})
			Convey("And modifying the copy", func() {
				copy[1] = "elemental"
				Convey("Should not change the original", func() {
					So(stringSliceOrig[1], ShouldEqual, "element2")
				})
			})
		})
	})
}

/*
func TestMapStringInterface(t *testing.T) {
	Convey("Given a map[string]interface{} with a slice", t, func() {
		Con
}
*/

func TestIntSlice(t *testing.T) {
	Convey("Given a nil slice", t, func() {
		Convey("Deep copying it", func() {
			copy := IntSlice(nil)
			Convey("Should result in nil", func() {
				So(copy, ShouldBeNil)
			})
		})
	})

	Convey("Given a slice of type int", t, func() {
		Convey("Deep copying it", func() {
			copy := IntSlice(intSliceOrig)
			Convey("Should result in a copy", func() {
				So(copy, ShouldResemble, intSliceOrig)
			})
			Convey("And their addresses should be different", func() {
				So(&copy, ShouldNotPointTo, &intSliceOrig)
			})
			Convey("And modifying the copy", func() {
				copy[1] = 333
				Convey("Should not change the original", func() {
					So(intSliceOrig[1], ShouldEqual, 2)
				})
			})
		})
	})
}

func TestInterfaceIntSlice(t *testing.T) {
	Convey("Given a nil slice", t, func() {
		Convey("Deep copying it", func() {
			copy := InterfaceToIntSlice(nil)
			Convey("Should result in nil", func() {
				So(copy, ShouldBeNil)
			})
		})
	})

	Convey("Given a non-slice value", t, func() {
		Convey("Deep copying it with InterfaceToIntSlice", func() {
			copy := InterfaceToIntSlice("test")
			Convey("Should result in nil", func() {
				So(copy, ShouldBeNil)
			})
		})
	})

	Convey("Given a slice of type int", t, func() {
		Convey("Deep copying it", func() {
			copy := InterfaceToIntSlice(intSliceOrig)
			Convey("Should result in a copy", func() {
				So(copy, ShouldResemble, intSliceOrig)
			})
			Convey("And modifying the copy", func() {
				copy[1] = 3333
				Convey("Should not change the original", func() {
					So(intSliceOrig[1], ShouldEqual, 2)
				})
			})
		})
	})
}

/*
func TestMapStringInterface(t *testing.T) {
	Convey("Give a nil value", t, func() {
		Convey("Deep copying it", func() {
			copy := MapStringInterface(nil)
			Convey("Should result in nil", func() {
				So(copy, ShouldBeNil)
			})
		})
	})

	Convey("Given map[string]interfaces{}", t, func() {
		Convey("Deep copying a map[string]interface with []strings", func() {
			copy := MapStringInterface(mapStringInterfaceStringSlice)
			Convey("Should result in a copy", func() {
				So(copy, ShouldResemble, mapStringInterfaceStringSlice)
			})
			Convey("Modifying the copy should not change the original", func() {
			})
		})

		Convey("Deep copying a map[string]interface with []ints", func() {
			copy := MapStringInterface(mapStringInterfaceIntSlice)
			Convey("Should result in a copy", func() {
				So(copy, ShouldResemble, mapStringInterfaceIntSlice)
			})
			Convey("Modifying the copy should not change the original", func() {
			})
		})

		Convey("Deep copying a map[string]interface with variable slice types and embedded map[string]interfaces", func() {
			copy := MapStringInterface(mapStringInterfaceMaps)
			Convey("Should result in a copy", func() {
				So(copy, ShouldResemble, mapStringInterfaceMaps)
			})
			Convey("Modifying the copy", func() {
				copy["mapIface"] = nil
				copy["slice3"].([]string)[0] = "123"
				Convey("Should not change the original", func() {
					So(copy["mapIface"], ShouldNotResemble, mapStringInterfaceMaps["mapIface"])
					So(mapStringInterfaceMaps["mapIface"], ShouldNotBeNil)
					So(copy, ShouldNotResemble, mapStringInterfaceMaps)
				})
			})
		})

		Convey("Deep copying an unsupported interface type", func() {
			Convey("Deep copying a map[string]interface with []bool", func() {
				copy := MapStringInterface(mapStringInterfaceBoolSlice)
				Convey("Should result in a not supported meassage", func() {
					So(copy["slice1"], ShouldEqual, "Error unsupported Type: []bool is not supported.")
					So(copy["slice2"], ShouldEqual, "Error unsupported Type: []bool is not supported.")
				})
				Convey("Modifying the copy should not change the original", func() {
				})
			})
		})
	})
}
*/
/*
func TestInterfaceToSliceInterfaces(t * testing.T) {
	Convey("Given an interface with slices", t, func() {
		Convey("copying it to  a slice of interfaces", func() {
			copy :=  InterfaceToSliceInterfaces(mapStringInterfaceStringSlice["slice1"])
//			copy2 := InterfaceToStringSlice(copy)
			Convey("Should result in a copy", func() {
				So(reflect.ValueOf(copy[0]).Interface(), ShouldResemble, "value1")
//				So(reflect.KindOf(copy[0]).Interface(), ShouldResemble, "")
				So(copy, ShouldContain, "value2")
				So(copy, ShouldContain, "value3")
			})
		})
	})
}
*/

func TestRecursiveCopy(t *testing.T) {
	Convey("Given a nil pointer to an interface", t, func() {
		var tst interface{}
		Convey("copying it", func() {
			copy := Iface(tst)
			Convey("should result in nil", func() {
				So(copy, ShouldBeNil)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(stringSliceOrig)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, stringSliceOrig)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(intSliceOrig)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, intSliceOrig)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(mapStringInterfaceStringSlice)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, mapStringInterfaceStringSlice)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(mapStringInterfaceIntSlice)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, mapStringInterfaceIntSlice)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(mapStringInterfaceBoolSlice)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, mapStringInterfaceBoolSlice)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(stringSliceOrig)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, stringSliceOrig)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(mapStringInterfaceMaps)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, mapStringInterfaceMaps)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(testS)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testS)
			})
		})
	})

	Convey("Given a type-aliased string", t, func() {
		type aliased string
		testValue := aliased("hello")
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a type-aliased int", t, func() {
		type aliased int
		testValue := aliased(42)
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a type-aliased int8", t, func() {
		type aliased int8
		testValue := aliased(42)
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a type-aliased int16", t, func() {
		type aliased int16
		testValue := aliased(42)
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a type-aliased int32", t, func() {
		type aliased int32
		testValue := aliased(42)
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a type-aliased int64", t, func() {
		type aliased int64
		testValue := aliased(42)
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a type-aliased uint", t, func() {
		type aliased uint
		testValue := aliased(42)
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a type-aliased uint8", t, func() {
		type aliased uint8
		testValue := aliased(42)
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a type-aliased uint16", t, func() {
		type aliased uint16
		testValue := aliased(42)
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a type-aliased uint32", t, func() {
		type aliased uint32
		testValue := aliased(42)
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a type-aliased uint64", t, func() {
		type aliased uint64
		testValue := aliased(42)
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a type-aliased float32", t, func() {
		type aliased float32
		testValue := aliased(42.0)
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a type-aliased float64", t, func() {
		type aliased float64
		testValue := aliased(42.0)
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a type-aliased bool", t, func() {
		type aliased bool
		testValue := aliased(true)
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a string", t, func() {
		var testValue string = "hello"
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a int", t, func() {
		var testValue int = 42
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a int8", t, func() {
		var testValue int8 = 42
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a int16", t, func() {
		var testValue int16 = 42
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a int32", t, func() {
		var testValue int32 = 42
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a int64", t, func() {
		var testValue int64 = 42
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a uint", t, func() {
		var testValue uint = 42
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a uint8", t, func() {
		var testValue uint8 = 42
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a uint16", t, func() {
		var testValue uint16 = 42
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a uint32", t, func() {
		var testValue uint32 = 42
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a uint64", t, func() {
		var testValue uint64 = 42
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a float32", t, func() {
		var testValue float32 = 42.0
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a float64", t, func() {
		var testValue float64 = 42.0
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})

	Convey("Given a bool", t, func() {
		var testValue bool = true
		Convey("copying it", func() {
			copy := Iface(testValue)
			Convey("should result in result resembling original", func() {
				So(copy, ShouldResemble, testValue)
			})
		})
	})
}
