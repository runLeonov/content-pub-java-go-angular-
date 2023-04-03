package com.example.contentpublicationadmin.entity;

import lombok.Data;

import java.util.List;

@Data
public class TitleContent {
    private Long ID;
    private Title title;
    private Long titleID;
    private int likesCount;
    private int views;
    private List<Like> likes;
    private List<Image> images;
    private List<Comment> comments;
}
