package com.dislinkt.agent.dto;

import javax.validation.constraints.Email;
import javax.validation.constraints.NotBlank;

public class CompanyDTO {
    private String id;
    @NotBlank
    private String name;
    @Email
    private String email;
    @NotBlank
    private String phoneNumber;
    @NotBlank
    private String address;
    private String description;
    private String ownerId;

    public CompanyDTO(String name, String email, String phoneNumber, String address, String description, String ownerId) {
        this.name = name;
        this.email = email;
        this.phoneNumber = phoneNumber;
        this.address = address;
        this.description = description;
        this.ownerId = ownerId;
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

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }
}
