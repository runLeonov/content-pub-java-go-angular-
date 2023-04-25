package com.example.contentpublicationadmin.service;

import com.example.contentpublicationadmin.entity.*;

import java.util.List;

public interface IContentService {
    AllContent getAllContent();

    Tag updateTag(Tag o);

    Category updateCategory(Category o);

    Serial updateSerial(Serial o);

    Tag addTag(Tag o);

    Category addCategory(Category o);

    Serial addSerial(Serial o);

    boolean deleteSerial(Long id);

    boolean deleteCategory(Long id);

    boolean deleteTag(Long id);

}
