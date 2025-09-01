# VK\_ONLY\_EXPORTED\_PROTOTYPES(3) Manual Page

## Name

VK\_ONLY\_EXPORTED\_PROTOTYPES - Vulkan header file exported prototype inclusion control



## [](#_description)Description

If the `VK_ONLY_EXPORTED_PROTOTYPES` macro is defined by an application at compile time, only prototypes for Vulkan APIs tagged as \`"exported"\`in the API XML will be included. For non-tagged APIs, only typedefs for API function pointers will be defined.

This is intended to match APIs which are statically exported by the Vulkan loader. At present, the exported APIs are only those defined by Vulkan core versions.

If the macro is not defined by the application, prototypes for all Vulkan APIs will be included.

## [](#_see_also)See Also

[VK\_NO\_PROTOTYPES](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NO_PROTOTYPES.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_ONLY_EXPORTED_PROTOTYPES).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0