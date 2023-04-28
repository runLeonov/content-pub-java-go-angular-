package com.example.contentpublicationadmin.entity;

import lombok.Data;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;
import org.thymeleaf.expression.Lists;

import javax.validation.constraints.Email;
import javax.validation.constraints.NotEmpty;
import javax.validation.constraints.Size;
import java.time.LocalDateTime;
import java.util.Arrays;
import java.util.Collection;
import java.util.Date;
import java.util.List;

@Data
public class User implements UserDetails {
    private Long ID;
    @Email
    @NotEmpty(message = "Email field can't be empty!")
    private String email;
    @NotEmpty(message = "Password can't be empty!")
    @Size(min = 6, message = "Password must contains at least 6 chars")
    private String password;
    @NotEmpty(message = "Name field can't be empty!")
    private String name;
    private String img;
    private String role = "USER";
    private List<Like> likes;
    private List<Comment> comments;
    private LocalDateTime creationDate = LocalDateTime.now();
    private String lastName;
    private String firstName;
    private boolean banned;

    @Override
    public Collection<? extends GrantedAuthority> getAuthorities() {
        return List.of((GrantedAuthority) this::getRole);
    }

    @Override
    public String getUsername() {
        return name;
    }

    @Override
    public boolean isAccountNonExpired() {
        return true;
    }

    @Override
    public boolean isAccountNonLocked() {
        return !banned;
    }

    @Override
    public boolean isCredentialsNonExpired() {
        return true;
    }


    @Override
    public String getPassword() {
        return password;
    }

    @Override
    public boolean isEnabled() {
        return !banned;
    }
}
