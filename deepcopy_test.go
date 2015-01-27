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
			Convey("should result in nil", func() {
				So(copy, ShouldResemble, stringSliceOrig)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(intSliceOrig)
			Convey("should result in nil", func() {
				So(copy, ShouldResemble, intSliceOrig)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(mapStringInterfaceStringSlice)
			Convey("should result in nil", func() {
				So(copy, ShouldResemble, mapStringInterfaceStringSlice)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(mapStringInterfaceIntSlice)
			Convey("should result in nil", func() {
				So(copy, ShouldResemble, mapStringInterfaceIntSlice)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(mapStringInterfaceBoolSlice)
			Convey("should result in nil", func() {
				So(copy, ShouldResemble, mapStringInterfaceBoolSlice)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(stringSliceOrig)
			Convey("should result in nil", func() {
				So(copy, ShouldResemble, stringSliceOrig)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(mapStringInterfaceMaps)
			Convey("should result in nil", func() {
				So(copy, ShouldResemble, mapStringInterfaceMaps)
			})
		})
	})

	Convey("Given a slice of strings", t, func() {
		Convey("copying it", func() {
			copy := Iface(testS)
			Convey("should result in nil", func() {
				So(copy, ShouldResemble, testS)
			})
		})
	})
}
