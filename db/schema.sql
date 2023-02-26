-- public.cards definition

-- Drop table

-- DROP TABLE public.cards;

CREATE TABLE public.cards (
	value varchar NOT NULL,
	CONSTRAINT cards_pk PRIMARY KEY (value)
);


-- public.center_deck definition

-- Drop table

-- DROP TABLE public.center_deck;

CREATE TABLE public.center_deck (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	CONSTRAINT center_deck_pk PRIMARY KEY (id)
);


-- public.gestures definition

-- Drop table

-- DROP TABLE public.gestures;

CREATE TABLE public.gestures (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	value varchar NOT NULL,
	CONSTRAINT gestures_pk PRIMARY KEY (id)
);


-- public.center_deck_cards definition

-- Drop table

-- DROP TABLE public.center_deck_cards;

CREATE TABLE public.center_deck_cards (
	center_deck_id uuid NOT NULL,
	card_value varchar NULL,
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	CONSTRAINT center_deck_cards_pk PRIMARY KEY (id),
	CONSTRAINT center_deck_cards_fk FOREIGN KEY (center_deck_id) REFERENCES public.center_deck(id),
	CONSTRAINT center_deck_cards_fk_1 FOREIGN KEY (card_value) REFERENCES public.cards(value)
);


-- public.player_cards definition

-- Drop table

-- DROP TABLE public.player_cards;

CREATE TABLE public.player_cards (
	player_id uuid NOT NULL,
	card_value varchar NOT NULL,
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	CONSTRAINT player_cards_pk PRIMARY KEY (id)
);


-- public.players definition

-- Drop table

-- DROP TABLE public.players;

CREATE TABLE public.players (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	room_id uuid NULL,
	username varchar NOT NULL,
	ready bool NOT NULL DEFAULT false,
	CONSTRAINT players_pk PRIMARY KEY (id)
);


-- public.room_gestures definition

-- Drop table

-- DROP TABLE public.room_gestures;

CREATE TABLE public.room_gestures (
	room_id uuid NOT NULL,
	card_id varchar NOT NULL,
	gesture_id uuid NOT NULL,
	CONSTRAINT room_gestures_pk PRIMARY KEY (room_id, card_id)
);


-- public.rooms definition

-- Drop table

-- DROP TABLE public.rooms;

CREATE TABLE public.rooms (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	center_deck_id uuid NULL,
	game_status bool NOT NULL DEFAULT false,
	host uuid NOT NULL,
	turn uuid NULL,
	CONSTRAINT rooms_pk PRIMARY KEY (id),
	CONSTRAINT rooms_un UNIQUE (host)
);


-- public.rounds definition

-- Drop table

-- DROP TABLE public.rounds;

CREATE TABLE public.rounds (
	room_id uuid NOT NULL,
	loser uuid NULL,
	CONSTRAINT rounds_pk PRIMARY KEY (room_id)
);
CREATE UNIQUE INDEX rounds_un ON public.rounds USING btree (room_id);


-- public.winners definition

-- Drop table

-- DROP TABLE public.winners;

CREATE TABLE public.winners (
	room_id uuid NOT NULL,
	first_rank uuid NULL,
	second_rank uuid NULL,
	third_rank uuid NULL,
	CONSTRAINT winners_check CHECK (((first_rank <> second_rank) AND (first_rank <> third_rank) AND (second_rank <> third_rank))),
	CONSTRAINT winners_pk PRIMARY KEY (room_id)
);


-- public.player_cards foreign keys

ALTER TABLE public.player_cards ADD CONSTRAINT player_cards_fk FOREIGN KEY (player_id) REFERENCES public.players(id);
ALTER TABLE public.player_cards ADD CONSTRAINT player_cards_value_fk FOREIGN KEY (card_value) REFERENCES public.cards(value);


-- public.players foreign keys

ALTER TABLE public.players ADD CONSTRAINT players_fk FOREIGN KEY (room_id) REFERENCES public.rooms(id) ON DELETE SET NULL;


-- public.room_gestures foreign keys

ALTER TABLE public.room_gestures ADD CONSTRAINT room_gestures_fk FOREIGN KEY (room_id) REFERENCES public.rooms(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE public.room_gestures ADD CONSTRAINT room_gestures_fk_1 FOREIGN KEY (card_id) REFERENCES public.cards(value) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE public.room_gestures ADD CONSTRAINT room_gestures_fk_2 FOREIGN KEY (gesture_id) REFERENCES public.gestures(id) ON DELETE CASCADE;


-- public.rooms foreign keys

ALTER TABLE public.rooms ADD CONSTRAINT rooms_center_deck_fk FOREIGN KEY (center_deck_id) REFERENCES public.center_deck(id);
ALTER TABLE public.rooms ADD CONSTRAINT rooms_host_fk FOREIGN KEY (host) REFERENCES public.players(id);
ALTER TABLE public.rooms ADD CONSTRAINT rooms_turn_fk FOREIGN KEY (turn) REFERENCES public.players(id);


-- public.rounds foreign keys

ALTER TABLE public.rounds ADD CONSTRAINT rounds_fk FOREIGN KEY (room_id) REFERENCES public.rooms(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE public.rounds ADD CONSTRAINT rounds_fk_1 FOREIGN KEY (loser) REFERENCES public.players(id);


-- public.winners foreign keys

ALTER TABLE public.winners ADD CONSTRAINT winners_fk FOREIGN KEY (room_id) REFERENCES public.rooms(id) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE public.winners ADD CONSTRAINT winners_fk_1 FOREIGN KEY (first_rank) REFERENCES public.players(id);
ALTER TABLE public.winners ADD CONSTRAINT winners_fk_2 FOREIGN KEY (second_rank) REFERENCES public.players(id);
ALTER TABLE public.winners ADD CONSTRAINT winners_fk_3 FOREIGN KEY (third_rank) REFERENCES public.players(id);