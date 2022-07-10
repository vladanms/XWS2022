package com.dislinkt.agent.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document
public class JobPosition{

    @Id
    private String id;

   // private Company company;
    private String name;
    private float pay;
    private String description;

    
    public JobPosition(String id, String name, float pay, String description)
    {
     //   this.company = company;
        this.id = id;
        this.name = name;
        this.pay = pay;
        this.description = description;
    }

    public String getId()
    {
        return id;
    }
    public String getName() 
    {
        return name;
    }
    public float getPay() 
    {
        return pay;
    }
    public String getDescription() 
    {
        return description;
    }

    public void setId(String id)
    {
        this.id = id;
    }
    public void setName(String name)
    {
        this.name = name;
    }
    public void setId(Float pay)
    {
        this.pay = pay;
    }
    public void setDescription(String description)
    {
        this.description = description;
    } 

    /*public String getCompany() 
    {
        return company;
    }
    public void setCompany(Company company)
    {
        this.company= company;
    }*/
}