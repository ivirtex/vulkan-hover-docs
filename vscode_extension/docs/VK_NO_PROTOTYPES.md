# VK\_NO\_PROTOTYPES(3) Manual Page

## Name

VK\_NO\_PROTOTYPES - Vulkan header file prototype inclusion control



## [](#_description)Description

If the `VK_NO_PROTOTYPES` macro is defined by an application at compile time, prototypes for Vulkan APIs will not be included. Only typedefs for API function pointers will be defined.

This is intended for applications using their own function loader and dispatch mechanism.

If the macro is not defined by the application, prototypes for Vulkan APIs will be included.

## [](#_see_also)See Also

[VK\_ONLY\_EXPORTED\_PROTOTYPES](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ONLY_EXPORTED_PROTOTYPES.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NO_PROTOTYPES).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0