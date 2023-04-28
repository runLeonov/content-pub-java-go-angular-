package com.example.contentpublicationadmin.controller;

import com.example.contentpublicationadmin.entity.Sort;
import com.example.contentpublicationadmin.entity.Title;
import com.example.contentpublicationadmin.entity.User;
import com.example.contentpublicationadmin.service.IAccountService;
import com.example.contentpublicationadmin.service.ITitleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.ModelAndView;

import java.util.List;

@Controller
@RequestMapping("/users")
public class AccountController {
    @Autowired
    private IAccountService userService;
    @Autowired
    private ITitleService service;

    @GetMapping("/")
    public ModelAndView getAllUsers(ModelMap model) {
        List<User> users = userService.getAllUsers();
        model.addAttribute("users", users);
        return new ModelAndView("users", model);
    }

    @GetMapping("")
    public ModelAndView getAllUsersRedirect(ModelMap model) {
        return getAllUsers(model);
    }


    @GetMapping("/f")
    public ModelAndView filterUsers(@RequestParam(name = "filter", required = false) String filter, ModelMap model) {
        filter = filter.trim().toLowerCase();
        if (filter.isEmpty()) {
            model.addAttribute("users", userService.getAllUsers());
            return new ModelAndView("users", model);
        }
        List<User> users = userService.getFilteredUsers(filter);
        model.addAttribute("users", users);
        return new ModelAndView("users", model);
    }

    @GetMapping("/sort/id/")
    public ModelAndView sortByID(ModelMap model) {
        model.addAttribute("users", userService.getAllUsersSort(Sort.ID));
        model.addAttribute("byID", true);
        return new ModelAndView("users", model);
    }

    @GetMapping("/sort/date/")
    public ModelAndView sortByDate(ModelMap model) {
        model.addAttribute("users", userService.getAllUsersSort(Sort.DATE));
        model.addAttribute("byDate", true);
        return new ModelAndView("users", model);
    }


    @DeleteMapping("/{id}")
    public void deleteUser(@PathVariable Long id) {
        boolean isTrue = userService.deleteUser(id);
        return;
    }

    @GetMapping("/{id}")
    public ModelAndView getUser(@PathVariable Long id, ModelMap model) {
        User user = userService.getUserById(id);
        List<Title> titles = service.getAllUserTitles(user.getID());
        model.addAttribute("user", user);
        model.addAttribute("titles", titles);
        return new ModelAndView("user", model);
    }

    @PostMapping("/{id}/ban")
    public ModelAndView banOrUnban(@PathVariable Long id, ModelMap model) {
        User user = userService.getUserById(id);
        model.addAttribute("user", userService.banOrUnbanUser(user));
        return new ModelAndView("redirect:/users/" + id);
    }

    @PostMapping("/{id}/make-author")
    public ModelAndView giveAuthorRole(@PathVariable Long id, ModelMap model) {
        User user = userService.getUserById(id);
        model.addAttribute("user", userService.updateUserRole("AUTHOR" ,user));
        return new ModelAndView("redirect:/users/" + id);
    }

    @PostMapping("/{id}/cancel-author")
    public ModelAndView cancelAuthorRole(@PathVariable Long id, ModelMap model) {
        User user = userService.getUserById(id);
        model.addAttribute("user", userService.updateUserRole("USER" ,user));
        return new ModelAndView("redirect:/users/" + id);
    }

    @PostMapping("/{id}/comment/{commentId}")
    public ModelAndView deleteCommentByUser(ModelMap model, @PathVariable Long commentId, @PathVariable Long id) {
        User user = userService.getUserById(id);
        model.addAttribute("user", userService.deleteCommentByUser(user, commentId));
        return new ModelAndView("redirect:/users/" + id);
    }
}
