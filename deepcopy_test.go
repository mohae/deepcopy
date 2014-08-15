package deepcopy

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var sliceStringsOrig = []string{
	"element1",
	"element2",
	"element3",
}

var sliceIntsOrig = []int{
	1,
	2,
	3,
}

var mapStringInterfaceSliceStrings = map[string]interface{}{
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

var mapStringInterfaceSliceInts = map[string]interface{}{
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

var mapStringInterfaceSliceBools = map[string]interface{}{
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
			"sliceStrings": []string{
				"value1",
				"value2",
				"value3",
			},
			"sliceInts": []int{
				1,
				2,
				3,
			},

		},
	},	
	"mapIface2": map[string]interface{}{
		"sliceStrings": []string{
			"val1",
			"val2",
			"val3",
		},
		"sliceInts": []int{
			1,
			2,
			3,
		},
	},	
	"slice3": []string{
		"v1",
		"v2",
		"v3",
	},	
}
//var origMapStringInterface = map[string]interface{}{
//}

func TestSliceStrings(t *testing.T) {
	Convey("Given a nil slice", t, func() {
		Convey("Deep copying it", func() {
			copy := SliceStrings(nil)
			Convey("Should result in nil", func() {
				So(copy, ShouldBeNil)
			})
		})
	})

	Convey("Given a slice of type string", t, func() {
		Convey("Deep copying it", func() {
			copy := SliceStrings(sliceStringsOrig)
			Convey("Should result in a copy", func() {
				So(copy, ShouldResemble, sliceStringsOrig)
			})
			Convey("And modifying the copy", func() {
				copy[1] = "elemental"
				Convey("Should not change the original", func() {
					So(sliceStringsOrig[1], ShouldEqual, "element2")
				})
			})
		})
	})
}

func TestInterfaceSliceStrings(t *testing.T) {
	Convey("Given a nil slice", t, func() {
		Convey("Deep copying it", func() {
			copy := InterfaceToSliceStrings(nil)
			Convey("Should result in nil", func() {
				So(copy, ShouldBeNil)
			})
		})
	})

	Convey("Given a non-slice value", t, func() {
		Convey("Deep copying it with InterfaceToSliceStrings", func() {
			copy := InterfaceToSliceStrings(1233)
			Convey("Should result in nil", func() {
				So(copy, ShouldBeNil)
			})
		})
	})

	Convey("Given a slice of type string", t, func() {
		Convey("Deep copying it", func() {
			copy := InterfaceToSliceStrings(sliceStringsOrig)
			Convey("Should result in a copy", func() {
				So(copy, ShouldResemble, sliceStringsOrig)
			})
			Convey("And their addresses should be different", func() {
				So(&copy, ShouldNotPointTo, &sliceStringsOrig)
			})
			Convey("And modifying the copy", func() {
				copy[1] = "elemental"
				Convey("Should not change the original", func() {
					So(sliceStringsOrig[1], ShouldEqual, "element2")
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

func TestSliceInts(t *testing.T) {
	Convey("Given a nil slice", t, func() {
		Convey("Deep copying it", func() {
			copy := SliceInts(nil)
			Convey("Should result in nil", func() {
				So(copy, ShouldBeNil)
			})
		})
	})

	Convey("Given a slice of type int", t, func() {
		Convey("Deep copying it", func() {
			copy := SliceInts(sliceIntsOrig)
			Convey("Should result in a copy", func() {
				So(copy, ShouldResemble, sliceIntsOrig)
			})
			Convey("And their addresses should be different", func() {
				So(&copy, ShouldNotPointTo, &sliceIntsOrig)
			})
			Convey("And modifying the copy", func() {
				copy[1] = 333
				Convey("Should not change the original", func() {
					So(sliceIntsOrig[1], ShouldEqual, 2)
				})
			})
		})
	})
}

func TestInterfaceSliceInts(t *testing.T) {
	Convey("Given a nil slice", t, func() {
		Convey("Deep copying it", func() {
			copy := InterfaceToSliceInts(nil)
			Convey("Should result in nil", func() {
				So(copy, ShouldBeNil)
			})
		})
	})

	Convey("Given a non-slice value", t, func() {
		Convey("Deep copying it with InterfaceToSliceInts", func() {
			copy := InterfaceToSliceInts("test")
			Convey("Should result in nil", func() {
				So(copy, ShouldBeNil)
			})
		})
	})

	Convey("Given a slice of type int", t, func() {
		Convey("Deep copying it", func() {
			copy := InterfaceToSliceInts(sliceIntsOrig)
			Convey("Should result in a copy", func() {
				So(copy, ShouldResemble, sliceIntsOrig)
			})
			Convey("And modifying the copy", func() {
				copy[1] = 3333
				Convey("Should not change the original", func() {
					So(sliceIntsOrig[1], ShouldEqual, 2)
				})
			})
		})
	})
}	

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
			copy := MapStringInterface(mapStringInterfaceSliceStrings)
			Convey("Should result in a copy", func() {
				So(copy, ShouldResemble, mapStringInterfaceSliceStrings)
			})
			Convey("Modifying the copy should not change the original", func() {
			})
		})

		Convey("Deep copying a map[string]interface with []ints", func() {
			copy := MapStringInterface(mapStringInterfaceSliceInts)
			Convey("Should result in a copy", func() {
				So(copy, ShouldResemble, mapStringInterfaceSliceInts)
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
				copy := MapStringInterface(mapStringInterfaceSliceBools)
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
