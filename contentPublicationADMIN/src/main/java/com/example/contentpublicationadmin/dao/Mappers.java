package com.example.contentpublicationadmin.dao;

import com.example.contentpublicationadmin.entity.*;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.List;

public class Mappers {


    public static final class UserMapper implements RowMapper<User> {
        private JdbcTemplate jdbcTemplate;

        public UserMapper(JdbcTemplate jdbcTemplate) {
            this.jdbcTemplate = jdbcTemplate;
        }

        public User mapRow(ResultSet rs, int rowNum) throws SQLException {
            User user = new User();
            user.setID(rs.getInt("id"));
            user.setName(rs.getString("name"));
            user.setName(rs.getString("email"));
            user.setName(rs.getString("img"));
            user.setName(rs.getString("role"));
            user.setName(rs.getString("last_name"));
            user.setName(rs.getString("first_name"));

            String sql = "SELECT likes.* FROM likes JOIN users ON users.id = likes.user_id WHERE likes.user_id = ?";
            List<Like> likes = jdbcTemplate.query(sql, new Object[]{user.getID()}, new Mappers.LikeMapper(jdbcTemplate));
            user.setLikes(likes);

            sql = "SELECT comments.* FROM comments JOIN users ON users.id = comments.user_id WHERE comments.user_id = ?";
            List<Comment> comments = jdbcTemplate.query(sql, new Object[]{user.getID()}, new Mappers.CommentMapper(jdbcTemplate));
            user.setComments(comments);

            return user;
        }
    }

    public static final class ImagesMapper implements RowMapper<Image> {
        private JdbcTemplate jdbcTemplate;

        public ImagesMapper(JdbcTemplate jdbcTemplate) {
            this.jdbcTemplate = jdbcTemplate;
        }

        public Image mapRow(ResultSet rs, int rowNum) throws SQLException {
            Image image = new Image();
            image.setID(rs.getLong("id"));
            image.setImgBase64(rs.getString("img_base64"));
            image.setTitleContentID(rs.getLong("title_content_id"));

            return image;
        }
    }

    public static final class LikeMapper implements RowMapper<Like> {

        private JdbcTemplate jdbcTemplate;

        public LikeMapper(JdbcTemplate jdbcTemplate) {
            this.jdbcTemplate = jdbcTemplate;
        }

        public Like mapRow(ResultSet rs, int rowNum) throws SQLException {
            Like o = new Like();
            o.setTitleContentID(rs.getLong("title_content_id"));
            o.setUserID(rs.getLong("user_id"));

            return o;
        }
    }

    public static final class CategoryMapper implements RowMapper<Category> {

        private JdbcTemplate jdbcTemplate;

        public CategoryMapper(JdbcTemplate jdbcTemplate) {
            this.jdbcTemplate = jdbcTemplate;
        }

        public Category mapRow(ResultSet rs, int rowNum) throws SQLException {
            Category o = new Category();
            o.setGenre(rs.getString("genre"));
            o.setID(rs.getLong("id"));

            return o;
        }
    }

    public static final class TagMapper implements RowMapper<Tag> {
        private JdbcTemplate jdbcTemplate;

        public TagMapper(JdbcTemplate jdbcTemplate) {
            this.jdbcTemplate = jdbcTemplate;
        }

        public Tag mapRow(ResultSet rs, int rowNum) throws SQLException {
            Tag o = new Tag();
            o.setTagName(rs.getString("tag_name"));
            o.setID(rs.getLong("id"));

            return o;
        }
    }

    public static final class SerialMapper implements RowMapper<Serial> {
        private JdbcTemplate jdbcTemplate;

        public SerialMapper(JdbcTemplate jdbcTemplate) {
            this.jdbcTemplate = jdbcTemplate;
        }

        public Serial mapRow(ResultSet rs, int rowNum) throws SQLException {
            Serial o = new Serial();
            o.setSerialName(rs.getString("serial_name"));
            o.setID(rs.getLong("id"));

            return o;
        }
    }

    public static final class CommentMapper implements RowMapper<Comment> {
        private JdbcTemplate jdbcTemplate;

        public CommentMapper(JdbcTemplate jdbcTemplate) {
            this.jdbcTemplate = jdbcTemplate;
        }

        public Comment mapRow(ResultSet rs, int rowNum) throws SQLException {
            Comment o = new Comment();
            o.setCommentVal(rs.getString("comment_val"));
            o.setCreationDate(rs.getTimestamp("creation_date").toLocalDateTime());
            o.setUserID(rs.getLong("user_id"));
            o.setTitleContentID(rs.getLong("title_content_id"));

            return o;
        }
    }

    public static final class TitleContentMapper implements RowMapper<TitleContent> {
        private JdbcTemplate jdbcTemplate;

        public TitleContentMapper(JdbcTemplate jdbcTemplate) {
            this.jdbcTemplate = jdbcTemplate;
        }

        public TitleContent mapRow(ResultSet rs, int rowNum) throws SQLException {
            TitleContent o = new TitleContent();
            o.setID(rs.getLong("id"));
            o.setTitleID(rs.getLong("title_id"));
            o.setLikesCount(rs.getInt("likes_count"));
            o.setViews(rs.getInt("views"));


            String sql = "SELECT c.* FROM comments c JOIN title_contents t ON c.title_content_id = t.id WHERE c.title_content_id = ?";
            List<Comment> comments = jdbcTemplate.query(sql, new Object[]{o.getID()}, new Mappers.CommentMapper(jdbcTemplate));
            o.setComments(comments);

            sql = "SELECT i.* FROM images i JOIN title_contents tc ON tc.id = i.title_content_id WHERE i.title_content_id = ?";
            List<Image> images = jdbcTemplate.query(sql, new Object[]{o.getID()}, new Mappers.ImagesMapper(jdbcTemplate));
            o.setImages(images);

            return o;
        }
    }

    public static final class StaticTypeMapper implements RowMapper<StaticType> {
        private JdbcTemplate jdbcTemplate;

        public StaticTypeMapper(JdbcTemplate jdbcTemplate) {
            this.jdbcTemplate = jdbcTemplate;
        }

        public StaticType mapRow(ResultSet rs, int rowNum) throws SQLException {
            StaticType o = new StaticType();
            o.setID(rs.getLong("id"));
            o.setTypeName(rs.getString("type_name"));

            return o;
        }
    }

    public static final class TitleMapper implements RowMapper<Title> {
        private JdbcTemplate jdbcTemplate;

        public TitleMapper(JdbcTemplate jdbcTemplate) {
            this.jdbcTemplate = jdbcTemplate;
        }

        public Title mapRow(ResultSet rs, int rowNum) throws SQLException {
            Title o = new Title();
            o.setID(rs.getLong("id"));
            o.setAuthorId(rs.getLong("author_id"));
            o.setTitleName(rs.getString("title_name"));
            o.setOriginalAuthor(rs.getString("original_author"));
            o.setCreationDate(rs.getTimestamp("creation_date").toLocalDateTime());
            o.setDescription(rs.getString("description"));
            o.setTitleImgBase64(rs.getString("title_img_base64"));
            o.setReleased(rs.getBoolean("released"));


            String sql = "SELECT c.* FROM title_contents c INNER JOIN titles tc ON tc.id = c.title_id WHERE tc.id = ?";
            TitleContent content = jdbcTemplate.queryForObject(sql, new Object[]{o.getID()}, new Mappers.TitleContentMapper(jdbcTemplate));
            o.setContent(content);

            sql = "SELECT c.* FROM categories c INNER JOIN titles_categories tc ON tc.category_id = c.id WHERE tc.title_id = ?";
            List<Category> categories = jdbcTemplate.query(sql, new Object[]{o.getID()}, new Mappers.CategoryMapper(jdbcTemplate));
            o.setCategories(categories);

            sql = "SELECT c.* FROM tags c INNER JOIN titles_tags tc ON tc.tag_id = c.id WHERE tc.title_id = ?";
            List<Tag> tags = jdbcTemplate.query(sql, new Object[]{o.getID()}, new Mappers.TagMapper(jdbcTemplate));
            o.setTags(tags);

            sql = "SELECT c.* FROM serials c INNER JOIN titles_serials tc ON tc.serial_id = c.id WHERE tc.title_id = ?";
            List<Serial> serials = jdbcTemplate.query(sql, new Object[]{o.getID()}, new Mappers.SerialMapper(jdbcTemplate));
            o.setSerials(serials);

            return o;
        }
    }

    public static final class AllContentMapper implements RowMapper<AllContent> {
        private JdbcTemplate jdbcTemplate;

        public AllContentMapper(JdbcTemplate jdbcTemplate) {
            this.jdbcTemplate = jdbcTemplate;
        }

        public AllContent mapRow(ResultSet rs, int rowNum) throws SQLException {
            AllContent o = new AllContent();

            String sql = "SELECT * FROM categories";
            List<Category> categories = jdbcTemplate.query(sql, new Mappers.CategoryMapper(jdbcTemplate));
            sql = "SELECT * FROM tags";
            List<Tag> tags = jdbcTemplate.query(sql, new Mappers.TagMapper(jdbcTemplate));
            sql = "SELECT * FROM serials";
            List<Serial> serials = jdbcTemplate.query(sql, new Mappers.SerialMapper(jdbcTemplate));
            sql = "SELECT * FROM static_types";
            List<StaticType> types = jdbcTemplate.query(sql, new Mappers.StaticTypeMapper(jdbcTemplate));

            o.setCategories(categories);
            o.setTags(tags);
            o.setSerials(serials);
            o.setTypes(types);
            return o;
        }
    }

}
