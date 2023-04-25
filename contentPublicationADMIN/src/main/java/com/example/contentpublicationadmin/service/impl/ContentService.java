package com.example.contentpublicationadmin.service.impl;

import com.example.contentpublicationadmin.dao.ContentsDAO;
import com.example.contentpublicationadmin.dao.TitleDAO;
import com.example.contentpublicationadmin.entity.AllContent;
import com.example.contentpublicationadmin.entity.Category;
import com.example.contentpublicationadmin.entity.Serial;
import com.example.contentpublicationadmin.entity.Tag;
import com.example.contentpublicationadmin.service.IContentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ContentService implements IContentService {

    @Autowired
    private ContentsDAO contentsDAO;

    @Override
    public AllContent getAllContent() {
        AllContent content = new AllContent();
        List<Tag> tags = contentsDAO.getAllTags();
        content.setTags(tags);

        List<Category> categories = contentsDAO.getAllCategories();
        content.setCategories(categories);

        List<Serial> serials = contentsDAO.getAllSerials();
        content.setSerials(serials);
        return content;
    }

    @Override
    public Tag updateTag(Tag o) {
        contentsDAO.updateTag(o);
        return o;
    }

    @Override
    public Category updateCategory(Category o) {
        contentsDAO.updateCategory(o);
        return o;
    }

    @Override
    public Serial updateSerial(Serial o) {
        contentsDAO.updateSerial(o);
        return o;
    }

    @Override
    public boolean deleteSerial(Long id) {
        contentsDAO.deleteSerial(id);
        return true;

    }

    @Override
    public boolean deleteCategory(Long id) {
        contentsDAO.deleteCategory(id);
        return true;
    }

    @Override
    public boolean deleteTag(Long id) {
        contentsDAO.deleteTag(id);
        return true;
    }

    @Override
    public Tag addTag(Tag o) {
        contentsDAO.addTag(o);
        return o;
    }

    @Override
    public Category addCategory(Category o) {
        contentsDAO.addCategory(o);
        return o;
    }

    @Override
    public Serial addSerial(Serial o) {
        contentsDAO.addSerial(o);
        return o;
    }
}
