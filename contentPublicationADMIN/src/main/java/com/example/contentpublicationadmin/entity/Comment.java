package com.example.contentpublicationadmin.entity;

import lombok.Data;

import java.time.LocalDateTime;

@Data
public class Comment {
    private Long ID;
    private String commentVal;
    private LocalDateTime creationDate;
    private User user;
    private Long userID;
    private TitleContent titleContent;
    private Long titleContentID;
}
