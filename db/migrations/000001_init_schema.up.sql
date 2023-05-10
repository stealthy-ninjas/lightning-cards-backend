create extension "uuid-ossp";
-- public.cards definition
-- Drop table
-- DROP TABLE public.cards;
create table public.cards (
                              value varchar not null,
                              constraint cards_pk primary key (value)
);
-- public.center_deck definition
-- Drop table
-- DROP TABLE public.center_deck;
create table public.center_deck (
                                    id uuid not null default uuid_generate_v4(),
                                    constraint center_deck_pk primary key (id)
);

-- public.gestures definition
-- Drop table
-- DROP TABLE public.gestures;



create table public.gestures (
                                 id uuid not null default uuid_generate_v4(),
                                 value varchar not null,
                                 constraint gestures_pk primary key (id)
);


-- public.center_deck_cards definition


-- Drop table


-- DROP TABLE public.center_deck_cards;



create table public.center_deck_cards (
                                          center_deck_id uuid not null,
                                          card_value varchar null,
                                          id uuid not null default uuid_generate_v4(),
                                          constraint center_deck_cards_pk primary key (id),
                                          constraint center_deck_cards_fk foreign key (center_deck_id) references public.center_deck(id),
                                          constraint center_deck_cards_fk_1 foreign key (card_value) references public.cards(value)
);


-- public.player_cards definition


-- Drop table


-- DROP TABLE public.player_cards;



create table public.player_cards (
                                     player_id uuid not null,
                                     card_value varchar not null,
                                     id uuid not null default uuid_generate_v4(),
                                     constraint player_cards_pk primary key (id)
);


-- public.players definition


-- Drop table


-- DROP TABLE public.players;



create table public.players (
                                id uuid not null default uuid_generate_v4(),
                                room_id uuid null,
                                username varchar not null,
                                ready bool not null default false,
                                constraint players_pk primary key (id)
);


-- public.room_gestures definition


-- Drop table


-- DROP TABLE public.room_gestures;



create table public.room_gestures (
                                      room_id uuid not null,
                                      card_id varchar not null,
                                      gesture_id uuid not null,
                                      constraint room_gestures_pk primary key (room_id,
                                                                               card_id)
);


-- public.rooms definition


-- Drop table


-- DROP TABLE public.rooms;



create table public.rooms (
                              id uuid not null default uuid_generate_v4(),
                              center_deck_id uuid null,
                              game_status bool not null default false,
                              host uuid not null,
                              turn uuid null,
                              constraint rooms_pk primary key (id)
);


-- public.rounds definition


-- Drop table


-- DROP TABLE public.rounds;



create table public.rounds (
                               room_id uuid not null,
                               loser uuid null,
                               constraint rounds_pk primary key (room_id)
);

create unique index rounds_un on
    public.rounds
    using btree (room_id);


-- public.winners definition


-- Drop table


-- DROP TABLE public.winners;



create table public.winners (
                                room_id uuid not null,
                                first_rank uuid null,
                                second_rank uuid null,
                                third_rank uuid null,
                                constraint winners_check check (((first_rank <> second_rank)
                                    and (first_rank <> third_rank)
                                    and (second_rank <> third_rank))),
                                constraint winners_pk primary key (room_id)
);


-- public.player_cards foreign keys



alter table public.player_cards add constraint player_cards_fk foreign key (player_id) references public.players(id);

alter table public.player_cards add constraint player_cards_value_fk foreign key (card_value) references public.cards(value);


-- public.players foreign keys



alter table public.players add constraint players_fk foreign key (room_id) references public.rooms(id) on
    delete
    set
    null;


-- public.room_gestures foreign keys



alter table public.room_gestures add constraint room_gestures_fk foreign key (room_id) references public.rooms(id) on
    delete
    cascade on
    update
    cascade;

alter table public.room_gestures add constraint room_gestures_fk_1 foreign key (card_id) references public.cards(value) on
    delete
    cascade on
    update
    cascade;

alter table public.room_gestures add constraint room_gestures_fk_2 foreign key (gesture_id) references public.gestures(id) on
    delete
    cascade;


-- public.rooms foreign keys



alter table public.rooms add constraint rooms_center_deck_fk foreign key (center_deck_id) references public.center_deck(id);

alter table public.rooms add constraint rooms_host_fk foreign key (host) references public.players(id);

alter table public.rooms add constraint rooms_turn_fk foreign key (turn) references public.players(id);


-- public.rounds foreign keys



alter table public.rounds add constraint rounds_fk foreign key (room_id) references public.rooms(id) on
    delete
    cascade on
    update
    cascade;

alter table public.rounds add constraint rounds_fk_1 foreign key (loser) references public.players(id);


-- public.winners foreign keys



alter table public.winners add constraint winners_fk foreign key (room_id) references public.rooms(id) on
    delete
    cascade on
    update
    cascade;

alter table public.winners add constraint winners_fk_1 foreign key (first_rank) references public.players(id);

alter table public.winners add constraint winners_fk_2 foreign key (second_rank) references public.players(id);

alter table public.winners add constraint winners_fk_3 foreign key (third_rank) references public.players(id);