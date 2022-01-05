package placeholder

import (
    "fmt"
    "platform/logging"
)

var names = []string{"Alice", "Bob", "Charlie", "Dora"}

type NameHandler struct {
	logging.Logger
}

func (n NameHandler) GetName(i int) string {
    n.Logger.Debugf("GetName method invoked with argument: %v", i)
    if (i < len(names)) {
        return fmt.Sprintf("Name #%v: %v", i, names[i])
    } else {
        return fmt.Sprintf("Index out of bounds")
    }
}

func (n NameHandler) GetNames() string {
    n.Logger.Debug("GetNames method invoked")
    return fmt.Sprintf("Names: %v", names)
}

type NewName struct {
    Name string
    InsertAtStart bool
}

func (n NameHandler) PostName(new NewName) string {
    n.Logger.Debugf("PostName method invoked with argument %v", new)
    if (new.InsertAtStart) {
        names = append([] string { new.Name}, names... )
    } else {
        names = append(names, new.Name)
    }
    return fmt.Sprintf("Names: %v", names)
}
