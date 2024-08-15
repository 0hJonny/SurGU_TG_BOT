package models

import (
	"encoding/json"
	"telegram_bot/src/logger"
)

type Response struct {
	Group              Group         `json:"group"`
	WithJobOfferNumber bool          `json:"with_job_offer_number"`
	List               []UserProfile `json:"list"`
}

type Group struct {
	ID                             int                `json:"id"`
	Name                           string             `json:"name"`
	ID1c                           string             `json:"id_1c"`
	SpecialityID                   int                `json:"speciality_id"`
	BudgetTypeID                   int                `json:"budget_type_id"`
	EducationFormatID              int                `json:"education_format_id"`
	EducationQuotaID               int                `json:"education_quota_id"`
	EducationConditionID           int                `json:"education_condition_id"`
	CreatedAt                      string             `json:"created_at"`
	UpdatedAt                      string             `json:"updated_at"`
	CurriculumID                   int                `json:"curriculum_id"`
	CampaignID                     int                `json:"campaign_id"`
	CountOfPlaces                  int                `json:"count_of_places"`
	EducationProfileID             int                `json:"education_profile_id"`
	StartAt                        string             `json:"start_at"`
	EndInternalExamsAt             string             `json:"end_internal_exams_at"`
	EndEgExamsAt                   string             `json:"end_ege_exams_at"`
	EduDocumentDeadlineAt          string             `json:"edu_document_deadline_at"`
	OrderAt                        string             `json:"order_at"`
	ForbiddenSpoInternal           int                `json:"forbidden_spo_internal"`
	IsHealthCertificateRequirement int                `json:"is_health_certificate_requirement"`
	ExamsStartAt                   string             `json:"exams_start_at"`
	ExamsEndAt                     string             `json:"exams_end_at"`
	FormatName                     string             `json:"format_name"`
	EducationCondition             EducationCondition `json:"education_condition"`
	EducationQuota                 EducationQuota     `json:"education_quota"`
	EducationFormat                EducationFormat    `json:"education_format"`
}

type EducationCondition struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ID1c        string `json:"id_1c"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	LandingName string `json:"landing_name"`
}

type EducationQuota struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ID1c           string `json:"id_1c"`
	IsSpecialRight int    `json:"is_special_right"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type EducationFormat struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	ID1c            string `json:"id_1c"`
	LandingName     string `json:"landing_name"`
	JobInRussiaName string `json:"job_in_russia_name"`
}

type UserProfile struct {
	ID                    int     `json:"id"`
	CompetitiveGroupID    int     `json:"competitive_group_id"`
	ID1c                  string  `json:"id_1c"`
	Snils                 string  `json:"snils"`
	Priority              int     `json:"priority"`
	Delivery              string  `json:"delivery"`
	Original              int     `json:"original"`
	OriginalEpgu          int     `json:"original_epgu"`
	ScoresSum             int     `json:"scores_sum"`
	ScoresSubjectsSum     int     `json:"scores_subjects_sum"`
	ScoresAchievementsSum int     `json:"scores_achievements_sum"`
	ToOrder               int     `json:"to_order"`
	CreatedAt             string  `json:"created_at"`
	UpdatedAt             string  `json:"updated_at"`
	Scores                []Score `json:"scores"`
	IsOk                  int     `json:"is_ok"`
	IsOrdered             int     `json:"is_ordered"`
	JobOfferNumber        *string `json:"job_offer_number"`
	OriginalIs            int     `json:"original_is"`
	Identity              string  `json:"identity"`
	Status                Status  `json:"status"`
}

func OrderUserProfiles(u []UserProfile) []UserProfile {
	var result []UserProfile
	originalFirst := make([]UserProfile, 0, len(u))
	originalOther := make([]UserProfile, 0, len(u))
	for _, user := range u {
		if user.Original == 1 || user.OriginalEpgu == 1 && user.ToOrder == 1 {
			originalFirst = append(originalFirst, user)
		} else {
			originalOther = append(originalOther, user)
		}
	}

	result = append(originalFirst, originalOther...)

	return result
}

type Status struct {
	Label string `json:"label"`
	Color string `json:"color"`
	Help  string `json:"help"`
}

type Score struct {
	Score        string `json:"score"`
	Discipline   string `json:"discipline"`
	CampaignID   string `json:"campaign_id"`
	DisciplineID string `json:"discipline_id"`
}

type ResponseGroup []Response

func (r *ResponseGroup) Parse(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *ResponseGroup) GetBytes() []byte {
	var log = logger.Logger{}
	data, err := json.Marshal(r)
	if err != nil {
		log.Panic("Failed to get bytes: " + err.Error())
	}
	return data
}
