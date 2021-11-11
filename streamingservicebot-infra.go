package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk"
	ec2 "github.com/aws/aws-cdk-go/awscdk/awsec2"
	ecs "github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type StreamingservicebotInfraStackProps struct {
	awscdk.StackProps
}

func NewStreamingservicebotInfraStack(scope constructs.Construct, id string, props *StreamingservicebotInfraStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Create VPC and Cluster
	vpc := ec2.NewVpc(stack, jsii.String("StreamingServiceBotVpc"), &ec2.VpcProps{
		MaxAzs: jsii.Number(1),
	})

	cluster := ecs.NewCluster(stack, jsii.String("StreamingServiceBotECSCluster"), &ecs.ClusterProps{
		ClusterName: jsii.String("streamingservicebot-cluster"),
		Vpc:         vpc,
	})

	// Create Task Definition
	taskDef := ecs.NewFargateTaskDefinition(stack, jsii.String("StreamingServiceBotTaskDef"), &ecs.FargateTaskDefinitionProps{
		MemoryLimitMiB: jsii.Number(512),
		Cpu:            jsii.Number(256),
		Family:         jsii.String("streamingservicebot-task"),
	})
	container := taskDef.AddContainer(jsii.String("StreamingServiceBotContainer"), &ecs.ContainerDefinitionOptions{
		Environment: &map[string]*string{
			"SSB_REDDIT_ID":             jsii.String(os.Getenv("SSB_REDDIT_ID")),
			"SSB_REDDIT_PASSWORD":       jsii.String(os.Getenv("SSB_REDDIT_PASSWORD")),
			"SSB_REDDIT_SECRET":         jsii.String(os.Getenv("SSB_REDDIT_SECRET")),
			"SSB_REDDIT_SUBREDDITS":     jsii.String("+2000smusic+2010smusic+311+50sMusic+60sMusic+70s+70sMusic+80sHardcorePunk+80sHipHop+80sMusic+80sremixes+90sAlternative+90sHipHop+90sMusic+90sPunk+90sRock+Acappella+acidhouse+AcousticCovers+ADTR+AfricanMusic+afrobeat+AliciaKeys+AltCountry+AlternativeRock+altrap+ambientfolk+ambientmusic+animemusic+AORMelodic+APerfectCircle+ArcadeFire+ArethaFranklin+asianrap+AStateOfTrance+AtmosphericDnB+audioinsurrection+ausmetal+AvengedSevenfold+BABYMETAL+backpacker+backspin+balkanbrass+balkanmusic+Bangtan+baroque+bassheavy+BayRap+BaysideIsACult+Beatles+BestOfDisney+BigBeat+billytalent+BinauralMusic+BlackMetal+Blink182+bluegrass+Blues+bluesrock+BMSR+BoBurnham+Boneyard+boniver+boogiemusic+boomswing+bossanova+BoyBands+brandnew+brazilianmusic+breakbeat+breakcore+britpop+BruceSpringsteen+Burial+CanadianClassicRock+CanadianMusic+carmusic+Catchysongs+ChamberMusic+ChanceTheRapper+ChapHop+cher+chicagohouse+ChiefKeef+ChiefKeef+chillmusic+chillout+chillwave+Chipbreak+Chiptunes+Chopping+choralmusic+Christcore+ChristinaAguilera+CircleMusic+cityandcolour+classicalmusic+ClassicRock+Coldplay+complextro+Complextro+concertband+concerts+contemporary+country+CoverSongs+cpop+CroatianMusic+crunkcore+CutCopy+cxd+cyberpunk_music+DaftPunk+DANCEPARTY+dancepunk+danktunes+darkstep+DavidBowie+Deadmau5+DeathCabforCutie+deathcore+DeathGrips+deathmetal+deepcuts+deephouse+DeepPurple+Deftones+dembow+DieAntwoord+disco+Djent+DMB+DnB+donaldglover+DoomMetal+DreamPop+DrillandBop+Drizzy+Drone+dub+DubStep+EarlyMusic+EarlyMusic+EBM+EDM+electrohiphop+electrohouse+ElectronicBlues+electronicdancemusic+ElectronicJazz+electronicmagic+ElectronicMusic+electropop+electroswing+Elephant6+ElitistClassical+elliegoulding+Eminem+Emo+Emo_Trap+EmoScreamo+empireofthesun+EnterShikari+epicmetal+ETIMusic+Evanescence+Exotica+ExperimentalMusic+FallOutBoy+feedme+FemaleVocalists+festivals+fidget+FilmMusic+filth+findaband+finkworld+FirstAidKit+FitTunes+Flamenco+flaminglips+fleet_foxes+flocked+folk+folkmetal+folkpunk+folkrock+folkunknown+Foreigner+FrankOcean+franzferdinand+FreeAlbums+freemusic+frenchelectro+frenchhouse+Frisson+funkhouse+FunkSouMusic+fusiondancemusic+futurebeats+FutureFunkAirlines+FutureGarage+futuresynth+gabber+gamemusic+gameofbands+GamesMusicMixMash+GaragePunk+GayMusic+germusic+gethightothis+Gfunk+Ghostbc+glitch+glitchop+Gorillaz+GothicMetal+grateful_dead+Greenday+Grime+Grunge+GuiltyPleasureMusic+GunslingerMusic+GunsNRoses+GypsyJazz+happyhardcore+hardcore+hardhouse+HardRock+hardstyle+HeadNodders+heady+HeyThatWasIn+HighFidelity+hiphop101+hiphopheads+hiphopheadsnorthwest+horrorpunk+house+icm+idm+ifyoulikeblank+ilikethissong+Incubus+indie+indie_rock+IndieFolk+Indieheads+IndieWok+industrialmusic+Instrumentals+international_music+ipm+Irishmusic+IsolatedVocals+ItalianMusic+ItaloDisco+JackJohnson+JackJohnson+JackWhite+jambands+JanetJackson+jazz+JazzFusion+JazzInfluence+jpop+jrock+Kanye+KendrickLamar+kings_of_leon+klezmer+KoreanRock+Korn+kpop+krautrock+LadiesofMetal+ladygaga+lanadelrey+latinhouse+Led_Zeppelin+LeeHallMusic+lennykravitz+LetsTalkMusic+LiquidDubstep+listentoconcerts+listentoconcerts+listentodynamic+listentomusic+listentonews+ListenToThis+ListenToUs+livemusic+llawenyddhebddiwedd+LofiHipHop+LongerJams+lorde+lt10k+Lyrics+macdemarco+Macklemore+Madlib+Madonna+mainstreammusic+makemeaplaylist+Manowar+MariahCarey+mashups+MathRock+MattAndKim+MedievalMusic+Megadeth+MelancholyMusic+melodicdeathmetal+melodichouse+MelodicMetal+MemphisRap+metal+metalcore+Metallica+Metalmusic+MetalNews+mfdoom+MGMT+MichaelJackson+MiddleEasternMusic+minimal+minimalism_music+MinusTheBear+mixes+MLPtunes+ModernRockMusic+ModestMouse+monsterfuzz+moombahcore+Morrissey+MoscowBeat+motown+mrbungle+Muse+Music+MusicAlbums+musicanova+MusicForConcentration+musicsuggestions+MusicToSleepTo+musicvideos+muzyka+MyChemicalRomance+NameThatSong+NeilYoung+neopsychedelia+NewAlbums+newmusic+NewWave+nightstep+NIN+Nirvana+NOFX+noiserock+NuDisco+numetal+NYrap+oasis+officialbadcompany+OfMonstersAndMen+ofMontreal+OFWGKTA-OddFuture+oldiemusic+OldiesMusic+OldskoolRave+onealbumaweek+Opera+Opeth+orchestra+OutKast+OutlawCountry+Outrun+panicatthedisco+partymusic+partymusic+PearlJam+phish+pianocovers+pianorock+Pinback+PinkFloyd+plunderphonics+popheads+poppunkers+PoptoRock+porcupinetree+PostHardcore+PostRock+powermetal+powerpop+prettylights+ProgMetal+progrockmusic+psybient+PsyBreaks+PsychedelicRock+psytrance+punk+Punk_Rock+Punkskahardcore+purplemusic+Puscifer+Queen+QuietStorm+Radiohead+raggajungle+rainymood+Rap+raprock+raprock+rapverses+RATM+RealDubstep+recordstorefinds+RedditOriginals+RedHotChiliPeppers+reggae+RelientK+remixxd+RepublicOfMusic+RetroMusic+rhymesandbeats+rhymesandbeats+RiseAgainst+rnb+RnBHeads+robertplant+Rock+RoMusic+rootsmusic+RoyaltyFreeMusic+runningmusic+SalsaMusic+SaveMoneyCrew+ScottishMusic+shoegaze+SigurRos+Ska+skweee+SlavicMusicVideos+Slayer+slipknot+SmashingPumpkins+Soca+SOCOTRAband+somluso+songbooks+songwriterscircle+SoulDivas+Soulies+SoundsVintage+SpaceMusic+SparksFTW+spop+streamingservicebot+stonerrock+SurfPunk+swing+swinghouse+symphonicblackmetal+symphonicmetal+synthrock+TameImpala+Tango+TaylorSwift+tech_house+Techno+TeganAndSara+The_Residents+TheAvettBrothers+TheBeachBoys+TheCure+TheKillers+TheMagneticZeros+ThemVoices+TheOffspring+theOverload+thepixies+ThePolice+TheRealBookVideos+TheStrokes+TheWeeknd+ThrowbackCore+tmbg+TodaysFavoriteSong+ToolBand+TouhouMusic+Tpain+TraditionalMusic+tragicallyhip+Trance+tranceandbass+trap+trapmuzik+treemusic+tribalbeats+triphop+TropicalHouse+Truemetal+truethrash+trunknation+TwentyOnePilots+U2Band+UKbands+ukfunky+ukhiphopheads+Umphreys+undergroundchicago+unheardof+UnicornsMusic+vaporwave+velvetunderground+VintageObscura+vocaloid+Ween+weezer+WeirdAl+WhatIListenTo+witch_house+WomenRock+WorldMusic+WTFMusicVideos+wuuB+yesband+Zappa"),
			"SSB_REDDIT_USERNAME":       jsii.String(os.Getenv("SSB_REDDIT_USERNAME")),
			"SSB_SPOTIFY_CLIENT_ID":     jsii.String(os.Getenv("SSB_SPOTIFY_CLIENT_ID")),
			"SSB_SPOTIFY_CLIENT_SECRET": jsii.String(os.Getenv("SSB_SPOTIFY_CLIENT_SECRET")),
			"SSB_YOUTUBE_KEY":           jsii.String(os.Getenv("SSB_YOUTUBE_KEY")),
		},
		Image: ecs.ContainerImage_FromRegistry(jsii.String("ghcr.io/scottbrenner/streamingservicebot:main"), &ecs.RepositoryImageProps{}),
	})

	container.AddPortMappings(&ecs.PortMapping{
		ContainerPort: jsii.Number(80),
		Protocol:      ecs.Protocol_TCP,
	})

	// Create Fargate Service
	service := ecs.NewFargateService(stack, jsii.String("StreamingServiceBotService"), &ecs.FargateServiceProps{
		Cluster:        cluster,
		ServiceName:    jsii.String("streamingservicebot-service"),
		TaskDefinition: taskDef,
	})

	awscdk.NewCfnOutput(stack, jsii.String("Service"), &awscdk.CfnOutputProps{Value: ecs.FargateService.ServiceName(service)})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewStreamingservicebotInfraStack(app, "StreamingservicebotInfraStack", &StreamingservicebotInfraStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	// return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("AWS_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("AWS_DEFAULT_REGION")),
	}
}
