package com.example.contentpublicationadmin.dao;

import com.example.contentpublicationadmin.entity.Title;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public class TitleDAO {

    private final JdbcTemplate jdbcTemplate;

    public TitleDAO(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }

    public List<Title> findAll() {
        String sql = "SELECT * FROM titles";
        return jdbcTemplate.query(sql, new Mappers.TitleMapper(jdbcTemplate));
    }

    public List<Title> findAllReleased() {
        String sql = "SELECT * FROM titles WHERE released = true";
        return jdbcTemplate.query(sql, new Mappers.TitleMapper(jdbcTemplate));
    }

    public List<Title> findAllUnReleased() {
        String sql = "SELECT * FROM titles WHERE released = false";
        return jdbcTemplate.query(sql, new Mappers.TitleMapper(jdbcTemplate));
    }

    public List<Title> findAllByID() {
        String sql = "SELECT * FROM titles ORDER BY id DESC ";
        return jdbcTemplate.query(sql, new Mappers.TitleMapper(jdbcTemplate));
    }

    public List<Title> findAllByDate() {
        String sql = "SELECT * FROM titles ORDER BY creation_date DESC";
        return jdbcTemplate.query(sql, new Mappers.TitleMapper(jdbcTemplate));
    }


    public List<Title> findFilteredAll(String filter) {
        String sql = "SELECT * FROM titles WHERE title_name LIKE ?";
        return jdbcTemplate.query(sql, new Object[]{'%' + filter + '%'}, new Mappers.TitleMapper(jdbcTemplate));
    }

    public Title findById(Long id) {
        String sql = "SELECT * FROM titles WHERE id = ?";
        return jdbcTemplate.queryForObject(sql, new Object[]{id}, new Mappers.TitleMapper(jdbcTemplate));
    }

    public void releaseOrBanTitleById(Title title) {
        String sql = "UPDATE titles SET released = ? WHERE id = ?";
        jdbcTemplate.update(sql, !title.isReleased(), title.getID());
    }
    public void deleteCommentForTitle(Long commentId) {
        String sql = "DELETE FROM comments WHERE id = ?";
        jdbcTemplate.update(sql, commentId);
    }

    public List<Title> findTitlesByUserId(Long id) {
        String sql = "SELECT * FROM titles WHERE user_id = ?";
        return jdbcTemplate.query(sql, new Object[]{id}, new Mappers.TitleMapper(jdbcTemplate));
    }

    public void save(Title title) {
        String sql = "INSERT INTO titles (name, author, published_date) VALUES (?, ?, ?)";
        jdbcTemplate.update(sql, title.getTitleName(), title.getOriginalAuthor(), title.getAuthorId());
    }

    public void update(Title title) {
        String sql = "UPDATE titles SET name = ?, author = ?, published_date = ? WHERE id = ?";
        jdbcTemplate.update(sql, title.getTitleName(), title.getOriginalAuthor(), title.getAuthorId());
    }

    public void deleteById(Long id) {
        String sql = "DELETE FROM titles WHERE id = ?";
        jdbcTemplate.update(sql, id);
    }

}
