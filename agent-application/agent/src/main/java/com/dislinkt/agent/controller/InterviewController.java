package com.dislinkt.agent.controller;


import com.dislinkt.agent.service.CompanyService;
import com.dislinkt.agent.dto.CompanyDTO;
import com.dislinkt.agent.dto.InterviewDTO;
import com.dislinkt.agent.model.User;
import com.dislinkt.agent.model.Company;
import com.dislinkt.agent.model.JobOffer;
import com.dislinkt.agent.security.UserDetailsImpl;
import com.dislinkt.agent.security.util.TokenUtil;
import com.dislinkt.agent.service.UserService;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpHeaders;
import org.springframework.http.RequestEntity;
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
import org.springframework.security.access.prepost.PreAuthorize;

@RestController
public class InterviewController {
    
}
