package com.example.contentpublicationadmin.entity;

import lombok.Data;

import java.time.LocalDateTime;
import java.util.List;

@Data
public class Title {
    private Long ID;
    private String titleName;
    private String type;
    private String originalAuthor;
    private Long authorId;
    private LocalDateTime creationDate;
    private String description;
    private String titleImgBase64;
    private boolean released;
    private TitleContent content;
    private List<Category> categories;
    private List<Tag> tags;
    private List<Serial> serials;
}
