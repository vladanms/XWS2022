package com.dislinkt.agent.service;

import com.dislinkt.agent.model.Role;
import com.dislinkt.agent.model.User;

import java.util.List;

public interface UserService {
    boolean addUser(User user);
    User getUserByUsername(String username);
    List<User> getAllUsers();
    User getUserByEmail(String email);
    User getUserById(String id);
    boolean isEmailUnique(User user);
    boolean isUsernameUnique(User user);

    void addRole(String id, Role role);
}
