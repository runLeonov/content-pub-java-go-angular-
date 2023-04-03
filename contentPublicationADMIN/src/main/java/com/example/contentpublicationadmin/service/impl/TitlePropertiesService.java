package com.example.contentpublicationadmin.service.impl;

import com.example.contentpublicationadmin.dao.TitlePropsDAO;
import com.example.contentpublicationadmin.entity.Category;
import com.example.contentpublicationadmin.entity.Serial;
import com.example.contentpublicationadmin.entity.StaticType;
import com.example.contentpublicationadmin.entity.Tag;
import com.example.contentpublicationadmin.service.ITitlePropertiesService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class TitlePropertiesService implements ITitlePropertiesService {

    @Autowired
    private TitlePropsDAO titlePropsDAO;

    @Override
    public List<Category> getAllCategories() {
        return titlePropsDAO.findCategories();
    }

    @Override
    public List<Tag> getAllTags() {
        return titlePropsDAO.findTags();
    }

    @Override
    public List<Serial> getAllSerials() {
        return titlePropsDAO.findSerials();
    }

    @Override
    public List<StaticType> getAllTypes() {
        return titlePropsDAO.findTypes();
    }

    @Override
    public Category createCategory(Category cat) {
        titlePropsDAO.saveCategory(cat);
        return cat;
    }

    @Override
    public Tag createTag(Tag tag) {
        titlePropsDAO.saveTag(tag);
        return tag;
    }

    @Override
    public Serial createSerial(Serial ser) {
        titlePropsDAO.saveSerial(ser);
        return ser;
    }

    @Override
    public StaticType createStaticType(StaticType staticType) {
        titlePropsDAO.saveType(staticType);
        return staticType;
    }

    @Override
    public boolean deleteType(Long id) {
        titlePropsDAO.delType(id);
        return true;
    }

    @Override
    public boolean deleteTag(Long id) {
        titlePropsDAO.delTag(id);
        return true;
    }

    @Override
    public boolean deleteCategory(Long id) {
        titlePropsDAO.delSerial(id);
        return true;
    }

    @Override
    public boolean deleteSerial(Long id) {
        titlePropsDAO.delSerial(id);
        return true;
    }
}
