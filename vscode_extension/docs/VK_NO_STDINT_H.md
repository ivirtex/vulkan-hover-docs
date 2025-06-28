# VK\_NO\_STDINT\_H(3) Manual Page

## Name

VK\_NO\_STDINT\_H - Control definition of types



## [](#_description)Description

If the `VK_NO_STDINT_H` macro is defined by the application at compile time, extended integer types used by the Vulkan API, such as `uint8_t`, **must** also be defined by the application. Otherwise, the Vulkan headers will not compile. If `VK_NO_STDINT_H` is not defined, the system `<stdint.h>` is used to define these types. There is a fallback path when Microsoft Visual Studio version 2008 and earlier versions are detected at compile time.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#boilerplate-platform-specific-header-control)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0