from dataclasses import dataclass, field
from typing import List


@dataclass
class EducationalInstitutionDTO:
    Name: str

    def to_dict(self):
        return {
            "name": self.Name
        }


@dataclass
class FacultyDTO:
    Id: int
    Name: str

    def to_dict(self):
        return {
            "id": self.Id,
            "name": self.Name
        }


@dataclass
class YearDTO:
    Id: int
    Name: str

    def to_dict(self):
        return {
            "id": self.Id,
            "name": self.Name
        }


@dataclass
class GroupDTO:
    Id: int
    Name: str

    def to_dict(self):
        return {
            "id": self.Id,
            "name": self.Name
        }


@dataclass
class LessonDTO:
    StartTime: str
    EndTime: str
    Name: str
    Teacher: str
    Auditorium: str
    Type: str
    SubGroup: int = None

    def to_dict(self):
        lesson = {
            "name": self.Name,
            "teacher": self.Teacher,
            "auditorium": self.Auditorium,
            "type": self.Type,
            "startTime": self.StartTime,
            "endTime": self.EndTime
        }

        if self.SubGroup != None:
            lesson["subGroup"] = str(self.SubGroup)

        return lesson


@dataclass
class DayDTO:
    Number: int
    Lessons: List[LessonDTO] = field(default_factory=list)

    def to_dict(self):
        return {
            "number": self.Number,
            "lessons": [lesson.to_dict() for lesson in self.Lessons],
        }


@dataclass
class ScheduleDTO:
    EducationalInstitution: EducationalInstitutionDTO
    Faculty: FacultyDTO = FacultyDTO
    Year: YearDTO = YearDTO
    Group: GroupDTO = GroupDTO
    EvenWeek: List[DayDTO] = field(default_factory=list)
    OddWeek: List[DayDTO] = field(default_factory=list)

    def to_dict(self):
        return {
            "educationalInstitution": self.EducationalInstitution.to_dict(),
            "faculty": self.Faculty.to_dict(),
            "year": self.Year.to_dict(),
            "group": self.Group.to_dict(),
            "evenWeek": [week.to_dict() for week in self.EvenWeek],
            "oddWeek": [week.to_dict() for week in self.EvenWeek],
        }
