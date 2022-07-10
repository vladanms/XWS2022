package com.dislinkt.agent.security.service;

import com.dislinkt.agent.model.User;
import com.dislinkt.agent.security.UserDetailsImpl;
import com.dislinkt.agent.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

@Service
public class UserDetailsServiceImpl implements UserDetailsService {

    @Autowired
    private UserService userService;

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        User user = userService.getUserByUsername(username);
        if(user != null) {
            return UserDetailsImpl.getUserDetailsFromUser(user);
        }
        return null;
    }
}
