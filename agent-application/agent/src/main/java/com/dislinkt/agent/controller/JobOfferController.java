package com.dislinkt.agent.controller;

import com.dislinkt.agent.dto.JobOfferDTO;
import com.dislinkt.agent.model.JobOffer;
import com.dislinkt.agent.service.CompanyService;
import com.dislinkt.agent.service.JobOfferService;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;

@RestController
public class JobOfferController {

    @Autowired
    private JobOfferService jobOfferService;

    @PostMapping("/job-offer")
    @PreAuthorize("hasRole('COMPANY_OWNER')")
    public ResponseEntity<?> createJobOffer(@Valid @RequestBody JobOfferDTO jobOfferDTO) {
        JobOffer jobOffer = new JobOffer();
        BeanUtils.copyProperties(jobOffer, jobOfferDTO);
        jobOfferService.addJobOffer(jobOffer);
        return ResponseEntity.ok().build();
    }

    @PutMapping("/job-offer")
    @PreAuthorize("hasRole('COMPANY_OWNER')")
    public ResponseEntity<?> removeJobOffer(@Valid @RequestBody JobOfferDTO jobOfferDTO) {
        JobOffer jobOffer = new JobOffer();
        BeanUtils.copyProperties(jobOffer, jobOfferDTO);
        jobOfferService.removeJobOffer(jobOffer);
        return ResponseEntity.ok().build();
    }

    @GetMapping("/job-offers")
    @PreAuthorize("hasRole('COMPANY_OWNER') or hasRole('REGULAR_USER') or hasRole('ADMINISTRATOR')")
    public ResponseEntity<?> getAllJobOffers() {
        return ResponseEntity.ok().body(jobOfferService.findAllJobOffers());
    }
}
