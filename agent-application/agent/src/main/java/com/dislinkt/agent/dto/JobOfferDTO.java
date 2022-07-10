package com.dislinkt.agent.dto;

import com.dislinkt.agent.model.Comment;
import com.dislinkt.agent.model.Company;
import com.dislinkt.agent.model.JobPosition;
import com.dislinkt.agent.model.User;

import javax.validation.constraints.NotNull;
import java.util.List;

public class JobOfferDTO {

    private String id;

    @NotNull
    private Company company;
    @NotNull
    private JobPosition position;
    private List<Comment> comments;
    private List<User> applications;
    private boolean postToDislinkt;

    public JobOfferDTO(String id, Company company, JobPosition position, List<Comment> comments, List<User> applications, boolean postToDislinkt) {
        this.id = id;
        this.company = company;
        this.position = position;
        this.comments = comments;
        this.applications = applications;
        this.postToDislinkt = postToDislinkt;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public Company getCompany() {
        return company;
    }

    public void setCompany(Company company) {
        this.company = company;
    }

    public JobPosition getPosition() {
        return position;
    }

    public void setPosition(JobPosition position) {
        this.position = position;
    }

    public List<Comment> getComments() {
        return comments;
    }

    public void setComments(List<Comment> comments) {
        this.comments = comments;
    }

    public List<User> getApplications() {
        return applications;
    }

    public void setApplications(List<User> applications) {
        this.applications = applications;
    }

    public boolean isPostToDislinkt() {
        return postToDislinkt;
    }

    public void setPostToDislinkt(boolean postToDislinkt) {
        this.postToDislinkt = postToDislinkt;
    }
}
