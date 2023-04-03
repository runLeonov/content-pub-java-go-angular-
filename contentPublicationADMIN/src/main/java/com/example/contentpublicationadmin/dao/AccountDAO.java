package com.example.contentpublicationadmin.dao;

import com.example.contentpublicationadmin.entity.User;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Repository;

import javax.sql.DataSource;
import java.util.List;

@Repository
public class AccountDAO {

    private final JdbcTemplate jdbcTemplate;

    public AccountDAO(DataSource dataSource) {
        this.jdbcTemplate = new JdbcTemplate(dataSource);
    }

    public List<User> getAllUsers() {
        String sql = "SELECT * FROM users";
        return jdbcTemplate.query(sql, new Mappers.UserMapper(jdbcTemplate));
    }

    public User getUserById(Long id) {
        String sql = "SELECT * FROM users WHERE id=?";
        return jdbcTemplate.queryForObject(sql, new Object[]{id}, new Mappers.UserMapper(jdbcTemplate));
    }

    public void addUser(User user) {
        String sql = "INSERT INTO users (name) VALUES (?)";
        jdbcTemplate.update(sql, user.getName());
    }

    public void updateUser(User user) {
        String sql = "UPDATE users SET name=? WHERE id=?";
        jdbcTemplate.update(sql, user.getName(), user.getID());
    }

    public void deleteUser(Long id) {
        String sql = "DELETE FROM users WHERE id=?";
        jdbcTemplate.update(sql, id);
    }

}
