package models

type FollowUpData struct {
	VitalStatus    string          `json:"vital_status"`
	DeathDetails   *DeathDetails   `json:"death_details,omitempty"`
	DateOfFollowUp string          `json:"date_of_follow_up,omitempty"`
	Symptoms       *Symptoms       `json:"symptoms"`
	Investigations *Investigations `json:"investigations"`
	Treatment      *Treatment      `json:"treatment"`
}

type DeathDetails struct {
	DateOfDeath  string        `json:"date_of_death"`
	CauseOfDeath string        `json:"cause_of_death"`
	CardiacDeath *CardiacDeath `json:"cardiac_death,omitempty"`
}

type CardiacDeath struct {
	Cause string `json:"cause"`
}

type Symptoms struct {
	ChestPain       string   `json:"chest_pain"`
	Breathlessness  string   `json:"breathlessness"`
	Stroke          *History `json:"stroke"`
	Reinfarction    *History `json:"reinfarction"`
	Hospitalization *History `json:"hospitalization"`
	PCI             *History `json:"pci"`
	CABG            *History `json:"cabg"`
}

type Biochemistry struct {
	FastingSugar  string         `json:"fasting_sugar"`
	HbA1C         string         `json:"hba1c"`
	LipidProfile  *LipidProfile  `json:"lipid_profile"`
	RenalFunction *RenalFunction `json:"renal_function"`
}

type LipidProfile struct {
	Status           string `json:"status"`
	TotalCholesterol string `json:"total_cholesterol"`
	LDLC             string `json:"ldlc"`
	HDLC             string `json:"hdlc"`
	TG               string `json:"tg"`
}

type RenalFunction struct {
	IsDone     bool   `json:"is_done"`
	Creatinine string `json:"creatinine,omitempty"`
	EGFR       string `json:"e_gfr,omitempty"`
	Na         string `json:"na,omitempty"`
	K          string `json:"k,omitempty"`
	UricAcid   string `json:"uric_acid,omitempty"`
}

type ECG struct {
	IsDone  bool        `json:"is_done"`
	Details *ECGDetails `json:"details,omitempty"`
}

type ECGDetails struct {
	Rhythm           string `json:"rhythm"`
	Conduction       string `json:"conduction"`
	STAndTWaveChange bool   `json:"st_and_t_wave_change"`
	PathologicalQ    bool   `json:"pathological_q"`
}

type Echocardiography struct {
	IsDone  bool                     `json:"is_done"`
	Details *EchocardiographyDetails `json:"details,omitempty"`
}

type EchocardiographyDetails struct {
	Result              string `json:"result"`
	LVFunction          string `json:"lv_function"`
	MitralRegurgitation string `json:"mitral_regurgitation"`
	LVThrombus          bool   `json:"lv_thrombus"`
	LVEDVolume          string `json:"lved_volume"`
	LVESV               string `json:"lvesv"`
}

type StressTest struct {
	IsDone              bool              `json:"is_done"`
	StressECGDetails    *StressECGDetails `json:"stress_ecg_details,omitempty"`
	StressEcho          string            `json:"stress_echo,omitempty"`
	StressPerfusionScan string            `json:"stress_perfusion_scan,omitempty"`
}

type StressECGDetails struct {
	TestType                string `json:"test_type"`
	DurationOfExercise      string `json:"duration_of_exercise"`
	STSegmentDepression     string `json:"st_segment_depression"`
	ChestPainDuringExercise bool   `json:"chest_pain_during_exercise"`
}

type HolterStudy struct {
	IsDone  bool                `json:"is_done"`
	Details *HolterStudyDetails `json:"details,omitempty"`
}

type HolterStudyDetails struct {
	Result string `json:"result"`
}

type CoronaryAngiography struct {
	IsDone  bool                        `json:"is_done"`
	Details *CoronaryAngiographyDetails `json:"details,omitempty"`
}

type CoronaryAngiographyDetails struct {
	Indication    string `json:"indication"`
	ExtentOfCAD   string `json:"extent_of_cad"`
	LMCAD         bool   `json:"lm_cad"`
	SeverityOfCAD string `json:"severity_of_cad"`
}

type Investigations struct {
	Biochemistry        *Biochemistry        `json:"biochemistry"`
	ECG                 *ECG                 `json:"ecg"`
	Echocardiography    *Echocardiography    `json:"echocardiography"`
	StressTest          *StressTest          `json:"stress_test"`
	HolterStudy         *HolterStudy         `json:"holter_study"`
	CoronaryAngiography *CoronaryAngiography `json:"coronary_angiography"`
}

type Treatment struct {
	PCI               *History           `json:"pci"`
	CABG              *History           `json:"cabg"`
	DrugTreatment     *DrugTreatment     `json:"drug_treatment"`
	DiabetesTreatment *DiabetesTreatment `json:"diabetes_treatment"`
	OtherDrugs        string             `json:"other_drugs"`
}

type History struct {
	IsDone     bool   `json:"is_done"`
	Date       string `json:"date,omitempty"`
	Indication string `json:"indication,omitempty"`
}

type DrugTreatment struct {
	AntiAnginals       *Details               `json:"anti_anginals"`
	AntiPlatelets      *Details               `json:"anti_platelets"`
	Anticoagulants     *Details               `json:"anticoagulants"`
	LipidLoweringDrugs *LipidLoweringDrugs    `json:"lipid_lowering_drugs"`
	BetaBlockers       *Details               `json:"beta_blockers"`
	RAASInhibitors     *RAASInhibitorsDetails `json:"raas_inhibitors"`
	MRA                string                 `json:"mra"`
	SGLT2Inhibitors    bool                   `json:"sglt2_inhibitors"`
}

type RAASInhibitorsDetails struct {
	IsPresent bool       `json:"is_present"`
	Inhibitor *Inhibitor `json:"inhibitor,omitempty"`
	ARNI      bool       `json:"arni"`
}

type Inhibitor struct {
	Type    string   `json:"type"`
	Salt    string   `json:"salt"`
	Details *Details `json:"details,omitempty"`
}

type DiabetesTreatment struct {
	Metformin      *Details `json:"metformin"`
	DPP4Inhibitors *Details `json:"dpp4_inhibitors"`
	GLP1Agonist    *Details `json:"glp1_agonist"`
	Sulphonylurea  *Details `json:"sulphonylurea"`
	Insulin        *Details `json:"insulin"`
}

type LipidLoweringDrugs struct {
	Statins       *Details `json:"statins"`
	Ezetimibe     bool     `json:"ezetimibe"`
	BempedoicAcid bool     `json:"bempedoic_acid"`
	Fenofibrate   bool     `json:"fenofibrate"`
}
type Details struct {
	IsPresent bool   `json:"is_present"`
	Type      string `json:"type,omitempty"`
	Dose      string `json:"dose,omitempty"`
	Frequency string `json:"frequency,omitempty"`
}
