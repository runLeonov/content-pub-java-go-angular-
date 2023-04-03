package com.example.contentpublicationadmin.service;

import com.example.contentpublicationadmin.entity.*;

import java.util.List;

public interface ITitlePropertiesService {

    List<Category> getAllCategories();

    List<Tag> getAllTags();

    List<Serial> getAllSerials();

    List<StaticType> getAllTypes();

    Category createCategory(Category tag);

    Tag createTag(Tag tag);

    Serial createSerial(Serial tag);

    StaticType createStaticType(StaticType tag);

     boolean deleteCategory(Long id);

     boolean deleteType(Long id);

     boolean deleteTag(Long id);

     boolean deleteSerial(Long id);

}
