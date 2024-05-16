package mapper

import (
	"fmt"

	"github.com/jpmoraess/service-scheduling/internal/domain/entity"
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
	"github.com/jpmoraess/service-scheduling/internal/infra/persistence/data"
)

func toBreak(b *data.BreakData) (*vo.Break, error) {
	if b == nil {
		return nil, nil
	}
	return vo.NewBreak(b.StartTime, b.EndTime)
}

func toBreakData(aBreak *vo.Break) (*data.BreakData, error) {
	return &data.BreakData{
		StartTime: aBreak.StartTime(),
		EndTime:   aBreak.EndTime(),
	}, nil
}

func toDay(d *data.DayData) (*vo.Day, error) {
	aBreak, err := toBreak(d.Break)
	if err != nil {
		return nil, err
	}
	day, err := vo.NewDay(d.StartTime, d.EndTime, aBreak)
	if err != nil {
		return nil, err
	}
	return day, nil
}

// func toDayData(day *vo.Day) (*data.DayData, error) {
// 	var breakData = &data.BreakData{}
// 	if day.Break() == nil {
// 		breakData = nil
// 	} else {
// 		b, err := toBreakData(day.Break())
// 		if err != nil {
// 			return nil, err
// 		}
// 		breakData = b
// 	}
// 	return &data.DayData{
// 		StartTime: day.StartTime(),
// 		EndTime:   day.EndTime(),
// 		Break:     breakData,
// 	}, nil
// }

func toDayData(day *vo.Day) (*data.DayData, error) {
	if day == nil {
		return nil, nil
	}
	breakVO := day.Break()
	var breakData *data.BreakData
	if breakVO != nil {
		b, err := toBreakData(breakVO)
		if err != nil {
			return nil, err
		}
		breakData = b
	}
	return &data.DayData{
		StartTime: day.StartTime(),
		EndTime:   day.EndTime(),
		Break:     breakData,
	}, nil
}

func toWorkPlan(w *data.WorkPlanData) (*vo.WorkPlan, error) {
	monday, err := toDay(w.Monday)
	if err != nil {
		return nil, err
	}
	tuesday, err := toDay(w.Tuesday)
	if err != nil {
		return nil, err
	}
	wednesday, err := toDay(w.Wednesday)
	if err != nil {
		return nil, err
	}
	thursday, err := toDay(w.Thursday)
	if err != nil {
		return nil, err
	}
	friday, err := toDay(w.Friday)
	if err != nil {
		return nil, err
	}
	saturday, err := toDay(w.Saturday)
	if err != nil {
		return nil, err
	}
	sunday, err := toDay(w.Sunday)
	if err != nil {
		return nil, err
	}
	workPlan, err := vo.NewWorkPlan(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
	if err != nil {
		return nil, err
	}
	return workPlan, nil
}

func toWorkPlanData(workPlan *vo.WorkPlan) (*data.WorkPlanData, error) {
	monday, err := toDayData(workPlan.Monday())
	if err != nil {
		return nil, err
	}
	tuesday, err := toDayData(workPlan.Tuesday())
	if err != nil {
		return nil, err
	}
	wednesday, err := toDayData(workPlan.Wednesday())
	if err != nil {
		return nil, err
	}
	thursday, err := toDayData(workPlan.Thursday())
	if err != nil {
		return nil, err
	}
	friday, err := toDayData(workPlan.Friday())
	if err != nil {
		return nil, err
	}
	saturday, err := toDayData(workPlan.Saturday())
	if err != nil {
		return nil, err
	}
	sunday, err := toDayData(workPlan.Sunday())
	if err != nil {
		return nil, err
	}

	return &data.WorkPlanData{
		Monday:    monday,
		Tuesday:   tuesday,
		Wednesday: wednesday,
		Thursday:  thursday,
		Friday:    friday,
		Saturday:  saturday,
		Sunday:    sunday,
	}, nil
}

func ToProfessionalData(professional *entity.Professional) (*data.ProfessionalData, error) {
	workplanData, err := toWorkPlanData(professional.WorkPlan())
	if err != nil {
		return nil, err
	}
	return &data.ProfessionalData{
		AccountID:       professional.AccountID(),
		EstablishmentID: professional.EstablishmentID(),
		Name:            professional.Name(),
		WorkPlan:        workplanData,
		Active:          professional.Active(),
		CreatedAt:       professional.CreatedAt(),
	}, nil
}

func FromProfessionalData(data *data.ProfessionalData) (*entity.Professional, error) {
	workPlan, err := toWorkPlan(data.WorkPlan)
	if err != nil {
		return nil, err
	}
	professional, err := entity.RestoreProfessional(data.ID.Hex(), data.AccountID, data.EstablishmentID, data.Name, workPlan, data.CreatedAt)
	if err != nil {
		fmt.Println("error to restore professional from database", err)
		return nil, err
	}
	return professional, nil
}
