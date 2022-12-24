package interfaces

import "engine/internal/core/entity"

type RepositoryManager interface {
	BeginTran() (RepositoryManager, error)
	CommitTran() error
	RollbackTran() error
	Transaction(callback func(RepositoryManager) error) error

	GetEducationalRepository() EducationalRepository
	GetGroupRepository() GroupRepository
	GetDayRepository() DayRepository
	GetLessonRepository() LessonRepository
	GetStudentRepository() StudentRepository
}

type EducationalRepository interface {
	SaveEducationalInstitution(name string) (entity.EducationalInstitution, error)
	SaveFaculty(name string) (entity.Faculty, error)
	SaveYear(name string) (entity.Year, error)
}

type GroupRepository interface {
	Save(*entity.Group) (*entity.Group, error)
	Get(int) (*entity.Group, error)
	GetByName(string) (*entity.Group, error)
}

type LessonRepository interface {
	Save(*entity.Lesson) (*entity.Lesson, error)
}

type DayRepository interface {
	Save(weekType string, number, groupId int) (*entity.Day, error)
	GetLessons(weekType string, weekDay entity.Weekday, groupId int) (*entity.Day, error)
}

type StudentRepository interface {
	Save(*entity.Student) (*entity.Student, error)
	Get(int) (*entity.Student, error)
	GetList() ([]entity.Student, error)
}
