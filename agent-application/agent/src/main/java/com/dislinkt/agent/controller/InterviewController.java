package com.dislinkt.agent.controller;

import com.dislinkt.agent.service.InterviewService;
import com.dislinkt.agent.dto.InterviewDTO;
import com.dislinkt.agent.model.Interview;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import javax.validation.Valid;
import org.springframework.security.access.prepost.PreAuthorize;

@RestController
public class InterviewController {

    @Autowired
    private InterviewService service;

    @PostMapping("/scheduleInterview")
    @PreAuthorize("hasRole('COMPANY_OWNER')")
    public ResponseEntity schedule(@Valid @RequestBody InterviewDTO interviewDTO)
    {
        Interview interview = new Interview();
        BeanUtils.copyProperties(interviewDTO, interview);
        service.scheduleInterview(interview);
        return ResponseEntity.ok().build();
    }

    @PostMapping("/removeInterview")
    @PreAuthorize("hasRole('COMPANY_OWNER')")
    public ResponseEntity delete(@Valid @RequestBody Interview interview)
    {
        service.removeInterview(interview);
        return ResponseEntity.ok().build();
    }
}
