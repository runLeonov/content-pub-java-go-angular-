package com.example.contentpublicationadmin.dao;

import com.example.contentpublicationadmin.entity.Comment;
import com.example.contentpublicationadmin.entity.User;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import javax.sql.DataSource;
import java.util.List;


@Repository
public class CommentDAO {

    private final JdbcTemplate jdbcTemplate;

    public CommentDAO(DataSource dataSource) {
        this.jdbcTemplate = new JdbcTemplate(dataSource);
    }

    public List<Comment> getAllComments() {
        String sql = "SELECT * FROM comments";
        return jdbcTemplate.query(sql, new Mappers.CommentMapper(jdbcTemplate));
    }

    public List<Comment> getAllFilteredComments(String filter) {
        String sql = "SELECT * FROM comments WHERE comment_val LIKE ?";
        return jdbcTemplate.query(sql, new String[]{'%' + filter + '%'}, new Mappers.CommentMapper(jdbcTemplate));
    }

    public List<Comment> findAllByID() {
        String sql = "SELECT * FROM comments ORDER BY id DESC ";
        return jdbcTemplate.query(sql, new Mappers.CommentMapper(jdbcTemplate));
    }

    public List<Comment> findAllByDate() {
        String sql = "SELECT * FROM comments ORDER BY creation_date DESC ";
        return jdbcTemplate.query(sql, new Mappers.CommentMapper(jdbcTemplate));
    }

    public Comment getCommentById(Long id) {
        String sql = "SELECT * FROM comments WHERE id=?";
        return jdbcTemplate.queryForObject(sql, new Object[]{id}, new Mappers.CommentMapper(jdbcTemplate));
    }


    public void deleteComment(Long id) {
        String sql = "DELETE FROM comments WHERE id=?";
        jdbcTemplate.update(sql, id);
    }
}
