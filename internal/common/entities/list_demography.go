package entities

// ListDemography defines model for listDemography.
type ListDemography []struct {
	Data *[]struct {
		Employment *EmploymentInfo `json:"employment,omitempty"`
		Household  *HouseholdInfo  `json:"household,omitempty"`
		Lodgment   *LodgmentInfo   `json:"lodgment,omitempty"`
		Population *PopulationInfo `json:"population,omitempty"`
		Year       *string         `json:"year,omitempty"`
	} `json:"data,omitempty"`

	InseeCode      *string `json:"insee_code,omitempty"`
	SectorName     *string `json:"sector_name,omitempty"`
	YearLastUpdate *int64  `json:"year_last_update,omitempty"`
}

// Répartition par age.
type AgeDivision struct {
	// Population pour la tranche d'âge 0-14 ans.
	Age014 *int64 `json:"age_0_14,omitempty"`

	// Population pour la tranche d'âge 15-29 ans.
	Age1529 *int64 `json:"age_15_29,omitempty"`

	// Population pour la tranche d'âge 30-44 ans.
	Age3044 *int64 `json:"age_30_44,omitempty"`

	// Population pour la tranche d'âge 45-59 ans.
	Age4559 *int64 `json:"age_45_59,omitempty"`

	// Population pour la tranche d'âge 60-74 ans.
	Age6074 *int64 `json:"age_60_74,omitempty"`

	// Population pour la tranche d'âge 75-89 ans.
	Age7589 *int64 `json:"age_75_89,omitempty"`

	// Population pour la tranche d'âge 90 ans et plus.
	Age90More *int64 `json:"age_90_more,omitempty"`
}

// BuildingDivision defines model for building-division.
type BuildingDivision struct {
	// Nombre d'appartement.
	Appartment *int64 `json:"appartment,omitempty"`

	// Nombre de maison.
	House *int64 `json:"house,omitempty"`

	// Total.
	Total *int64 `json:"total,omitempty"`
}

// Indicateurs du logement des ménages.
type DivisionHousehold struct {
	// Division ménages.
	Division *HouseholdDivision `json:"division,omitempty"`
}

// Indicateurs d'activité économique.
type EmploymentInfo struct {
	// Taux de personnes salariées.
	EmployeeRate *float64 `json:"employee_rate,omitempty"`

	// Taux d'évolution de l'emploi depuis la dernière année d'analyse.
	EmploymentEvolutionRate *float64 `json:"employment_evolution_rate,omitempty"`
	Psc                     *struct {
		// Taux employés
		EmployeeRate *float64 `json:"employee_rate,omitempty"`

		// Taux artisans, commerçants, chefs entreprise
		EntrepreneurRate *float64 `json:"entrepreneur_rate,omitempty"`

		// Taux agriculteurs exploitants
		FarmerRate *float64 `json:"farmer_rate,omitempty"`

		// Taux professions intermédiaires
		IntermediaiteProfessionRate *float64 `json:"intermediaite_profession_rate,omitempty"`

		// Taux retraités
		RetiredRate *float64 `json:"retired_rate,omitempty"`

		// Taux cadres et professions intellectuelles supérieures
		SeniorExecutiveRate *float64 `json:"senior_executive_rate,omitempty"`

		// Taux chomeurs
		WithoutActivityRate *float64 `json:"without_activity_rate,omitempty"`

		// Taux ouvriers
		WorkerRate *float64 `json:"worker_rate,omitempty"`
	} `json:"psc,omitempty"`

	// Total travailleurs.
	TotalWorkers *int64 `json:"total_workers,omitempty"`

	// Taux d'évolution du chômage depuis la dernière année d'analyse.
	UnemploymentEvolutionRate *float64 `json:"unemployment_evolution_rate,omitempty"`

	// Taux de chômage.
	UnemploymentRate *float64 `json:"unemployment_rate,omitempty"`
}

// Division des familles.
type FamilyDivision struct {
	// Autre.
	Other *int64 `json:"other,omitempty"`

	// Parent seul.
	SingleParent *int64 `json:"single_parent,omitempty"`

	// Avec enfant.
	WithChild *int64 `json:"with_child,omitempty"`

	// Sans enfant.
	WithoutChild *int64 `json:"without_child,omitempty"`
}

// Division ménages.
type HouseholdDivision struct {
	// Nombre de famille.
	Family *int64 `json:"family,omitempty"`

	// Division des familles.
	FamilyDivision *FamilyDivision `json:"family_division,omitempty"`

	// Nombre de personne seule.
	SinglePerson *int64 `json:"single_person,omitempty"`

	// Division des personnes seules.
	SinglePersonDivision *SinglePersonDivision `json:"single_person_division,omitempty"`
}

// Indicateurs des ménages.
type HouseholdInfo struct {
	// Evolution des emménagements.
	LastYearRotation *LastYearRotation `json:"last_year_rotation,omitempty"`

	// Ancienneté d'emménagement dans la résidence principale.
	MoveInAge *MoveInAge `json:"move_in_age,omitempty"`
}

// Evolution des emménagements.
type LastYearRotation struct {
	// Population habitant auparavant dans un autre logement de la même agglomeration.
	ChangeCitySameAgglomeration *int64 `json:"change_city_same_agglomeration,omitempty"`

	// Population habitant auparavant dans un autre logement dans le même département.
	ChangeCitySameDepartment *int64 `json:"change_city_same_department,omitempty"`

	// Population habitant auparavant dans la même région dans un autre département.
	ChangeDepartementSameRegion *int64 `json:"change_departement_same_region,omitempty"`

	// Population habitant auparavant dans un autre logement de la même commune.
	ChangeLodgmentSameCity *int64 `json:"change_lodgment_same_city,omitempty"`

	// Population habitant auparavant à l'étrangé.
	ChangeNoMainland *int64 `json:"change_no_mainland,omitempty"`

	// Population habitant auparavant dans une région sur un autre continent.
	ChangeRegionMainland *int64 `json:"change_region_mainland,omitempty"`

	// Population habitant auparavant dans le même logement.
	SameLodgment *int64 `json:"same_lodgment,omitempty"`
}

// Division des logements.
type LodgmentDivision struct {
	// Taux d'appartement.
	ApartmentRate *float64 `json:"apartment_rate,omitempty"`

	// Taux de résidences vides.
	EmptyLodgementRate *float64 `json:"empty_lodgement_rate,omitempty"`

	// Taux de maison.
	HouseRate *float64 `json:"house_rate,omitempty"`

	// Indicateurs du logement des ménages.
	Household *DivisionHousehold `json:"household,omitempty"`

	// Division des résidences principales par ancièneté de contruction et par typologie.
	MainLodgmentBuildingTimeDivision *MainLodgmentBuildingTimeDivision `json:"main_lodgment_building_time_division,omitempty"`

	// Taux de résidences principales.
	MainLodgmentRate *float64 `json:"main_lodgment_rate,omitempty"`

	// Division des résidences principales par nombre de pièce.
	MainLodgmentRoomsDivision *MainLodgmentRoomsDivision `json:"main_lodgment_rooms_division,omitempty"`

	// Taux de résidences secondaires.
	SecondaryLodgementRate *float64 `json:"secondary_lodgement_rate,omitempty"`

	// Number of properties.
	TotalLodgments *int64 `json:"total_lodgments,omitempty"`
}

// Indicateurs sur le logement des ménages.
type LodgmentHousehold struct {
	// Revenu moyen des ménages.
	AverageSalaryHousehold *int64 `json:"average_salary_household,omitempty"`

	// Taux de ménages imposables.
	RateTaxableHousehold *float64 `json:"rate_taxable_household,omitempty"`

	// Nombre de ménages imposés.
	TotalTaxableHousehold *int64 `json:"total_taxable_household,omitempty"`
}

// Indicateurs d'occupation des logements.
type LodgmentInfo struct {
	// Division des logements.
	Division *LodgmentDivision `json:"division,omitempty"`

	// Indicateurs sur le logement des ménages.
	Household *LodgmentHousehold `json:"household,omitempty"`

	// Occupation des logements.
	Occupation *LodgmentOccupation `json:"occupation,omitempty"`
}

// Occupation des logements.
type LodgmentOccupation struct {
	// Taux de locataires hébergés gratuitement.
	FreeTenantRate *float64 `json:"free_tenant_rate,omitempty"`

	// Taux de propriétaires.
	OwnerRate *float64 `json:"owner_rate,omitempty"`

	// Taux de locataires socialement aidés.
	SocialTenantRate *float64 `json:"social_tenant_rate,omitempty"`

	// Taux de locataires.
	TenantRate *float64 `json:"tenant_rate,omitempty"`
}

// Division des résidences principales par ancièneté de contruction et par typologie.
type MainLodgmentBuildingTimeDivision struct {
	Before1919     *BuildingDivision `json:"before_1919,omitempty"`
	From1919To1945 *BuildingDivision `json:"from_1919_to_1945,omitempty"`
	From1946To1970 *BuildingDivision `json:"from_1946_to_1970,omitempty"`
	From1971To1990 *BuildingDivision `json:"from_1971_to_1990,omitempty"`
	From1991To2005 *BuildingDivision `json:"from_1991_to_2005,omitempty"`
	From2006To2015 *BuildingDivision `json:"from_2006_to_2015,omitempty"`
}

// Division des résidences principales par nombre de pièce.
type MainLodgmentRoomsDivision struct {
	// Total T1.
	T1 *int64 `json:"t1,omitempty"`

	// Total T2.
	T2 *int64 `json:"t2,omitempty"`

	// Total T3.
	T3 *int64 `json:"t3,omitempty"`

	// Total T4.
	T4 *int64 `json:"t4,omitempty"`

	// Total T5.
	T5 *int64 `json:"t5,omitempty"`
}

// Ancienneté d'emménagement dans la résidence principale.
type MoveInAge struct {
	// Depuis moins de 2 ans.
	Age02 *int64 `json:"age_0_2,omitempty"`

	// 10 ans ou plus.
	Age10More *int64 `json:"age_10_more,omitempty"`

	// De 2 à 4 ans.
	Age24 *int64 `json:"age_2_4,omitempty"`

	// De 5 à 9 ans.
	Age59 *int64 `json:"age_5_9,omitempty"`
}

// Indicateurs de répartition de la population.
type PopulationInfo struct {
	// Répartition par age.
	AgeDivision *AgeDivision `json:"age_division,omitempty"`

	// Taux d'évolution de la population depuis la dernière année d'analyse.
	EvolutionPopulationRate *float64 `json:"evolution_population_rate,omitempty"`

	// Total population.
	Population *int64 `json:"population,omitempty"`
}

// Division des personnes seules.
type SinglePersonDivision struct {
	// Homme.
	Man *int64 `json:"man,omitempty"`

	// Femme.
	Woman *int64 `json:"woman,omitempty"`
}
