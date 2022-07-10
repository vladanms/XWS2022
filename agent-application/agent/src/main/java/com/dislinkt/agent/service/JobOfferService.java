package com.dislinkt.agent.service;

import com.dislinkt.agent.model.JobOffer;

import java.util.List;

public interface JobOfferService {
    JobOffer addJobOffer(JobOffer jobOffer);

    boolean removeJobOffer(JobOffer jobOffer);

    boolean updateJobOffer(JobOffer jobOffer);

    List<JobOffer> findAllJobOffers();
}
