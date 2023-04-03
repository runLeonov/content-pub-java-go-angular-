package com.example.contentpublicationadmin.service.impl;

import com.example.contentpublicationadmin.dao.AccountDAO;
import com.example.contentpublicationadmin.entity.User;
import com.example.contentpublicationadmin.service.IAccountService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class AccountService implements IAccountService {

    @Autowired
    private AccountDAO accountDAO;

    @Override
    public List<User> getAllUsers() {
        return accountDAO.getAllUsers();
    }

    @Override
    public User getUserById(Long id) {
        return accountDAO.getUserById(id);
    }

    @Override
    public User createUser(User user) {
        accountDAO.addUser(user);
        return user;
    }

    @Override
    public User updateUser(User user) {
        accountDAO.updateUser(user);
        return user;
    }

    @Override
    public boolean deleteUser(Long id) {
        accountDAO.deleteUser(id);
        return true;
    }
}
