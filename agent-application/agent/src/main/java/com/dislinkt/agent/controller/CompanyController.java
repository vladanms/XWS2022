package com.dislinkt.agent.controller;

import com.dislinkt.agent.dto.CompanyDTO;
import com.dislinkt.agent.model.Company;
import com.dislinkt.agent.service.CompanyService;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;

@RestController
@RequestMapping("/api/lt")
public class CompanyController {
    @Autowired
    private CompanyService companyService;

    @GetMapping ("/companies")
    public ResponseEntity<?> getAllApprovedCompanies() {
        return ResponseEntity.ok().body(companyService.findAllApprovedCompanies());
    }

    @GetMapping ("/company-requests")
    public ResponseEntity<?> getAllCompanyRequests() {
        return ResponseEntity.ok().body(companyService.findAllUnapprovedCompanies());
    }

    @PostMapping("/companies")
    @PreAuthorize("hasRole('REGULAR_USER')")
    public ResponseEntity<?> createCompanyRequest(@Valid @RequestBody CompanyDTO companyDTO) {
        Company company = new Company();
        BeanUtils.copyProperties(companyDTO, company);
        companyService.addCompany(company);
        return ResponseEntity.ok().build();
    }

    @PostMapping("/company-requests/approve")
    @PreAuthorize("hasRole('ADMINISTRATOR')")
    public ResponseEntity<?> approveCompanyRequest(@RequestBody CompanyDTO companyDTO) {
        Company company = new Company();
        BeanUtils.copyProperties(companyDTO, company);
        companyService.approveCompanyRequest(company);
        return ResponseEntity.ok().build();
    }

    @PostMapping("/company-requests/reject")
    @PreAuthorize("hasRole('ADMINISTRATOR')")
    public ResponseEntity<?> rejectCompanyRequest(@RequestBody CompanyDTO companyDTO) {
        Company company = new Company();
        BeanUtils.copyProperties(companyDTO, company);
        companyService.approveCompanyRequest(company);
        return ResponseEntity.ok().build();
    }

    @PostMapping("/company")
    @PreAuthorize("hasRole('COMPANY_OWNER')")
    public ResponseEntity<?> updateCompanyDetails(@RequestBody CompanyDTO companyDTO) {
        Company company = new Company();
        BeanUtils.copyProperties(companyDTO, company);
        companyService.updateCompany(company);
        return ResponseEntity.ok().build();
    }
}
