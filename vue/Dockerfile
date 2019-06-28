FROM cander56/vue
USER root
ENV HOMEDIR /home/vue
WORKDIR $HOMEDIR
COPY . .
RUN chown -R vue:vue /home/vue && apt -y install nginx 
USER vue
RUN npm install 
ENTRYPOINT ["./entrypoint.sh"]
