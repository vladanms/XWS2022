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
    private String information;
    private User owner;
    private List<JobPosition> positions;
    private List<JobOffer> offers;

    
    public Company() {}

    public Company(String id, String name, String email, String phoneNumber, String address, String information, User owner) {
        this.id = id;
        this.name = name;
        this.email = email;
        this.phoneNumber = phoneNumber;
        this.address = address;
        this.information = information;
        this.owner = owner;
        this.positions = new ArrayList<JobPosition>();
        this.offers = new ArrayList<JobOffer>();

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

    public String getInformation() {
        return information;
    }

    public void setInformation(String information) {
        this.information = information;
    }

    public User getOwner() {
        return owner;
    }

    public void setOwner(User owner) {
        this.owner = owner;
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