package com.dislinkt.agent.controller;

import com.dislinkt.agent.dto.UserLoginDTO;
import com.dislinkt.agent.model.User;
import com.dislinkt.agent.dto.UserDTO;
import com.dislinkt.agent.security.UserDetailsImpl;
import com.dislinkt.agent.security.util.TokenUtil;
import com.dislinkt.agent.service.UserService;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpHeaders;
import org.springframework.http.ResponseCookie;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.validation.Valid;

@RestController
@RequestMapping("/api/auth")
public class UserController {
    @Autowired
    private UserService userService;
    @Autowired
    AuthenticationManager authenticationManager;
    @Autowired
    TokenUtil tokenUtil;

    @PostMapping("/register")
    public ResponseEntity<?> register(@Valid @RequestBody UserDTO userDTO) {
        User user = new User();
        BeanUtils.copyProperties(userDTO, user);
        if(!userService.isUsernameUnique(user)) {
            return ResponseEntity.badRequest().body("Error: Username is already in use.");
        } else if(!userService.isEmailUnique(user)) {
            return ResponseEntity.badRequest().body("Error: Email is already in use.");
        }
        userService.addUser(user);
        return ResponseEntity.ok().build();
    }

    @PostMapping("/sign-in")
    public ResponseEntity<?> logIn(@Valid @RequestBody UserLoginDTO dto) {
        Authentication authentication = authenticationManager
                .authenticate(new UsernamePasswordAuthenticationToken(dto.getUsername(), dto.getPassword()));
        SecurityContextHolder.getContext().setAuthentication(authentication);

        UserDetailsImpl userDetails = (UserDetailsImpl) authentication.getPrincipal();
        ResponseCookie cookie = tokenUtil.getCookie(userDetails);
        return ResponseEntity.ok().header(HttpHeaders.SET_COOKIE, cookie.toString())
                .body(userDetails.getUsername());
    }

    @PostMapping("/sign-out")
    public ResponseEntity<?> logOut() {
        ResponseCookie cookie = tokenUtil.getCleanCookie();
        return ResponseEntity.ok().header(HttpHeaders.SET_COOKIE, cookie.toString())
                .body("Signed out successfully.");
    }
}
