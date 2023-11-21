package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/petme/api/cmd/serverd/router"
	"github.com/petme/api/internal/repository"
	"github.com/petme/api/pkg/app"
	"github.com/petme/api/pkg/db/pg"
	"github.com/petme/api/pkg/env"
	"github.com/petme/api/pkg/httpserv"
	pkgerrors "github.com/pkg/errors"
)

func main() {
	ctx := context.Background()

	appCfg := app.NewConfigFromEnv()

	if err := appCfg.IsValid(); err != nil {
		log.Fatal(err)
	}

	if err := run(ctx, appCfg); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, appCfg app.Config) error {
	// initialize db connection
	pgConn, err := initPG(ctx, appCfg)
	if err != nil {
		return err
	}
	defer pgConn.Close()

	repo, err := initRepo(pgConn, appCfg)
	if err != nil {
		return err
	}

	rtr, err := initRouter(ctx, appCfg, repo)
	if err != nil {
		return err
	}

	srv := httpserv.NewServer(
		rtr.Handler(),
		httpserv.ServerAddr(env.GetAndValidateF("PORT")),
	)

	return nil
}

func initRepo(pgConn pg.BeginnerExecutor, appCfg app.Config) (repository.Registry, error) {
	return repository.New(pgConn, nil, nil), nil
}

func initRouter(
	ctx context.Context,
	appCfg app.Config,
	repo repository.Registry,
) (router.Router, error) {
	accountCtrl := account.New(
		repo,
		gwys.bifrostGwy,
		gwys.boltGwy,
		gwys.sapGwy,
		gwys.profileGwy,
		gwys.skalboxGwy,
		gwys.ixos2eProxy,
		gwys.yggdrasilGwy,
		account.Config{
			EGiroEligibilityAPIEnabled: env.GetAndValidateF("EGIRO_ELIGIBILITY_API_ENABLED") == "true",
			EGiroPayablesAPIEnabled:    env.GetAndValidateF("EGIRO_PAYABLES_API_ENABLED") == "true",
		},
		env.GetAndValidateF("BILL_DOWNLOAD_URL"),
		env.GetAndValidateF("BILL_ENCRYPTION_KEY"),
	)
	return router.New(
		ctx,
		strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
		os.Getenv("GQL_INTROSPECTION_ENABLED") == "true",
		router.DevToolsConfig{
			BasicAuthRealm: fmt.Sprintf("%s-%s", appCfg.AppName, appCfg.Env.String()),
			BasicAuthUname: os.Getenv("DEV_TOOLS_UNAME"),
			BasicAuthPwd:   os.Getenv("DEV_TOOLS_PWD"),
		},
		iamValidator,
		iamCustomValidator, // Only used by automic njord
		system.New(repo, iamValidator),
		paynow.New(repo, gwys.boltGwy, gwys.odinGwy, gwys.ppmsGwy, accountCtrl),
		sap.New(repo,
			gwys.boltGwy,
			nil, // profile gwy
			nil, // jarvis gwy
			nil, // notify gwy
			event.New(repo, gwys.kafkaGwy),
			sap.Config{
				KafkaTopicAdhocPaymentFileUploaded:     env.GetAndValidateF("SAP_FILE_UPLOADED_KAFKA_TOPIC_ADHOC_PAYMENT"),
				KafkaTopicRecurringPaymentFileUploaded: env.GetAndValidateF("SAP_FILE_UPLOADED_KAFKA_TOPIC_RECURRING_PAYMENT"),
				KafkaTopicRecurringSetupFileUploaded:   env.GetAndValidateF("SAP_FILE_UPLOADED_KAFKA_TOPIC_RECURRING_SETUP"),
			}),
		accountCtrl,
		cardpayCtrl.New(repo, gwys.aresGwy, accountCtrl),
		recurring.New(
			repo,
			gwys.aresGwy,
			accountCtrl,
			gwys.notifyGwy,
			gwys.profileGwy,
			gwys.jarvisGwy,
			nil, // kafkaGwy
			nil, // sapGwy
			nil, // oneUpGwy
			recurring.Config{
				EmailReceiptFromAddress:             env.GetAndValidateF("SENDGRID_FROM_EMAIL"),
				EmailReceiptFromName:                env.GetAndValidateF("SENDGRID_FROM_NAME"),
				EmailTemplateIDRecurringSetup:       env.GetAndValidateF("RECURRING_SETUP_EMAIL_TEMPLATE_ID"),
				EmailTemplateIDRecurringTermination: env.GetAndValidateF("RECURRING_TERMINATION_EMAIL_TEMPLATE_ID"),
			},
		),
		gwys.ixos2eProxy,
		reports.New(repo),
		order.New(
			repo,
			gwys.boltGwy,
			gwys.notifyGwy,
			gwys.profileGwy,
			gwys.jarvisGwy,
			nil, // oneUpGwy
			nil, // odinGwy
			nil, // kafkaGwy,
			nil, // mgcCtrl
			nil, // eventCtrl
			order.Config{
				EmailReceiptFromAddress:                env.GetAndValidateF("SENDGRID_FROM_EMAIL"),
				EmailReceiptFromName:                   env.GetAndValidateF("SENDGRID_FROM_NAME"),
				EmailTemplateIDRecurringTermination:    env.GetAndValidateF("RECURRING_TERMINATION_EMAIL_TEMPLATE_ID"),
				EmailTemplateIDSuccessRecurringPayment: env.GetAndValidateF("CARDRECURRING_SUCCESSFUL_EMAIL_TEMPLATE_ID"),
				EmailTemplateIDSuccessAdhocPayment:     env.GetAndValidateF("CARDADHOC_SUCCESSFUL_EMAIL_TEMPLATE_ID"),
			},
		), unidollar.New(
			repo,
			accountCtrl,
			event.New(repo, gwys.kafkaGwy),
			gwys.sersiGwy,
			nil,
			nil,
			nil,
			"",
			"",
			"",
			env.GetAndValidateF("SERSI_ADHOC_PAYMENT_REQUEST_KAFKA_TOPIC"),
		),
	), nil
}

func initPG(ctx context.Context, appCfg app.Config) (pg.BeginnerExecutor, error) {
	pgOpenConns, err := strconv.Atoi(env.GetAndValidateF("PG_POOL_MAX_OPEN_CONNS"))
	if err != nil {
		return nil, pkgerrors.WithStack(fmt.Errorf("invalid pg poolmax open conns: %w", err))
	}

	pgIdleConns, err := strconv.Atoi(env.GetAndValidateF("PG_POOL_MAX_IDLE_CONNS"))
	if err != nil {
		return nil, pkgerrors.WithStack(fmt.Errorf("invalid pg poolmax idle conns: %w", err))
	}
	var pgOpts []pg.Option
	return pg.NewPool(ctx, appCfg, env.GetAndValidateF("PG_URL"),
		pgOpenConns,
		pgIdleConns,
		pgOpts...,
	)
}
