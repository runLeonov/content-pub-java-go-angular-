package com.example.contentpublicationadmin.controller;


import com.example.contentpublicationadmin.entity.User;
import com.example.contentpublicationadmin.service.impl.AccountService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.validation.BindingResult;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.servlet.ModelAndView;

import javax.validation.Valid;
import java.util.Objects;

@Controller()
@RequestMapping("/registration")
public class RegistrationController {


    @Autowired
    AccountService accountService;

    @GetMapping("/")
    public ModelAndView getPage(ModelMap modelMap) {
        modelMap.addAttribute("usr", new User());
        return new ModelAndView("registration", modelMap);
    }

    @GetMapping("")
    public ModelAndView redirect(ModelMap modelMap) {
        return getPage(modelMap);
    }


    @PostMapping()
    public ModelAndView registration(@ModelAttribute("usr") @Valid User user, BindingResult bindingResult, ModelMap modelMap) {
        if (bindingResult.hasErrors()) {
            return new ModelAndView("registration", modelMap);
        }

        if (Objects.nonNull(((User) accountService.loadUserByUsername(user.getUsername())).getID()  )) {
            modelMap.addAttribute("err", "User already exist!");
            return new ModelAndView("registration", modelMap);
        }

        accountService.createUser(user);
        return new ModelAndView("redirect:/login", modelMap);
    }

}
