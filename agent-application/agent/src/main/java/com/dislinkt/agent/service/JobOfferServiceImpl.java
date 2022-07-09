package com.dislinkt.agent.service;

import com.dislinkt.agent.model.JobOffer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class JobOfferServiceImpl implements JobOfferService {
    @Autowired
    private MongoTemplate mongoTemplate;
    @Override
    public JobOffer addJobOffer(JobOffer jobOffer) {
        mongoTemplate.save(jobOffer);

        if(jobOffer.isPostToDislinkt()) {
            postJobOfferToDislinkt(jobOffer);
        }

        return jobOffer;
    }

    private void postJobOfferToDislinkt(JobOffer jobOffer) {
    }

    @Override
    public boolean removeJobOffer(JobOffer jobOffer) {
        for(JobOffer offer : mongoTemplate.findAll(JobOffer.class)) {
            if(offer.getId().equals(jobOffer.getId())) {
                mongoTemplate.remove(jobOffer);
                return true;
            }
        }
        return false;
    }

    @Override
    public List<JobOffer> findAllJobOffers() {
        return mongoTemplate.findAll(JobOffer.class);
    }
}
