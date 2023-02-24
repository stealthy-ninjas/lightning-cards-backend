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
	id uuid NOT NULL,
	CONSTRAINT center_deck_pk PRIMARY KEY (id)
);


-- public.gestures definition

-- Drop table

-- DROP TABLE public.gestures;

CREATE TABLE public.gestures (
	id uuid NOT NULL,
	value varchar NOT NULL,
	CONSTRAINT gestures_pk PRIMARY KEY (id)
);


-- public.center_deck_cards definition

-- Drop table

-- DROP TABLE public.center_deck_cards;

CREATE TABLE public.center_deck_cards (
	center_deck_id uuid NOT NULL,
	card_value varchar NULL,
	CONSTRAINT center_deck_cards_fk FOREIGN KEY (center_deck_id) REFERENCES public.center_deck(id),
	CONSTRAINT center_deck_cards_fk_1 FOREIGN KEY (card_value) REFERENCES public.cards(value)
);


-- public.rooms definition

-- Drop table

-- DROP TABLE public.rooms;

CREATE TABLE public.rooms (
	id uuid NOT NULL,
	center_deck_id uuid NOT NULL,
	game_status bool NOT NULL,
	CONSTRAINT rooms_pk PRIMARY KEY (id),
	CONSTRAINT rooms_fk FOREIGN KEY (center_deck_id) REFERENCES public.center_deck(id)
);


-- public.players definition

-- Drop table

-- DROP TABLE public.players;

CREATE TABLE public.players (
	id uuid NOT NULL,
	room_id uuid NULL,
	username varchar NOT NULL,
	ready bool NOT NULL DEFAULT false,
	CONSTRAINT players_pk PRIMARY KEY (id),
	CONSTRAINT players_fk FOREIGN KEY (id) REFERENCES public.rooms(id) ON DELETE SET NULL ON UPDATE CASCADE
);


-- public.room_gestures definition

-- Drop table

-- DROP TABLE public.room_gestures;

CREATE TABLE public.room_gestures (
	room_id uuid NOT NULL,
	card_id varchar NOT NULL,
	gesture_id uuid NOT NULL,
	CONSTRAINT room_gestures_pk PRIMARY KEY (room_id, card_id),
	CONSTRAINT room_gestures_fk FOREIGN KEY (room_id) REFERENCES public.rooms(id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT room_gestures_fk_1 FOREIGN KEY (card_id) REFERENCES public.cards(value) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT room_gestures_fk_2 FOREIGN KEY (gesture_id) REFERENCES public.gestures(id) ON DELETE CASCADE
);


-- public.rounds definition

-- Drop table

-- DROP TABLE public.rounds;

CREATE TABLE public.rounds (
	room_id uuid NOT NULL,
	loser uuid NULL,
	CONSTRAINT rounds_un UNIQUE (room_id),
	CONSTRAINT rounds_fk FOREIGN KEY (room_id) REFERENCES public.rooms(id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT rounds_fk_1 FOREIGN KEY (loser) REFERENCES public.players(id)
);


-- public.winners definition

-- Drop table

-- DROP TABLE public.winners;

CREATE TABLE public.winners (
	room_id uuid NOT NULL,
	first_rank uuid NULL,
	second_rank uuid NULL,
	third_rank uuid NULL,
	CONSTRAINT winners_fk FOREIGN KEY (room_id) REFERENCES public.rooms(id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT winners_fk_1 FOREIGN KEY (first_rank) REFERENCES public.players(id),
	CONSTRAINT winners_fk_2 FOREIGN KEY (second_rank) REFERENCES public.players(id),
	CONSTRAINT winners_fk_3 FOREIGN KEY (third_rank) REFERENCES public.players(id)
);


-- public.player_cards definition

-- Drop table

-- DROP TABLE public.player_cards;

CREATE TABLE public.player_cards (
	player_id uuid NOT NULL,
	card_value varchar NOT NULL,
	CONSTRAINT player_cards_fk FOREIGN KEY (player_id) REFERENCES public.players(id),
	CONSTRAINT player_cards_fk_1 FOREIGN KEY (card_value) REFERENCES public.cards(value)
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";