package com.dislinkt.agent.service;
import com.dislinkt.agent.model.Role;
import com.dislinkt.agent.model.User;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.ArrayList;
import java.util.UUID;
import com.dislinkt.agent.model.JobPosition;
import com.dislinkt.agent.model.JobOffer;
import com.dislinkt.agent.model.Company;
import com.dislinkt.agent.model.Interview;
import java.time.LocalDateTime;

@Service
public class CompanyServiceImpl implements CompanyService{

    @Autowired
    private MongoTemplate mongoTemplate;
    @Autowired
    UserService userService;

    @Override
    public Company addCompany(Company company) {
        company.setApproved(false);
        mongoTemplate.insert(company);
        return company;
    }

    @Override
    public List<Company> getCompanies() {
        return mongoTemplate.findAll(Company.class);
    }

    @Override
    public List<Company> getCompaniesByName(String name) {
        List<Company> companies = new ArrayList<Company>();
        for (Company company : getCompanies()) {
            if(company.getName().equals(name))
            {
                companies.add(company);
            }
        }
        return companies;
    }

    @Override
    public JobOffer postOffer(Company company, JobPosition position)
    {
        String id = UUID.randomUUID().toString();
        JobOffer offer = new JobOffer(id, company, position);
        company.addOffer(offer);
        return offer;
    }

    @Override
    public JobPosition addPosition(Company company, String name, Float pay, String description)
    {
         String id = UUID.randomUUID().toString();
         JobPosition position = new JobPosition(id, name, pay, description);
         company.addPosition(position);
         return position;
    }

    @Override
    public List<Company> findAllApprovedCompanies() {
        List<Company> companies = new ArrayList<>();
        for(Company company : mongoTemplate.findAll(Company.class)) {
            if(company.isApproved()) {
                companies.add(company);
            }
        }
        return companies;
    }

    @Override
    public List<Company> findAllUnapprovedCompanies() {
        List<Company> companies = new ArrayList<>();
        for(Company company : mongoTemplate.findAll(Company.class)) {
            if(!company.isApproved()) {
                companies.add(company);
            }
        }
        return companies;
    }

    @Override
    public boolean approveCompanyRequest(Company unverifiedCompany) {
        for(Company company : mongoTemplate.findAll(Company.class)) {
            if(company.getId().equals(unverifiedCompany.getId())) {
                company.setApproved(true);
                mongoTemplate.save(company);

                userService.addRole(company.getOwnerId(), Role.COMPANY_OWNER);
                return true;
            }
        }
        return false;
    }

    @Override
    public boolean removeCompany(Company unverifiedCompany) {
        for(Company company : mongoTemplate.findAll(Company.class)) {
            if(company.getId().equals(unverifiedCompany.getId())) {
                mongoTemplate.remove(company);
                return true;
            }
        }
        return false;
    }

    @Override
    public boolean updateCompany(Company updatedCompany) {
        for(Company company : mongoTemplate.findAll(Company.class)) {
            if(company.getId().equals(updatedCompany.getId())) {
                company.setName(updatedCompany.getName());
                company.setDescription(updatedCompany.getDescription());
                company.setPhoneNumber(updatedCompany.getPhoneNumber());
                company.setAddress(updatedCompany.getAddress());
                mongoTemplate.save(company);

                return true;
            }
        }
        return false;
    }

    @Override
    public boolean apply(User user, JobOffer offer)
    {
        List<User> applications = new ArrayList<User>();
        if(offer.getCompany().getOwnerId().equals((user.getId())))
        {
            //vlasnik ne moze aplicirati sam sebi
            return false;
        }
        else
            applications.add(user);
            offer.setApplications(applications);
            mongoTemplate.save(offer);
            return true;
    }
}
