package com.dislinkt.agent.controller;

import com.dislinkt.agent.model.User;
import com.dislinkt.agent.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class UserController {

    @Autowired
    private UserService userService;

    @PostMapping("/register")
    public boolean register(@RequestBody User user) {
        return(userService.register(user));
    }
}
