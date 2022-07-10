package com.dislinkt.agent.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import java.util.List;
import java.util.ArrayList;

@Document
public class JobOffer{

    @Id
    private String id;

    private Company company;
    private JobPosition position;
    private List<Comment> comments;
    private List<User> applications;
    private boolean postToDislinkt;

    public JobOffer() {}

    public JobOffer(String id, Company company, JobPosition position)
    {
        this.id = id;
        this.company = company;
        this.position = position;
        comments = new ArrayList<Comment>();
        applications = new ArrayList<User>();
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
    public void setComments(List<Comment> comments)
    {
        this.comments = comments;
    }
    public List<Comment> getComments()
    {
        return comments;
    }
    public void setApplications(List<User> applications)
    {
        this.applications = applications;
    }
    public List<User> getApplications()
    {
        return applications;
    }

    public boolean isPostToDislinkt() {
        return postToDislinkt;
    }

    public void setPostToDislinkt(boolean postToDislinkt) {
        this.postToDislinkt = postToDislinkt;
    }
}