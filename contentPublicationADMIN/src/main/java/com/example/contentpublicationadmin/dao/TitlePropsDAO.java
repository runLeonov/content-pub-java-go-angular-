package com.example.contentpublicationadmin.dao;

import com.example.contentpublicationadmin.entity.*;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public class TitlePropsDAO {

    private final JdbcTemplate jdbcTemplate;

    public TitlePropsDAO(JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }

    public List<Tag> findTags() {
        String sql = "SELECT * FROM tags";
        return jdbcTemplate.query(sql, new Mappers.TagMapper(jdbcTemplate));
    }

    public List<StaticType> findTypes() {
        String sql = "SELECT * FROM static_types";
        return jdbcTemplate.query(sql, new Mappers.StaticTypeMapper(jdbcTemplate));
    }

    public List<Serial> findSerials() {
        String sql = "SELECT * FROM serials";
        return jdbcTemplate.query(sql, new Mappers.SerialMapper(jdbcTemplate));
    }

    public List<Category> findCategories() {
        String sql = "SELECT * FROM categories";
        return jdbcTemplate.query(sql, new Mappers.CategoryMapper(jdbcTemplate));
    }

    public void saveTag(Tag o) {
        String sql = "INSERT INTO tags (tag_name) VALUES (?)";
        jdbcTemplate.update(sql, o.getTagName());
    }

    public void saveCategory(Category o) {
        String sql = "INSERT INTO categories (genre) VALUES (?)";
        jdbcTemplate.update(sql, o.getGenre());
    }

    public void saveSerial(Serial o) {
        String sql = "INSERT INTO serials (serial_name) VALUES (?)";
        jdbcTemplate.update(sql, o.getSerialName());
    }

    public void saveType(StaticType o) {
        String sql = "INSERT INTO static_types (type_name) VALUES (?)";
        jdbcTemplate.update(sql, o.getTypeName());
    }

    public void delType(Long id) {
        String sql = "DELETE FROM static_types WHERE id = ?";
        jdbcTemplate.update(sql, id);
    }

    public void delTag(Long id) {
        String sql = "DELETE FROM titles_tags WHERE tag_id = ?";
        jdbcTemplate.update(sql, id);
        sql = "DELETE FROM tags WHERE id = ?";
        jdbcTemplate.update(sql, id);
    }

    public void delCategory(Long id) {
        String sql = "DELETE FROM titles_categories WHERE category_id = ?";
        jdbcTemplate.update(sql, id);
        sql = "DELETE FROM categories WHERE id = ?";
        jdbcTemplate.update(sql, id);
    }

    public void delSerial(Long id) {
        String sql = "DELETE FROM titles_serials WHERE serial_id = ?";
        jdbcTemplate.update(sql, id);
        sql = "DELETE FROM serials WHERE id = ?";
        jdbcTemplate.update(sql, id);
    }


}
