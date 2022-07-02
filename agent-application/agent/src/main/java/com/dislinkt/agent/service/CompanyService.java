package com.dislinkt.agent.service;

import com.dislinkt.agent.model.JobPosition;
import com.dislinkt.agent.model.JobOffer;
import com.dislinkt.agent.model.Company;

import java.util.List;

public interface CompanyService {
    Company register (Company company);
    JobOffer postOffer(Company company, JobPosition position);
    JobPosition addPosition(Company company, String name, Float pay, String description);   
    List<Company> getCompanies();
    List<Company> getCompaniesByName(String name);

}
