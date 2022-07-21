#Dockerfile
FROM microsoft/windowsservercore as builder

WORKDIR "C:/"

COPY "jre-8u181-windows-x64.exe" jre.exe
COPY "apache-activemq-5.15.6-bin.zip" activemq.zip

# Silently install the JRE
RUN	jre.exe /s INSTALLDIR=C:\jre WEB_JAVA=0 SPONSORS=0  && \
	DEL jre.exe

# Extract ActiveMQ
RUN	powershell -Command Expand-Archive activemq.zip C:/  && \
	REN apache-activemq-5.15.6 activemq  && \
	DEL activemq.zip

FROM microsoft/nanoserver
COPY --from=builder ["C:/jre", "C:/jre"]
COPY --from=builder ["C:/activemq", "C:/activemq"]

# Set JAVA_HOME
ENV JAVA_HOME "C:\\jre"
# Add the option to prevent heap reservation issues
ENV _JAVA_OPTIONS "-Xmx128M -Xms128M"
# Add Java to the path
ENV PATH "C:\\jre\\bin;C:\\Windows\\system32;C:\\Windows;C:\\Windows\\System32\\Wbem;C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\"

# Admin Port
EXPOSE 8616
# Listen Port
EXPOSE 61616
# STOMP is the Simple (or Streaming) Text Orientated Messaging Protocol
EXPOSE 61613

CMD "C:\\activemq\\bin\\activemq start"