package com.dislinkt.agent.dto;

import com.dislinkt.agent.model.JobOffer;
import com.dislinkt.agent.model.User;
import com.dislinkt.agent.model.Interview;

import java.time.LocalDateTime;

public class InterviewDTO {


    private String id;
    private User applicant;
    private JobOffer offer;
    private LocalDateTime time;

    public InterviewDTO(){};

    public InterviewDTO(String id, User user, JobOffer offer, LocalDateTime time)
    {
        this.id = id;
        this.applicant = user;
        this.offer = offer;
        this.time = time;
    }

    public User GetApplicant()
    {
        return applicant;
    }

    public JobOffer GetOffer()
    {
        return offer;
    }

    public LocalDateTime GetTime()
    {
        return time;
    }

    public void setApplicant(User applicant)
    {
        this.applicant = applicant;
    }

    public void setOffer(JobOffer offer)
    {
        this.offer = offer;
    }

    public void setTime(LocalDateTime time)
    {
        this.time = time;
    }

    public void setId(String id)
    {
        this.id = id;;
    }

    public String getId()
    {
        return id;
    }
    
}
