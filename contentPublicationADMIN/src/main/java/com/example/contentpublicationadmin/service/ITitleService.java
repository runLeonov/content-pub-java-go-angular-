package com.example.contentpublicationadmin.service;

import com.example.contentpublicationadmin.entity.Title;

import java.util.List;

public interface ITitleService {

    List<Title> getAllTitles();

    Title getTitleById(Long id);

    List<Title> getTitlesByCategories(List<Long> categoryId);

    List<Title> getTitlesByTags(List<Long> tagId);

    List<Title> getTitlesBySerials(List<Long> serialId);

    List<Title> searchTitles(String keyword);

    Title createTitle(Title title);

    Title updateTitle(Long id, Title title);

    boolean deleteTitle(Long id);

}