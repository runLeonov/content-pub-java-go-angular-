package com.example.contentpublicationadmin.controller;

import com.example.contentpublicationadmin.entity.*;
import com.example.contentpublicationadmin.service.IContentService;
import com.example.contentpublicationadmin.service.ITitleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.ModelAndView;

import java.util.List;
import java.util.Objects;


@Controller
@RequestMapping("/contents")
public class ContentController {

    @Autowired
    private IContentService service;

    @GetMapping("/")
    public ModelAndView getContents(ModelMap model) {
        AllContent content = service.getAllContent();
        model.addAttribute("contents", content);

        model.addAttribute("categoryInfo", new Category());
        model.addAttribute("tagInfo", new Tag());
        model.addAttribute("serialInfo", new Serial());
        return new ModelAndView("contents", model);
    }

    @GetMapping("")
    public ModelAndView getContentsRedirect(ModelMap model) {
        return getContents(model);
    }

    @PutMapping("/update-t")
    public ModelAndView updateTag(@ModelAttribute Tag o, ModelMap model) {
        model.addAttribute("updateTagOK", Objects.nonNull(service.updateTag(o)));
        return new ModelAndView("redirect:/contents", model);
    }

    @PutMapping("/update-s")
    public ModelAndView updateSerial(@ModelAttribute Serial o, ModelMap model) {
        model.addAttribute("updateSerialOK", Objects.nonNull(service.updateSerial(o)));
        return new ModelAndView("redirect:/contents", model);
    }

    @PutMapping("/update-c")
    public ModelAndView updateCategory(@ModelAttribute Category category, ModelMap model) {
        model.addAttribute("updateCategoryOK", Objects.nonNull(service.updateCategory(category)));
        return new ModelAndView("redirect:/contents", model);
    }


    @PostMapping("/create-t")
    public ModelAndView addTag(@ModelAttribute Tag tagInfo, ModelMap model) {
        model.addAttribute("addTagOK", Objects.nonNull(service.addTag(tagInfo)));
        return new ModelAndView("redirect:/contents", model);
    }

    @PostMapping("/create-s")
    public ModelAndView addSerial(@ModelAttribute Serial serialInfo, ModelMap model) {
        model.addAttribute("addSerialOK", Objects.nonNull(service.addSerial(serialInfo)));
        return new ModelAndView("redirect:/contents", model);
    }

    @PostMapping("/create-c")
    public ModelAndView addCategory(@ModelAttribute Category categoryInfo, ModelMap model) {
        model.addAttribute("addCategoryOK", Objects.nonNull(service.addCategory(categoryInfo)));
        return new ModelAndView("redirect:/contents", model);
    }

    @PostMapping("/t/{id}")
    public ModelAndView deleteTag(@PathVariable Long id, ModelMap model) {
        model.addAttribute("deleteTagOK", service.deleteTag(id));
        return new ModelAndView("redirect:/contents", model);
    }

    @PostMapping("/c/{id}")
    public ModelAndView deleteCategory(@PathVariable Long id, ModelMap model) {
        model.addAttribute("deleteCategoryOK", service.deleteCategory(id));
        return new ModelAndView("redirect:/contents", model);
    }

    @PostMapping("/s/{id}")
    public ModelAndView deleteSerial(@PathVariable Long id, ModelMap model) {
        model.addAttribute("deleteSerialOK", service.deleteSerial(id));
        return new ModelAndView("redirect:/contents", model);
    }
}
