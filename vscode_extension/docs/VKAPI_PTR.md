# VKAPI\_PTR(3) Manual Page

## Name

VKAPI\_PTR - Vulkan function pointer calling conventions macro



## [](#_description)Description

`VKAPI_PTR` is a macro placed between the '(' and '\*' in Vulkan API function pointer declarations. This macro also controls calling conventions, and typically has the same definition as `VKAPI_ATTR` or `VKAPI_CALL`, depending on the compiler.

## [](#_see_also)See Also

[VKAPI\_ATTR](https://registry.khronos.org/vulkan/specs/latest/man/html/VKAPI_ATTR.html), [VKAPI\_CALL](https://registry.khronos.org/vulkan/specs/latest/man/html/VKAPI_CALL.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#boilerplate-platform-specific-calling-conventions).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0