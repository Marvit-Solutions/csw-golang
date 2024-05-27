package quiz

import (
	"fmt"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) QuizScoreAll(req request.ParamQuizScoreAll) (*response.QuizScoreAllResponse, error) {

	subModul, err := u.subModulRepo.FindOneBy(map[string]interface{}{
		"uuid": req.SubModulUUID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to find sub modul: %v", err)
	}

	subjects, err := u.subjectRepo.FindBy(map[string]interface{}{
		"sub_module_id": subModul.ID,
	}, 0, 0)

	if err != nil {
		return nil, fmt.Errorf("failed to find subjects: %v", err)
	}

	// create array of id subjects
	subjectsId := make([]int, len(subjects))

	for i, subject := range subjects {
		subjectsId[i] = subject.ID
	}

	// create map of subject id to subjects
	subjectsMap := make(map[int]*model.Subject)
	for _, subject := range subjects {
		subjectsMap[subject.ID] = subject
	}

	quizzes, err := u.quizRepo.FindBy(map[string]interface{}{
		"subject_id": subjectsId,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find quiz: %v", err)
	}

	quizzesQuestionTotalMap := make(map[int]int)

	// create map quiz id to quizzes
	quizzesMap := make(map[int]*model.Quiz)
	for _, quiz := range quizzes {
		quizQuestionTotal := u.quizQuestionRepo.Count(map[string]interface{}{
			"quiz_id": quiz.ID,
		})
		quizzesQuestionTotalMap[quiz.ID] = quizQuestionTotal
		quizzesMap[quiz.ID] = quiz
	}

	// create array of id quizzes
	quizzesId := make([]int, len(quizzes))

	for i, quiz := range quizzes {
		quizzesId[i] = quiz.ID
	}

	quizSubmissions, err := u.quizSubmissionRepo.FindBy(map[string]interface{}{
		"quiz_id": quizzesId,
		"user_id": req.AuthenticatedUser,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find quiz: %v", err)
	}

	sort.Slice(quizSubmissions, func(i, j int) bool {
		return quizSubmissions[i].ID < quizSubmissions[j].ID
	})

	var quizScoreAllItem []response.QuizScoreAllItemTmp
	for _, quizSubmission := range quizSubmissions {
		quizScoreAllItem = append(quizScoreAllItem, response.QuizScoreAllItemTmp{
			SubjectID:          quizzesMap[quizSubmission.QuizID].SubjectID,
			Subject:            subjectsMap[quizzesMap[quizSubmission.QuizID].SubjectID].Name,
			QuizUUID:           quizzesMap[quizSubmission.QuizID].UUID,
			Quiz:               quizzesMap[quizSubmission.QuizID].Title,
			QuizSubmissionUUID: quizSubmission.UUID,
			Score:              quizSubmission.Score,
			MaxScore:           quizzesMap[quizSubmission.QuizID].MaxScore,
			TotalRightAnswers:  quizSubmission.RightAnswer,
			TotalQuestions:     quizzesQuestionTotalMap[quizSubmission.QuizID],
		})
	}

	// groupBy data by subject id, and transform it to response.QuizScoreAllItemGroupBySubject
	groupedItems := make(map[int]response.QuizScoreAllItemGroupBySubject)

	for _, item := range quizScoreAllItem {
		if group, exists := groupedItems[item.SubjectID]; exists {
			group.QuizScoresDetail = append(group.QuizScoresDetail, response.QuizScoreAllSubjectItem{
				QuizSubmissionUUID: item.QuizSubmissionUUID,
				QuizUUID:           item.QuizUUID,
				Quiz:               item.Quiz,
				Score:              item.Score,
				MaxScore:           item.MaxScore,
				TotalRightAnswers:  item.TotalRightAnswers,
				TotalQuestions:     item.TotalQuestions,
			})
			groupedItems[item.SubjectID] = group
		} else {
			groupedItems[item.SubjectID] = response.QuizScoreAllItemGroupBySubject{
				SubjectID: item.SubjectID,
				Subject:   item.Subject,
				QuizScoresDetail: []response.QuizScoreAllSubjectItem{
					{
						QuizSubmissionUUID: item.QuizSubmissionUUID,
						QuizUUID:           item.QuizUUID,
						Quiz:               item.Quiz,
						Score:              item.Score,
						MaxScore:           item.MaxScore,
						TotalRightAnswers:  item.TotalRightAnswers,
						TotalQuestions:     item.TotalQuestions,
					},
				},
			}
		}
	}

	// convert map to slice
	var quizScoreAllItemGroupedList []response.QuizScoreAllItemGroupBySubject
	for _, group := range groupedItems {
		quizScoreAllItemGroupedList = append(quizScoreAllItemGroupedList, group)
	}

	result := &response.QuizScoreAllResponse{
		ID:               subModul.ID,
		SubModul:         subModul.Name,
		QuizScoreAllItem: quizScoreAllItemGroupedList,
	}

	sort.Slice(result.QuizScoreAllItem, func(i, j int) bool {
		return result.QuizScoreAllItem[i].SubjectID < result.QuizScoreAllItem[j].SubjectID
	})

	return result, nil
}
