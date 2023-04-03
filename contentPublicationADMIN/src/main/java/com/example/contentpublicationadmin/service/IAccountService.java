package com.example.contentpublicationadmin.service;

import com.example.contentpublicationadmin.entity.User;

import java.util.List;

public interface IAccountService {

    List<User> getAllUsers();

    User getUserById(Long id);

    User createUser(User user);

    User updateUser(User title);

    boolean deleteUser(Long id);
}
