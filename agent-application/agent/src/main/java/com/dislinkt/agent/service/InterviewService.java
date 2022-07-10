package com.dislinkt.agent.service;

import com.dislinkt.agent.model.Interview;
import com.dislinkt.agent.model.JobOffer;
import com.dislinkt.agent.model.User;

import java.time.LocalDateTime;
import java.util.List;

public interface InterviewService {

    boolean scheduleInterview (Interview interview);

    boolean removeInterview(Interview interview);
    
}
