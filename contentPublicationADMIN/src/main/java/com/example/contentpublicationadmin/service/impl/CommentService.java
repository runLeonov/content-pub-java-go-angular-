package com.example.contentpublicationadmin.service.impl;


import com.example.contentpublicationadmin.dao.AccountDAO;
import com.example.contentpublicationadmin.dao.CommentDAO;
import com.example.contentpublicationadmin.entity.Comment;
import com.example.contentpublicationadmin.entity.Sort;
import com.example.contentpublicationadmin.entity.User;
import com.example.contentpublicationadmin.service.ICommentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class CommentService implements ICommentService {

    @Autowired
    private CommentDAO commentDAO;

    @Override
    @Transactional
    public Comment deleteComment(Long id) {
        commentDAO.deleteComment(id);
        return commentDAO.getCommentById(id);
    }

    @Override
    public List<Comment> getAllComments() {
        return commentDAO.getAllComments();
    }

    @Override
    public List<Comment> getAllCommentsSorted(Sort sort) {
        if (sort == Sort.ID) return commentDAO.findAllByID();
        else if (sort == Sort.DATE) return commentDAO.findAllByDate();
        else return commentDAO.getAllComments();
    }

    @Override
    public List<Comment> getAllFilteredComments(String filter) {
        return commentDAO.getAllFilteredComments(filter);
    }

    @Override
    public Comment getCommentById(Long id) {
        return commentDAO.getCommentById(id);
    }
}
