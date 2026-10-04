package calculator

type RiskProfileInput struct {
	MonthlyExpense    float64 `json:"monthly_expense"`
	IsMarried         bool    `json:"is_married"`
	NumberOfChildren  int     `json:"number_of_children"`
	ExistingInsurance float64 `json:"existing_insurance"`
	AnnualIncome      float64 `json:"annual_income"`
}

type RiskResult struct {
	IdealEmergencyFund float64 `json:"ideal_emergency_fund"`
	HumanLifeValue     float64 `json:"human_life_value"`
	InsuranceGap       float64 `json:"insurance_gap"`
}

func CalculateRiskAndEmergencyFund(input RiskProfileInput) RiskResult {
	// Perhitungan Dana Darurat (Standar CFP)
	multiplier := 4.0
	if input.IsMarried {
		multiplier = 6.0
	}
	multiplier += float64(input.NumberOfChildren * 2)
	idealEmergencyFund := input.MonthlyExpense * multiplier

	// Perhitungan Human Life Value sederhana (Standar AWP)
	humanLifeValue := input.AnnualIncome * 10.0

	// Perhitungan Insurance Gap
	insuranceGap := humanLifeValue - input.ExistingInsurance
	if insuranceGap < 0 {
		insuranceGap = 0
	}

	return RiskResult{
		IdealEmergencyFund: idealEmergencyFund,
		HumanLifeValue:     humanLifeValue,
		InsuranceGap:       insuranceGap,
	}
}
