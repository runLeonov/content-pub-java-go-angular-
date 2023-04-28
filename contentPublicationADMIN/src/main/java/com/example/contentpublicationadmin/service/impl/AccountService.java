package com.example.contentpublicationadmin.service.impl;

import com.example.contentpublicationadmin.dao.AccountDAO;
import com.example.contentpublicationadmin.entity.Sort;
import com.example.contentpublicationadmin.entity.Title;
import com.example.contentpublicationadmin.entity.User;
import com.example.contentpublicationadmin.service.IAccountService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.dao.EmptyResultDataAccessException;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.LocalDateTime;
import java.util.Date;
import java.util.List;
import java.util.Objects;

@Service
public class AccountService implements IAccountService, UserDetailsService {

    @Autowired
    private AccountDAO accountDAO;
    @Autowired
    BCryptPasswordEncoder bCryptPasswordEncoder;

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
        user.setRole("ADMIN");
        user.setCreationDate(LocalDateTime.now());
        user.setPassword(bCryptPasswordEncoder.encode(user.getPassword()));
        accountDAO.addUserAdmin(user);
        return user;
    }

    @Override
    public User updateUserRole(String role, User user) {
        accountDAO.giveRoleForUser(role, user);
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

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        UserDetails user = accountDAO.getUserByName(username);
        return Objects.isNull(user) ? new User() : user;
    }
}
