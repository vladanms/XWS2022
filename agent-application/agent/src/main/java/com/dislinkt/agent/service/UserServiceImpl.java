package com.dislinkt.agent.service;


import com.dislinkt.agent.model.Role;
import com.dislinkt.agent.model.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@Service
public class UserServiceImpl implements UserService{

    @Autowired
    private MongoTemplate mongoTemplate;

    @Override
    public boolean addUser(User user) {
        if(isEmailUnique(user) && isUsernameUnique(user)) {
            user.setPassword(new BCryptPasswordEncoder().encode(user.getPassword()));
            List<Role> roles = new ArrayList<>();
            roles.add(Role.REGULAR_USER);
            user.setRoles(roles);
            mongoTemplate.insert(user);
            return true;
        }
        return false;
    }

    @Override
    public List<User> getAllUsers() {
        return mongoTemplate.findAll(User.class);
    }

    @Override
    public User getUserByUsername(String username) {
        for (User user : getAllUsers()) {
            if(user.getUsername().equals(username)) {
                return user;
            }
        }
        return null;
    }

    @Override
    public User getUserByEmail(String email) {
        for (User user : getAllUsers()) {
            if(user.getEmail().equals(email)) {
                return user;
            }
        }
        return null;
    }

    @Override
    public User getUserById(String id) {
        for (User user : getAllUsers()) {
            if(user.getId().equals(id)) {
                return user;
            }
        }
        return null;
    }

    @Override
    public boolean isEmailUnique(User user) {
        if(getUserByEmail(user.getUsername()) != null) {
            return false;
        }
        return true;
    }

    @Override
    public boolean isUsernameUnique(User user) {
        if(getUserByUsername(user.getUsername()) != null) {
            return false;
        }
        return true;
    }

    @Override
    public void addRole(String id, Role role) {
        User owner = mongoTemplate.findById(id, User.class);
        List<Role> roles = owner.getRoles();
        roles.add(role);
        owner.setRoles(roles);
        mongoTemplate.save(owner);
    }
}
