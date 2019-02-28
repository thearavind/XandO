package game

import (
	"context"
	fmt "fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/pkg/errors"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	Broadcast     chan GameData
	ClientStreams map[string]chan GameData
	streamsMtx    sync.RWMutex
	openGames     []*GameData
}

func Server() *server {
	return &server{
		Broadcast:     make(chan GameData, 1000),
		ClientStreams: make(map[string]chan GameData),
		openGames:     []*GameData{},
	}
}

func (s *server) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	srv := grpc.NewServer()
	RegisterGameServiceServer(srv, s)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", os.Getenv("PORT")))
	if err != nil {
		return errors.WithMessage(err,
			"server unable to bind on provided host")
	}

	go s.broadcast(ctx)

	go func() {
		fmt.Printf("Inside go routine")
		srv.Serve(l)
		cancel()
	}()

	<-ctx.Done()

	close(s.Broadcast)

	srv.GracefulStop()
	return nil
}

func (s *server) RequestGameStart(ctx context.Context, gameRequest *GameRequestParams) (*GameData, error) {
	s.openStream(gameRequest.UserId)
	fmt.Printf("Got a request %s", gameRequest.UserId)
	if len(s.openGames) > 0 {
		s.streamsMtx.Lock()
		selectedGame := s.openGames[:1][0]
		s.openGames = s.openGames[1:]
		s.streamsMtx.Unlock()
		if selectedGame.Symbol == "x" {
			selectedGame.Symbol = "o"
		} else {
			selectedGame.Symbol = "x"
		}
		selectedGame.Players = append(selectedGame.Players, gameRequest.UserId)
		return selectedGame, nil
	}

	newGame := &GameData{
		GameId:    "1",
		Symbol:    "x",
		GameState: []string{"", "", "", "", "", "", "", "", ""},
		Turn:      "x",
		Result:    "",
		Players:   []string{gameRequest.UserId},
		Winner:    "",
	}
	s.openGames = append(s.openGames, newGame)
	return newGame, nil
}

func (s *server) UpdateGameState(stream GameService_UpdateGameStateServer) error {
	var userId string
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	}
	if t, ok := md["userid"]; ok {
		userId = strings.Join(t, "")
	} else {
		fmt.Print("No metadata")
	}
	go s.sendBroadcasts(stream, userId)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		s.Broadcast <- *in
	}
	<-stream.Context().Done()
	return stream.Context().Err()
}

func (s *server) sendBroadcasts(srv GameService_UpdateGameStateServer, id string) {
	stream := s.openStream(id)
	defer s.closeStream(id)

	for {
		select {
		case <-srv.Context().Done():
			return
		case res := <-stream:
			srv.Send(&res)
		}
	}
}

func (s *server) openStream(id string) (stream chan GameData) {
	stream = make(chan GameData, 100)

	s.streamsMtx.Lock()
	s.ClientStreams[id] = stream
	s.streamsMtx.Unlock()

	fmt.Printf("opened stream for client %s", id)

	return
}

func (s *server) closeStream(id string) {
	s.streamsMtx.Lock()

	if stream, ok := s.ClientStreams[id]; ok {
		delete(s.ClientStreams, id)
		close(stream)
	}

	s.streamsMtx.Unlock()
}

func (s *server) broadcast(ctx context.Context) {
	for res := range s.Broadcast {
		s.streamsMtx.RLock()
		for _, stream := range s.ClientStreams {
			select {
			case stream <- res:
				// noop
			default:
				// no idea
			}
		}
		s.streamsMtx.RUnlock()
	}
}
