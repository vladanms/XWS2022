package com.dislinkt.agent.service;
import com.dislinkt.agent.model.Interview;
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

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.stereotype.Service;

public class InterviewServiceImpl implements InterviewService {
    
    @Autowired
    private MongoTemplate mongoTemplate;
    @Autowired
    InterviewService service;

    @Override
    public boolean scheduleInterview(Interview interview)
    {
        mongoTemplate.insert(interview);
        return true;
    }

    @Override
    public boolean removeInterview(Interview interview) {
        for(Interview inter : mongoTemplate.findAll(Interview.class)) {
            if(inter.getId().equals(interview.getId())) {
                mongoTemplate.remove(interview);
                return true;
            }
        }
        return false;
    }
}
