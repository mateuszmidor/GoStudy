package sounds

import "strings"

// Sound represents a transcripted sound
type Sound string

const (
	// Silence represents no sound :)
	Silence Sound = ""

	// PipeBlast represents sound of exhaust explosion
	PipeBlast Sound = "Pipe Blast!"
)

// Sounds is a list of sounds
type Sounds []Sound

// Contains checks for existence
func (sounds Sounds) Contains(sound Sound) bool {
	for _, s := range sounds {
		if s == sound {
			return true
		}
	}
	return false
}

// Append appends :)
func (sounds Sounds) Append(sound Sound) Sounds {
	return append(sounds, sound)
}

func (sounds Sounds) String() string {
	var sb strings.Builder
	for _, s := range sounds {
		sb.WriteString(string(s))
		sb.WriteString(",")
	}

	if sb.Len() == 0 {
		return string(Silence)
	}

	s := sb.String()
	return s[0 : len(s)-1] // cut last comma
}
