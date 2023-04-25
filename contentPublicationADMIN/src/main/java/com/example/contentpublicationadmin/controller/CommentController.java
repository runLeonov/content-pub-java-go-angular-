package com.example.contentpublicationadmin.controller;


import com.example.contentpublicationadmin.entity.AllContent;
import com.example.contentpublicationadmin.entity.Comment;
import com.example.contentpublicationadmin.entity.Sort;
import com.example.contentpublicationadmin.entity.Title;
import com.example.contentpublicationadmin.service.ICommentService;
import com.example.contentpublicationadmin.service.IContentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.ModelAndView;

import java.util.List;

@Controller
@RequestMapping("/comments")
public class CommentController {

    @Autowired
    private ICommentService service;

    @GetMapping("/")
    public ModelAndView comments(ModelMap model) {
        List<Comment> comments = service.getAllComments();
        model.addAttribute("comments", comments);
        return new ModelAndView("comments", model);
    }


    @GetMapping("/f")
    public ModelAndView filterBy(@RequestParam(name = "filter", required = false) String filter, ModelMap model) {
        filter = filter.trim().toLowerCase();
        if (filter.isEmpty()) {
            model.addAttribute("comments", service.getAllComments());
            return new ModelAndView("comments", model);
        }
        List<Comment> comments = service.getAllFilteredComments(filter);
        model.addAttribute("comments", comments);
        return new ModelAndView("comments", model);
    }


    @GetMapping("/sort/id/")
    public ModelAndView sortByID(ModelMap model) {
        model.addAttribute("comments", service.getAllCommentsSorted(Sort.ID));
        model.addAttribute("byID", true);
        return new ModelAndView("comments", model);
    }

    @GetMapping("/sort/date/")
    public ModelAndView sortByDate(ModelMap model) {
        model.addAttribute("comments", service.getAllCommentsSorted(Sort.DATE));
        model.addAttribute("byDate", true);
        return new ModelAndView("comments", model);
    }


    @PostMapping("/{id}/del/")
    public ModelAndView deleteComment(@PathVariable Long id, ModelMap model) {
        model.addAttribute("delIsOK", service.deleteComment(id));
        return new ModelAndView("redirect:/comments", model);
    }

    @GetMapping("")
    public ModelAndView redirect(ModelMap model) {
        return comments(model);
    }
}
