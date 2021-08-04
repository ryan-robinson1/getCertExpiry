FROM centos/go-toolset-7-centos7

RUN git clone https://github.com/ryan-robinson1/getCertExpiry.git
RUN cd getCertExpiry
RUN go help