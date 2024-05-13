# VkImageConstraintsInfoFlagBitsFUCHSIA(3) Manual Page

## Name

VkImageConstraintsInfoFlagBitsFUCHSIA - Bitmask specifying image
constraints flags



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkImageConstraintsInfoFlagBitsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFlagBitsFUCHSIA.html)::`flags`
include:

``` c
// Provided by VK_FUCHSIA_buffer_collection
typedef enum VkImageConstraintsInfoFlagBitsFUCHSIA {
    VK_IMAGE_CONSTRAINTS_INFO_CPU_READ_RARELY_FUCHSIA = 0x00000001,
    VK_IMAGE_CONSTRAINTS_INFO_CPU_READ_OFTEN_FUCHSIA = 0x00000002,
    VK_IMAGE_CONSTRAINTS_INFO_CPU_WRITE_RARELY_FUCHSIA = 0x00000004,
    VK_IMAGE_CONSTRAINTS_INFO_CPU_WRITE_OFTEN_FUCHSIA = 0x00000008,
    VK_IMAGE_CONSTRAINTS_INFO_PROTECTED_OPTIONAL_FUCHSIA = 0x00000010,
} VkImageConstraintsInfoFlagBitsFUCHSIA;
```

## <a href="#_description" class="anchor"></a>Description

General hints about the type of memory that should be allocated by
Sysmem based on the expected usage of the images in the buffer
collection include:

- `VK_IMAGE_CONSTRAINTS_INFO_CPU_READ_RARELY_FUCHSIA`

- `VK_IMAGE_CONSTRAINTS_INFO_CPU_READ_OFTEN_FUCHSIA`

- `VK_IMAGE_CONSTRAINTS_INFO_CPU_WRITE_RARELY_FUCHSIA`

- `VK_IMAGE_CONSTRAINTS_INFO_CPU_WRITE_OFTEN_FUCHSIA`

For protected memory:

- `VK_IMAGE_CONSTRAINTS_INFO_PROTECTED_OPTIONAL_FUCHSIA` specifies that
  protected memory is optional for the buffer collection.

Note that if all participants in the buffer collection (Vulkan or
otherwise) specify that protected memory is optional, Sysmem will not
allocate protected memory.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_buffer_collection](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html),
[VkImageConstraintsInfoFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFlagsFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageConstraintsInfoFlagBitsFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
