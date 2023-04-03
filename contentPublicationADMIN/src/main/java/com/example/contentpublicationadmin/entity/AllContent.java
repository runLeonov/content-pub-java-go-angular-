package com.example.contentpublicationadmin.entity;

import lombok.Data;

import java.util.List;

@Data
public class AllContent {
    private List<StaticType> types;
    private List<Category> categories;
    private List<Tag> tags;
    private List<Serial> serials;
}

