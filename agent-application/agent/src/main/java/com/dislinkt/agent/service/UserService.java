package com.dislinkt.agent.service;

import com.dislinkt.agent.model.User;

import java.util.List;

public interface UserService {
    boolean register(User user);
    User getUserByUsername(String username);
    List<User> getAllUsers();
    User getUserByEmail(String email);
}
