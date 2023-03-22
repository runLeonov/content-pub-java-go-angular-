package service

import (
	app "contentPublicationBACK"
	"contentPublicationBACK/pkg/repository"
)

type TitleService struct {
	repoTitles     repository.Titles
	repoCategories repository.Categories
}

func NewTitleService(repo repository.Titles, repoCate repository.Categories) *TitleService {
	return &TitleService{repoTitles: repo, repoCategories: repoCate}
}

func (s *TitleService) GetTitles() ([]app.Title, error) {
	return s.repoTitles.GetAllTitles()
}

func (s *TitleService) GetAllPossibleContent() (app.AllContent, error) {
	return s.repoTitles.GetAllPossibleContent()
}

func (s *TitleService) GetTitle(id int) (app.Title, error) {
	return s.repoTitles.GetTitleById(id)
}

func (s *TitleService) LikeOrUnlike(likeObj app.Like, like bool) error {
	if like {
		return s.repoTitles.LikeById(likeObj)
	}
	return s.repoTitles.UnLikeById(likeObj)
}

func (s *TitleService) GetRandom() (app.Title, error) {
	return s.repoTitles.GetRandom()
}

func (s *TitleService) CreateTitle(title app.Title) (uint, error) {
	return s.repoTitles.SaveNewTitle(title)
}

func (s *TitleService) UpdateTitle(title app.Title, id int) (uint, error) {
	return s.repoTitles.UpdateTitle(title, id)
}

func (s *TitleService) UpdateViewsForTitle(id int) error {
	return s.repoTitles.AddView(id)
}

func (s *TitleService) GetImagesByTitleId(id int) ([]app.Image, error) {
	return s.repoTitles.GetImagesByTitleId(id)
}

func (s *TitleService) GetFilteredTitlesByParams(byLike, byDate, byViews bool, cats []app.Category, tags []app.Tag, serials []app.Serial) ([]app.Title, error) {
	titles, err := s.GetFilteredTitles(cats, tags, serials)
	//var orders []string
	//str := "SELECT * FROM titles INNER JOIN title_contents tc on titles.id = tc.title_id ORDER BY tc.likes_count DESC, tc.views DESC, titles.creation_date DESC"
	//if !byLike && !byDate && !byViews {
	//	return titles, err
	//}
	//if byDate {
	//	str +=
	//}
	//if byLike {
	//
	//}
	//if byViews {
	//
	//}
	return titles, err
}

func (s *TitleService) GetFilteredTitles(cats []app.Category, tags []app.Tag, serials []app.Serial) ([]app.Title, error) {
	if len(cats) == 0 && len(tags) == 0 && len(serials) == 0 {
		return s.GetTitles()
	}

	set := make(map[uint]app.Title)
	var titles []app.Title
	if len(cats) != 0 {
		var arrInt []uint
		for _, category := range cats {
			arrInt = append(arrInt, category.ID)
		}
		arr, _ := s.repoTitles.GetTitlesByCategories(arrInt)
		titles = append(titles, arr...)
	}

	if len(tags) != 0 {
		var arrInt []uint
		for _, tag := range tags {
			arrInt = append(arrInt, tag.ID)
		}
		arr, _ := s.repoTitles.GetTitlesByTags(arrInt)
		titles = append(titles, arr...)
	}

	if len(serials) != 0 {
		var arrInt []uint
		for _, ser := range serials {
			arrInt = append(arrInt, ser.ID)
		}
		arr, _ := s.repoTitles.GetTitlesBySerials(arrInt)
		titles = append(titles, arr...)
	}

	for _, title := range titles {
		set[title.ID] = title
	}
	res := make([]app.Title, 0, len(set))
	for _, t := range set {
		res = append(res, t)
	}

	return res, nil
}

func (s *TitleService) DeleteTitle(id int) (uint, error) {
	return s.repoTitles.DeleteTitleById(id)
}

func (s *TitleService) GetTitlesByCategories(categories []app.Category) ([]app.Title, error) {
	var arrInt []uint
	for _, category := range categories {
		arrInt = append(arrInt, category.ID)
	}
	return s.repoTitles.GetTitlesByCategories(arrInt)
}
