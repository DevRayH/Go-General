package child

import (
	parent "first_go/StructPractice/Struct1"
)

type Child struct{
	parent.Parent
	Gender string
}