package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/bwolf1/grpc-rest-kubernetes/pkg/service/echo"
	pb "github.com/bwolf1/grpc-rest-kubernetes/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var grpcServerCmd = &cobra.Command{
	Use:     "grpcServer",
	Aliases: []string{"grpc"},
	Short:   "Run the gRPC server",
	Long: `Echoer
	
	Echoes back to the caller the string it receives.`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.ReadInConfig()

		// Start the server.
		listen, err := net.Listen(
			viper.GetString("grpc.environment.development.network"),
			fmt.Sprintf(
				":%d",
				viper.GetInt64("grpc.environment.development.port")),
		)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterEchoerServer(s, &echo.Server{})
		log.Printf("server listening at %v", listen.Addr())
		if err := s.Serve(listen); err != nil {
			log.Fatalf("server failure: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(grpcServerCmd)
}
