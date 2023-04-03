package com.example.contentpublicationadmin.entity;

import lombok.Data;

@Data
public class Image {
    private Long ID;
    private String imgBase64;
    private Long titleContentID;
    private TitleContent titleContent;
}
