package meta

func (m *Meta) LockSection(section Section) {
	m.unlockedSections |= byte(section)
}

func (m *Meta) UnlockSection(section Section) {
	m.unlockedSections &^= byte(section)
}

func (m *Meta) IsLocked(section Section) bool {
	return m.unlockedSections&byte(section) == 0
}

func (m *Meta) HasUnlockedSections() bool {
	return m.unlockedSections != 0
}
