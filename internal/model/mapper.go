package model

func (p *CreatePartInput) ToModel() *Part {
	return &Part{
		Name:             p.Name,
		PartNumber:       p.PartNumber,
		ManufacturerID:   p.ManufacturerID,
		Description:      p.Description,
		TemperatureStage: p.TemperatureStage,
		Specifications:   p.Specifications,
		Images:           p.Images,
	}
}

func ApplyUpdatePartInput(part *Part, input UpdatePartInput) {
	if input.Name != nil {
		part.Name = *input.Name
	}
	if input.Description != nil {
		part.Description = input.Description
	}
	if input.TemperatureStage != nil {
		part.TemperatureStage = input.TemperatureStage
	}
	if input.Specifications != nil {
		part.Specifications = input.Specifications
	}
	if input.Images != nil {
		part.Images = input.Images
	}
}

func ApplyUpdateCategoryInput(c *Category, input UpdateCategoryInput) {
	if input.Name != nil {
		c.Name = *input.Name
	}
	if input.Description != nil {
		c.Description = input.Description
	}
}

func ApplyUpdateManufacturerInput(m *Manufacturer, input UpdateManufacturerInput) {
	if input.Name != nil {
		m.Name = *input.Name
	}
	if input.CountryOfOrigin != nil {
		m.CountryOfOrigin = input.CountryOfOrigin
	}
}
