package com.dislinkt.agent.service;

import com.dislinkt.agent.model.JobPosition;
import com.dislinkt.agent.model.JobOffer;
import com.dislinkt.agent.model.Company;
import com.dislinkt.agent.model.User;
import com.dislinkt.agent.model.Interview;

import java.time.LocalDateTime;
import java.util.List;

public interface CompanyService {
    Company addCompany (Company company);
    JobOffer postOffer(Company company, JobPosition position);
    JobPosition addPosition(Company company, String name, Float pay, String description);   
    List<Company> getCompanies();
    List<Company> getCompaniesByName(String name);

    List<Company> findAllApprovedCompanies();

    List<Company> findAllUnapprovedCompanies();

    boolean approveCompanyRequest(Company unverifiedCompany);

    boolean removeCompany(Company unverifiedCompany);

    boolean updateCompany(Company updatedCompany);

    boolean apply(User user, JobOffer offer);

    boolean scheduleInterview (Interview interview);

    boolean removeInterview(Interview interview);
}
