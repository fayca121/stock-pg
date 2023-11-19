CREATE TABLE public.stocks (
                              stockid serial NOT NULL,
                              "name" varchar NOT NULL,
                              price real NULL,
                              company varchar NULL,
                              CONSTRAINT stock_pk PRIMARY KEY (stockid)
);