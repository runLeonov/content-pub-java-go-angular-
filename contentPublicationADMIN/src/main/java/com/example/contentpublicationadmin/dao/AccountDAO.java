package com.example.contentpublicationadmin.dao;

import com.example.contentpublicationadmin.entity.Title;
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

    public List<User> findAllByID() {
        String sql = "SELECT * FROM users ORDER BY id DESC ";
        return jdbcTemplate.query(sql, new Mappers.UserMapper(jdbcTemplate));
    }

    public List<User> findAllByDate() {
        String sql = "SELECT * FROM users ORDER BY creation_date DESC ";
        return jdbcTemplate.query(sql, new Mappers.UserMapper(jdbcTemplate));
    }

    public User getUserById(Long id) {
        String sql = "SELECT * FROM users WHERE id=?";
        return jdbcTemplate.queryForObject(sql, new Object[]{id}, new Mappers.TitleUserMapper(jdbcTemplate));
    }

    public User getUserByName(String name) {
        String sql = "SELECT * FROM users WHERE name=? LIMIT 1";
        return jdbcTemplate.queryForObject(sql, new Object[]{name}, new Mappers.UserAuthMapper(jdbcTemplate));
    }

    public List<User> getFilteredUsers(String filter) {
        String sql = "SELECT * FROM users WHERE name LIKE ?";
        return jdbcTemplate.query(sql, new Object[]{'%' + filter + '%'}, new Mappers.UserMapper(jdbcTemplate));
    }

    public void banOrUnbanUserById(User user) {
        String sql = "UPDATE users SET banned = ? WHERE id = ?";
        jdbcTemplate.update(sql, !user.isBanned(), user.getID());
    }

    public void deleteUsersComment(Long commentId) {
        String sql = "DELETE FROM comments WHERE id = ?";
        jdbcTemplate.update(sql, commentId);
    }

    public void addUserAdmin(User user) {
        String sql = "INSERT INTO users ( email, password, name, role, last_name, first_name) VALUES (?,?,?,?,?,?)";
        jdbcTemplate.update(sql, user.getEmail(), user.getPassword(), user.getName(), user.getRole(), user.getLastName(), user.getFirstName());
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
