package com.dislinkt.agent.service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.ArrayList;
import java.util.UUID;
import com.dislinkt.agent.model.JobPosition;
import com.dislinkt.agent.model.JobOffer;
import com.dislinkt.agent.model.Company;


import org.springframework.stereotype.Service;

@Service
public class CompanyServiceImpl implements CompanyService{

    @Autowired
    private MongoTemplate mongoTemplate;

    @Override
    public Company register(Company company) {
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
}
