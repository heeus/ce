/*
 * Copyright (c) 2022-present unTill Pro, Ltd.
 */

package ce

type IServer interface {
	Run() error
}

type Config struct {
	AdminPort int
}
