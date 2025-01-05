package page

import "fyne.io/fyne/v2"

// Tutorial defines the data structure for a tutorial
type Tutorial struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

// var (
// 	// Tutorials defines the metadata for each tutorial
// 	Tutorials = map[string]Tutorial{
// 		"welcome": {"首页信息", "", WelcomeScreen},
// 		// "canvas": {"Canvas",
// 		// 	"See the canvas capabilities.",
// 		// 	canvasScreen,
// 		// },
// 		// "animations": {"Animations",
// 		// 	"See how to animate components.",
// 		// 	makeAnimationScreen,
// 		// },
// 		// "icons": {"Theme Icons",
// 		// 	"Browse the embedded icons.",
// 		// 	iconScreen,
// 		// },
// 		// "tree": {"Tree",
// 		// 	"A tree based arrangement of cached elements with the same styling.",
// 		// 	makeTreeTab,
// 		// },
// 		// "dialogs": {"Dialogs",
// 		// 	"Work with dialogs.",
// 		// 	dialogScreen,
// 		// },
// 		// "windows": {"Windows",
// 		// 	"Window function demo.",
// 		// 	windowScreen,
// 		// },
// 		// "binding": {"Data Binding",
// 		// 	"Connecting widgets to a data source.",
// 		// 	bindingScreen},
// 		// "advanced": {"Advanced",
// 		// 	"Debug and advanced information.",
// 		// 	advancedScreen,
// 		// },
// 	}

// 	// // TutorialIndex  defines how our tutorials should be laid out in the index tree
// 	// TutorialIndex = map[string][]string{
// 	// 	"":            {"welcome", "canvas", "animations", "icons", "widgets", "collections", "containers", "dialogs", "windows", "binding", "advanced"},
// 	// 	"collections": {"list", "table", "tree"},
// 	// 	"containers":  {"apptabs", "border", "box", "center", "doctabs", "grid", "scroll", "split"},
// 	// 	"widgets":     {"accordion", "button", "card", "entry", "form", "input", "progress", "text", "toolbar"},
// 	// }
// )
