package meta

func (m *Meta) ModifiedSection(section Section) {
	m.modifiedSections |= byte(section)
}

func (m *Meta) WasModified(section Section) bool {
	return m.modifiedSections&byte(section) != 0
}

func (m *Meta) HasModifiedSections() bool {
	return m.modifiedSections != 0
}
