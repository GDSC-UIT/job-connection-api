package controllers

import (
	"sort"

	"github.com/GDSC-UIT/job-connection-api/server/database"
	"github.com/GDSC-UIT/job-connection-api/server/middlewares"
	"github.com/GDSC-UIT/job-connection-api/server/models"
	"github.com/gofiber/fiber/v2"
	pq "github.com/lib/pq"
	"github.com/masatana/go-textdistance"
)

type CreateJobBody struct {
	CompanyID   string        `json:"company_id"`
	Title       string        `json:"title"`
	Address     string        `json:"address"`
	Description string        `json:"description"`
	SkillIds    pq.Int64Array `gorm:"type:integer[]" json:"skill_ids"`
}

func CreateJob(c *fiber.Ctx) error {
	p := new(models.Job)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	database.DBInstance.Db.Create(p)
	return c.JSON(json{Data: p})
}
func GetJobs(c *fiber.Ctx) error {
	info, ok := c.Locals("info").(middlewares.UserInfo)

	if ok {
		accountType := new(models.AccountType)
		err := database.DBInstance.Db.First(&accountType, "user_id = ?", info.ID).Error
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		if accountType.Type == "user" {
			var experiences []models.Experience

			if err := database.DBInstance.Db.Find(&experiences, "user_id = ?", info.ID).Error; err == nil {
				return getRecommenedJobs(experiences, c)
			}
		}
	}

	model := database.DBInstance.Db.Preload("Company").Model(&models.Job{})
	page := pg.Response(model, c.Request(), &[]models.Job{})
	return c.JSON(page)

}
func GetJobByID(c *fiber.Ctx) error {
	id := c.Params("id")
	job := new(models.Job)
	database.DBInstance.Db.Where("id = ?", id).First(job)
	return c.JSON(json{Data: job})
}
func UpdateJob(c *fiber.Ctx) error {
	var update_job_info models.Job
	id := c.Params("id")
	if err := c.BodyParser(&update_job_info); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var current_job_info models.Job
	result := database.DBInstance.Db.First(&current_job_info, "id = ?", id)

	if result.RowsAffected != 0 && result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	current_job_info.Title = update_job_info.Title
	current_job_info.Address = update_job_info.Address
	current_job_info.Description = update_job_info.Description
	current_job_info.SkillIds = update_job_info.SkillIds

	result = database.DBInstance.Db.Save(&current_job_info)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	return c.JSON(json{
		Data:    current_job_info,
		Message: "Informations update successful",
	})
}

func getRecommenedJobs(experiences []models.Experience, c *fiber.Ctx) error {
	// page, _ := strconv.Atoi(c.Query("page"))
	// size := 10

	experience_titles := ""
	experience_descriptions := ""
	for _, experience := range experiences {
		experience_titles = experience_titles + experience.JobTitle + " "
		experience_descriptions = experience_descriptions + experience.Description + " "
	}

	var jobs []models.Job
	if err := database.DBInstance.Db.Preload("Company").Find(&jobs).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	cache := make(map[uint]float64)
	sort.SliceStable(jobs[:], func(i, j int) bool {
		return getPoint(cache, jobs, jobs[i], experience_titles, experience_descriptions) > getPoint(cache, jobs, jobs[j], experience_titles, experience_descriptions)
	})

	return c.JSON(map[string]interface{}{
		"items": jobs,
	})
}

func getPoint(cache map[uint]float64, jobs []models.Job, job models.Job, experience_titles, experience_descriptions string) float64 {
	point, ok := cache[job.ID]
	if !ok {
		var prob_desciption float64
		var prob_title float64
		prob_title = textdistance.JaroWinklerDistance(experience_titles, job.Title)
		prob_desciption = textdistance.JaroWinklerDistance(experience_descriptions, job.Description)
		point = (prob_title + prob_desciption) * 0.5
		cache[job.ID] = point
	}
	return point
}
