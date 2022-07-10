package com.dislinkt.agent.security;

import com.dislinkt.agent.security.service.UserDetailsServiceImpl;
import com.dislinkt.agent.security.util.TokenUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.web.authentication.WebAuthenticationDetailsSource;
import org.springframework.web.filter.OncePerRequestFilter;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

public class TokenAuthFilter extends OncePerRequestFilter {
    @Autowired
    private TokenUtil tokenUtil;
    @Autowired
    private UserDetailsServiceImpl userDetailsService;
    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain)
            throws ServletException, IOException {
        try {
            String token = tokenUtil.getTokenFromCookie(request);
            if (token != null) {
                String username = tokenUtil.getUsername(token);
                UserDetails userDetails = userDetailsService.loadUserByUsername(username);
                if (tokenUtil.validateToken(token, userDetails)) {
                    UsernamePasswordAuthenticationToken authentication = new UsernamePasswordAuthenticationToken(userDetails, null, userDetails.getAuthorities());

                    authentication.setDetails(new WebAuthenticationDetailsSource().buildDetails(request));
                    SecurityContextHolder.getContext().setAuthentication(authentication);
                }

            }
        } catch (Exception e) {
            logger.debug("User authentication error", e);
        }
        filterChain.doFilter(request, response);
    }
}
