package dashboard

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
)

func (u *usecase) QuizAll(req request.ParamDashboard) ([]*response.QuizAllDashboardPerTest, error) {
	user, err := u.userRepo.FindOneBy(map[string]interface{}{
		"id":      req.AuthenticatedUser,
		"role_id": helper.PembeliPaketBimbel,
	})
	if user == nil {
		return nil, helper.ErrAccessDenied
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}

	//pretest data
	testTypePretest, err := u.testTypeRepo.FindOneBy(map[string]interface{}{
		"name": "Pretest",
	})

	if err != nil {
		return nil, fmt.Errorf("failed to find testTypePretest: %v", err)
	}

	testTypePosttest, err := u.testTypeRepo.FindOneBy(map[string]interface{}{
		"name": "Posttest",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find testTypePosttest: %v", err)
	}

	moduleSKD, err := u.moduleRepo.FindOneBy(map[string]interface{}{
		"name": "SKD",
	})

	if err != nil {
		return nil, fmt.Errorf("failed to find moduleSKD: %v", err)
	}
	moduleMatematika, err := u.moduleRepo.FindOneBy(map[string]interface{}{
		"name": "Matematika",
	})

	if err != nil {
		return nil, fmt.Errorf("failed to find moduleMatematika: %v", err)
	}

	pretestQuizAllSKD, err := u.dashboardLocalRepo.GetQuizAllDashboard(req.AuthenticatedUser, testTypePretest.ID, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to find pretestQuizAllSKD: %v", err)
	}

	pretestQuizAllMatematika, err := u.dashboardLocalRepo.GetQuizAllDashboard(req.AuthenticatedUser, testTypePretest.ID, 2)
	if err != nil {
		return nil, fmt.Errorf("failed to find pretestQuizAllMatematika: %v", err)
	}

	posttestQuizAllSKD, err := u.dashboardLocalRepo.GetQuizAllDashboard(req.AuthenticatedUser, testTypePosttest.ID, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to find posttestQuizAllSKD: %v", err)
	}

	posttestQuizAllMatematika, err := u.dashboardLocalRepo.GetQuizAllDashboard(req.AuthenticatedUser, testTypePosttest.ID, 2)
	if err != nil {
		return nil, fmt.Errorf("failed to find posttestQuizAllMatematika: %v", err)
	}

	pretestQuizAllSKDPerModul := &response.QuizAllDashboardPerModule{
		UUID:       moduleSKD.UUID,
		ModuleName: moduleSKD.Name,
		Quizzes:    pretestQuizAllSKD,
	}

	pretestQuizAllMatematikaPerModul := &response.QuizAllDashboardPerModule{
		UUID:       moduleMatematika.UUID,
		ModuleName: moduleMatematika.Name,
		Quizzes:    pretestQuizAllMatematika,
	}

	posttestQuizAllSKDPerModul := &response.QuizAllDashboardPerModule{
		UUID:       moduleSKD.UUID,
		ModuleName: moduleSKD.Name,
		Quizzes:    posttestQuizAllSKD,
	}

	posttestQuizAllMatematikaPerModul := &response.QuizAllDashboardPerModule{
		UUID:       moduleMatematika.UUID,
		ModuleName: moduleMatematika.Name,
		Quizzes:    posttestQuizAllMatematika,
	}

	var pretestQuizAllDashboardPerModul []*response.QuizAllDashboardPerModule
	pretestQuizAllDashboardPerModul = append(pretestQuizAllDashboardPerModul, pretestQuizAllSKDPerModul)
	pretestQuizAllDashboardPerModul = append(pretestQuizAllDashboardPerModul, pretestQuizAllMatematikaPerModul)

	var posttestQuizAllDashboardPerModul []*response.QuizAllDashboardPerModule
	posttestQuizAllDashboardPerModul = append(posttestQuizAllDashboardPerModul, posttestQuizAllSKDPerModul)
	posttestQuizAllDashboardPerModul = append(posttestQuizAllDashboardPerModul, posttestQuizAllMatematikaPerModul)

	pretest := &response.QuizAllDashboardPerTest{
		TestTypeName:              testTypePretest.Name,
		QuizAllDashboardPerModule: pretestQuizAllDashboardPerModul,
	}

	posttest := &response.QuizAllDashboardPerTest{
		TestTypeName:              testTypePosttest.Name,
		QuizAllDashboardPerModule: posttestQuizAllDashboardPerModul,
	}

	var quizAllDashboardPerTest []*response.QuizAllDashboardPerTest

	quizAllDashboardPerTest = append(quizAllDashboardPerTest, pretest)
	quizAllDashboardPerTest = append(quizAllDashboardPerTest, posttest)

	result := quizAllDashboardPerTest
	return result, nil

}
