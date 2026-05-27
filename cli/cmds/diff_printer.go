package cmds

type change struct {
	field    string
	oldValue string
	newValue string
}

func printDiff(changes []change) { _ = "STUB: not implemented"; return }

func printMapDiff(title string, changes []change) { _ = "STUB: not implemented"; return }

func diffMaps(oldMap, newMap map[string]string) []change { _ = "STUB: not implemented"; return nil }

// Check for new and changed keys
