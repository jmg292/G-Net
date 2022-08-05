package meta

func (m *Meta) LockSection(section sectionFlag) {
	m.unlockedSections |= byte(section)
}

func (m *Meta) UnlockSection(section sectionFlag) {
	m.unlockedSections &^= byte(section)
}

func (m *Meta) IsLocked(section sectionFlag) bool {
	return m.unlockedSections&byte(section) == 0
}

func (m *Meta) HasUnlockedSections() bool {
	return m.unlockedSections != 0
}
