package com.dislinkt.agent.service;


import com.dislinkt.agent.model.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class UserServiceImpl implements UserService{

    @Autowired
    private MongoTemplate mongoTemplate;

    @Override
    public boolean register(User user) {
        if (getUserByUsername(user.getUsername()) == null && getUserByEmail(user.getEmail()) == null) {
            mongoTemplate.insert(user);
            return true;
        } else {
            return false;
        }
    }

    @Override
    public List<User> getAllUsers() {
        return mongoTemplate.findAll(User.class);
    }

    @Override
    public User getUserByUsername(String username) {
        for (User user : getAllUsers()) {
            if(user.getUsername() == username) {
                return user;
            }
        }
        return null;
    }

    @Override
    public User getUserByEmail(String email) {
        for (User user : getAllUsers()) {
            if(user.getEmail() == email) {
                return user;
            }
        }
        return null;
    }
}
