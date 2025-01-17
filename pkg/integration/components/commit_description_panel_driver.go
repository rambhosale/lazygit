package components

type CommitDescriptionPanelDriver struct {
	t *TestDriver
}

func (self *CommitDescriptionPanelDriver) getViewDriver() *ViewDriver {
	return self.t.Views().CommitDescription()
}

func (self *CommitDescriptionPanelDriver) Type(value string) *CommitDescriptionPanelDriver {
	self.t.typeContent(value)

	return self
}

func (self *CommitDescriptionPanelDriver) SwitchToSummary() *CommitMessagePanelDriver {
	self.getViewDriver().PressTab()
	return &CommitMessagePanelDriver{t: self.t}
}

func (self *CommitDescriptionPanelDriver) AddNewline() *CommitDescriptionPanelDriver {
	self.t.press(self.t.keys.Universal.Confirm)
	return self
}
