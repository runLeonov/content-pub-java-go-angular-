package com.example.contentpublicationadmin.controller;


import com.example.contentpublicationadmin.entity.Title;
import com.example.contentpublicationadmin.service.ITitleService;
import com.example.contentpublicationadmin.service.impl.TitleService;
import jakarta.websocket.server.PathParam;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/titles")
public class TitleController {

    @Autowired
    private ITitleService service;

    @GetMapping("/{id}")
    public void getTitle(@PathVariable Long id) {
        Title title = service.getTitleById(id);
        return;
    }

    @GetMapping
    public void getTitles() {
        List<Title> titles = service.getAllTitles();
        return;
    }

    @DeleteMapping("/{id}")
    public void deleteTitle(@PathVariable Long id) {
        boolean isOk = service.deleteTitle(id);
        return;
    }
}
