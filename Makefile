
all:
	@echo " * argument required"
	@echo
	@echo "        eg: make proto-base"
	@echo
	@echo " * options:"
	@echo "  - proto-admin"
	@echo "  - proto-base"
	@echo "  - proto-bass"
	@echo "  - proto-dasset"
	@echo "  - proto-dscreen"
	@echo "  - proto-merchant"
	@echo "  - proto-payment"
	@echo "  - proto-rbac"
	@echo "  - proto-user"

######## buf install ###########

buf-init: _buf_update

_buf_update:
	@buf mod prune
	@buf mod update

######## Compile .proto source code ###########

.PHONY: all build bin clean proto-admin proto-base proto-bass proto-dscreen proto-dasset proto-payment proto-rbac proto-user

proto-admin:
	@echo ---------- adminpb ----------
	@buf generate --debug --path pb/adminpb
proto-base:
	@echo ---------- basepb ----------
	@buf generate --debug --path pb/basepb
proto-bass:
	@echo ---------- bass ----------
	@buf generate --debug --path pb/bass
proto-dasset:
	@echo ---------- dassetpb ----------
	@buf generate --debug --path pb/dassetpb
proto-dscreen:
	@echo ---------- user ----------
	@buf generate --debug --path pb/dscreenpb
proto-merchant:
	@echo ---------- merchantpb ----------
	@buf generate --debug --path pb/merchantpb
proto-payment:
	@echo ---------- paymentpb ----------
	@buf generate --debug --path pb/paymentpb
proto-rbac:
	@echo ---------- rbacpb ----------
	@buf generate --debug --path pb/rbacpb
proto-user:
	@echo ---------- user ----------
	@buf generate --debug --path pb/userpb
proto--all:
	@buf generate --debug
proto--all-pub: proto-admin proto-base proto-bass proto-dasset proto-payment proto-rbac proto-user
# proto--all-docs:
# 	@buf generate --debug \
# 		--path adminpb \
# 		--path basepb \
# 		--path bass \
# 		--path dassetpb \
# 		--path merchantpb \
# 		--path paymentpb \
# 		--path rbacpb \
# 		--path user \
