# VkImageConstraintsInfoFlagBitsFUCHSIA(3) Manual Page

## Name

VkImageConstraintsInfoFlagBitsFUCHSIA - Bitmask specifying image constraints flags



## [](#_c_specification)C Specification

Bits which **can** be set in [VkImageConstraintsInfoFlagBitsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFlagBitsFUCHSIA.html)::`flags` include:

```c++
// Provided by VK_FUCHSIA_buffer_collection
typedef enum VkImageConstraintsInfoFlagBitsFUCHSIA {
    VK_IMAGE_CONSTRAINTS_INFO_CPU_READ_RARELY_FUCHSIA = 0x00000001,
    VK_IMAGE_CONSTRAINTS_INFO_CPU_READ_OFTEN_FUCHSIA = 0x00000002,
    VK_IMAGE_CONSTRAINTS_INFO_CPU_WRITE_RARELY_FUCHSIA = 0x00000004,
    VK_IMAGE_CONSTRAINTS_INFO_CPU_WRITE_OFTEN_FUCHSIA = 0x00000008,
    VK_IMAGE_CONSTRAINTS_INFO_PROTECTED_OPTIONAL_FUCHSIA = 0x00000010,
} VkImageConstraintsInfoFlagBitsFUCHSIA;
```

## [](#_description)Description

General hints about the type of memory that should be allocated by Sysmem based on the expected usage of the images in the buffer collection include:

- `VK_IMAGE_CONSTRAINTS_INFO_CPU_READ_RARELY_FUCHSIA`
- `VK_IMAGE_CONSTRAINTS_INFO_CPU_READ_OFTEN_FUCHSIA`
- `VK_IMAGE_CONSTRAINTS_INFO_CPU_WRITE_RARELY_FUCHSIA`
- `VK_IMAGE_CONSTRAINTS_INFO_CPU_WRITE_OFTEN_FUCHSIA`

For protected memory:

- `VK_IMAGE_CONSTRAINTS_INFO_PROTECTED_OPTIONAL_FUCHSIA` specifies that protected memory is optional for the buffer collection.

Note that if all participants in the buffer collection (Vulkan or otherwise) specify that protected memory is optional, Sysmem will not allocate protected memory.

## [](#_see_also)See Also

[VK\_FUCHSIA\_buffer\_collection](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_buffer_collection.html), [VkImageConstraintsInfoFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFlagsFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImageConstraintsInfoFlagBitsFUCHSIA).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0