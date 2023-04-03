package com.example.contentpublicationadmin.service.impl;

import com.example.contentpublicationadmin.dao.TitleDAO;
import com.example.contentpublicationadmin.entity.Title;
import com.example.contentpublicationadmin.service.ITitleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;


@Service
public class TitleService implements ITitleService {
    @Autowired
    private TitleDAO titleDAO;

    @Override
    public List<Title> getAllTitles() {
        return titleDAO.findAll();
    }

    @Override
    public Title getTitleById(Long id) {
        return titleDAO.findById(id);
    }

    @Override
    public List<Title> getTitlesByCategories(List<Long> categoryId) {
        return null;
    }

    @Override
    public List<Title> getTitlesByTags(List<Long> tagId) {
        return null;
    }

    @Override
    public List<Title> getTitlesBySerials(List<Long> serialId) {
        return null;
    }

    @Override
    public List<Title> searchTitles(String keyword) {
        return null;
    }

    @Override
    public Title createTitle(Title title) {
        return null;
    }

    @Override
    public Title updateTitle(Long id, Title title) {
        return null;
    }

    @Override
    public boolean deleteTitle(Long id) {
        return false;
    }
}
