package com.example.contentpublicationadmin.controller;


import com.example.contentpublicationadmin.entity.Sort;
import com.example.contentpublicationadmin.entity.Title;
import com.example.contentpublicationadmin.service.ITitleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.ModelAndView;

import java.util.List;

@Controller
@RequestMapping("/titles")
public class TitleController {

    @Autowired
    private ITitleService service;

    @GetMapping("/{id}")
    public ModelAndView getTitle(@PathVariable Long id, ModelMap model) {
        Title title = service.getTitleById(id);
        model.addAttribute("title", title);
        return new ModelAndView("title", model);
    }

    @PostMapping("/{id}/release")
    public ModelAndView releaseOrBan(@PathVariable Long id, ModelMap model) {
        Title title = service.getTitleById(id);
        model.addAttribute("title", service.releaseOrBanTitleById(title));
        return new ModelAndView("title", model);
    }

    @PostMapping("/{id}/comment/{commentId}")
    public ModelAndView deleteComment(@PathVariable Long id, @PathVariable Long commentId, ModelMap model) {
        Title title = service.getTitleById(id);
        model.addAttribute("title", service.deleteCommentForTitle(title, commentId));
        return new ModelAndView("title", model);
    }

    @GetMapping("/")
    public ModelAndView titles(ModelMap model) {
        List<Title> titles = service.getAllTitles();
        model.addAttribute("titles", titles);
        return new ModelAndView("titles", model);
    }

    @GetMapping("/f")
    public ModelAndView filterTitles(@RequestParam(name = "filter", required = false) String filter, ModelMap model) {
        filter = filter.trim().toLowerCase();
        if (filter.isEmpty()) {
            model.addAttribute("titles", service.getAllTitles());
            return new ModelAndView("titles", model);
        }
        List<Title> titles = service.getAllFilteredTitles(filter);
        model.addAttribute("titles", titles);
        return new ModelAndView("titles", model);
    }

    @GetMapping("/sort/id/")
    public ModelAndView sortByID(ModelMap model) {
        model.addAttribute("titles", service.getAllTitlesSort(Sort.ID));
        model.addAttribute("byID", true);
        return new ModelAndView("titles", model);
    }

    @GetMapping("/sort/date/")
    public ModelAndView sortByDate(ModelMap model) {
        model.addAttribute("titles", service.getAllTitlesSort(Sort.DATE));
        model.addAttribute("byDate", true);
        return new ModelAndView("titles", model);
    }

    @DeleteMapping("/{id}")
    public ModelAndView deleteTitle(@PathVariable Long id, ModelMap model) {
        boolean isOk = service.deleteTitle(id);
        return new ModelAndView("titles", model);
    }
}
