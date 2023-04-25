package com.example.contentpublicationadmin.entity;

import lombok.Data;

import java.time.LocalDateTime;
import java.util.List;

@Data
public class User {
    private Long ID;
    private String email;
    private String password;
    private String name;
    private String img;
    private String role = "USER";
    private List<Like> likes;
    private List<Comment> comments;
    private LocalDateTime creationDate;
    private String lastName;
    private String firstName;
    private boolean banned;

}
