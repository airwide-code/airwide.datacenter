/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package main

import (
	"flag"
	"github.com/airwide-code/airwide.datacenter/access/session/server"
	"github.com/airwide-code/airwide.datacenter/baselib/app"
)

/*
  // Subscriber-related

  def subscriber(consumer: ActorRef): Receive = {
    case OnNext(cmd: SubscribeCommand) ⇒
      cmd match {
        case SubscribeToOnline(userIds) ⇒
          consumer ! UpdatesConsumerMessage.SubscribeToUserPresences(userIds.toSet)
        case SubscribeFromOnline(userIds) ⇒
          consumer ! UpdatesConsumerMessage.UnsubscribeFromUserPresences(userIds.toSet)
        case SubscribeToGroupOnline(groupIds) ⇒
          consumer ! UpdatesConsumerMessage.SubscribeToGroupPresences(groupIds.toSet)
        case SubscribeFromGroupOnline(groupIds) ⇒
          consumer ! UpdatesConsumerMessage.UnsubscribeFromGroupPresences(groupIds.toSet)
        case SubscribeToSeq(_) ⇒
          consumer ! UpdatesConsumerMessage.SubscribeToSeq
        case SubscribeToWeak(Some(group)) ⇒
          consumer ! UpdatesConsumerMessage.SubscribeToWeak(Some(group))
        case SubscribeToWeak(None) ⇒
          log.error("Subscribe to weak is done implicitly on UpdatesConsumer start")
      }
    case OnComplete ⇒
      context.stop(self)
    case OnError(cause) ⇒
      log.error(cause, "Error in upstream")
  }
 */

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func main() {
	flag.Parse()

	instance := server.NewSessionServer("./session.toml")
	app.DoMainAppInstance(instance)
}
