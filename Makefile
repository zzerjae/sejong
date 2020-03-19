major:
		bumpversion major; cat VERSION

minor:
		bumpversion minor; cat VERSION

patch:
		bumpversion patch; cat VERSION

release:
	bumpversion --tag release; cat VERSION