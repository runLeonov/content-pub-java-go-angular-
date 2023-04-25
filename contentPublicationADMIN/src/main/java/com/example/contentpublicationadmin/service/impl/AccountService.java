package com.example.contentpublicationadmin.service.impl;

import com.example.contentpublicationadmin.dao.AccountDAO;
import com.example.contentpublicationadmin.entity.Sort;
import com.example.contentpublicationadmin.entity.Title;
import com.example.contentpublicationadmin.entity.User;
import com.example.contentpublicationadmin.service.IAccountService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

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
    public List<User> getFilteredUsers(String filter) {
        return accountDAO.getFilteredUsers(filter);
    }

    @Override
    public List<User> getAllUsersSort(Sort sortBy) {
        if (sortBy == Sort.ID) return accountDAO.findAllByID();
        else if (sortBy == Sort.DATE) return accountDAO.findAllByDate();
        else return accountDAO.getAllUsers();
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

    @Override
    @Transactional
    public User deleteCommentByUser(User user, Long commentId) {
        accountDAO.deleteUsersComment(commentId);
        return accountDAO.getUserById(user.getID());
    }

    @Override
    @Transactional
    public User banOrUnbanUser(User user) {
        accountDAO.banOrUnbanUserById(user);
        return accountDAO.getUserById(user.getID());
    }
}
