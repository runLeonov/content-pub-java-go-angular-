package com.example.contentpublicationadmin.dao;

import com.example.contentpublicationadmin.entity.*;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public class ContentsDAO {
    private final JdbcTemplate jdbcTemplate;

    public ContentsDAO(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }


    public List<Tag> getAllTags() {
        String sql = "SELECT * FROM tags";
        return jdbcTemplate.query(sql, new Mappers.TagMapper(jdbcTemplate));
    }

    public List<Category> getAllCategories() {
        String sql = "SELECT * FROM categories";
        return jdbcTemplate.query(sql, new Mappers.CategoryMapper(jdbcTemplate));
    }

    public List<Serial> getAllSerials() {
        String sql = "SELECT * FROM serials";
        return jdbcTemplate.query(sql, new Mappers.SerialMapper(jdbcTemplate));
    }

    public void addTag(Tag o) {
        String sql = "INSERT INTO tags (tag_name) VALUES (?)";
        jdbcTemplate.update(sql, o.getTagName());
    }

    public void addSerial(Serial o) {
        String sql = "INSERT INTO serials (serial_name) VALUES (?)";
        jdbcTemplate.update(sql, o.getSerialName());
    }

    public void addCategory(Category o) {
        String sql = "INSERT INTO categories (genre) VALUES (?)";
        jdbcTemplate.update(sql, o.getGenre());
    }

    public void updateTag(Tag o) {
        String sql = "UPDATE tags SET tag_name = ? WHERE id = ?";
        jdbcTemplate.update(sql, o.getTagName(), o.getID());
    }

    public void updateCategory(Category o) {
        String sql = "UPDATE categories SET genre = ? WHERE id = ?";
        jdbcTemplate.update(sql, o.getGenre(), o.getID());
    }

    public void updateSerial(Serial o) {
        String sql = "UPDATE serials SET serial_name = ? WHERE id = ?";
        jdbcTemplate.update(sql, o.getSerialName(), o.getID());
    }

    public void deleteSerial(Long id) {
        String sql = "DELETE FROM titles_serials WHERE serial_id = ?";
        jdbcTemplate.update(sql, id);
        sql = "DELETE FROM serials WHERE id = ?";
        jdbcTemplate.update(sql, id);
    }

    public void deleteCategory(Long id) {
        String sql = "DELETE FROM titles_categories WHERE category_id = ?";
        jdbcTemplate.update(sql, id);
        sql = "DELETE FROM categories WHERE id = ?";
        jdbcTemplate.update(sql, id);
    }

    public void deleteTag(Long id) {
        String sql = "DELETE FROM titles_tags WHERE tag_id = ?";
        jdbcTemplate.update(sql, id);
        sql = "DELETE FROM tags WHERE id = ?";
        jdbcTemplate.update(sql, id);
    }

}
