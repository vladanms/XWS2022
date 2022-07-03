package com.dislinkt.agent.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document
public class JobOffer{

    @Id
    private String id;

    private Company company;
    private JobPosition position;

    public JobOffer(String id, Company company, JobPosition position)
    {
        this.id = id;
        this.company = company;
        this.position = position;
    }

     public String getId()
    {
        return id;
    }
    public Company getCompany() 
    {
        return company;
    }
    public JobPosition getPosition() 
    {
        return position;
    }

    public void setId(String id)
    {
        this.id = id;
    }
    public void setCompany(Company company)
    {
        this.company= company;
    }
    public void setPosition(JobPosition position)
    {
        this.position = position;
    }
    
}