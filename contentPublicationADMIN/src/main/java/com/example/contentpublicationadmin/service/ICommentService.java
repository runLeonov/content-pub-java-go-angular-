package com.example.contentpublicationadmin.service;

import com.example.contentpublicationadmin.entity.Comment;
import com.example.contentpublicationadmin.entity.Sort;
import com.example.contentpublicationadmin.entity.User;

import java.util.List;

public interface ICommentService {

    List<Comment> getAllComments();

    List<Comment> getAllCommentsSorted(Sort sort);

    List<Comment> getAllFilteredComments(String filter);

    Comment getCommentById(Long id);

    Comment deleteComment(Long id);
}
