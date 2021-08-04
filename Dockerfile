FROM golang:latest

RUN git clone https://github.com/ryan-robinson1/getCertExpiry.git
RUN cd getCertExpiry
RUN go help