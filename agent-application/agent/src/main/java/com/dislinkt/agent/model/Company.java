package com.dislinkt.agent.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import java.util.List;
import java.util.ArrayList;


@Document
public class Company {
    @Id
    private String id;

    private String name;
    private String email;
    private String phoneNumber;
    private String address;
    private String description;
    private String ownerId;
    private List<JobPosition> positions;
    private List<JobOffer> offers;
    private boolean approved;

    
    public Company() {}

    public Company(String id, String name, String email, String phoneNumber, String address, String description, String ownerId, boolean approved) {
        this.id = id;
        this.name = name;
        this.email = email;
        this.phoneNumber = phoneNumber;
        this.address = address;
        this.description = description;
        this.ownerId = ownerId;
        this.positions = new ArrayList<JobPosition>();
        this.offers = new ArrayList<JobOffer>();
        this.approved = approved;

    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPhoneNumber() {
        return phoneNumber;
    }

    public void setPhoneNumber(String phoneNumber) {
        this.phoneNumber = phoneNumber;
    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getOwnerId() {
        return ownerId;
    }

    public void setOwnerId(String ownerId) {
        this.ownerId = ownerId;
    }

    public boolean isApproved() {
        return approved;
    }

    public void setApproved(boolean approved) {
        this.approved = approved;
    }

    public List<JobPosition> getPositions() {
        return positions;
    }

    public void setPositions(List<JobPosition> positions) {
        this.positions = positions;
    }

    public List<JobOffer> getOffers() {
        return offers;
    }

    public void setOffers(List<JobOffer> offers) {
        this.offers = offers;
    }

    public void addOffer(JobOffer offer)
    {
        this.offers.add(offer);
    }
    public void addPosition(JobPosition position)
    {
        this.positions.add(position);
    }
}
