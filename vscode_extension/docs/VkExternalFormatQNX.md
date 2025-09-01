# VkExternalFormatQNX(3) Manual Page

## Name

VkExternalFormatQNX - Structure containing a QNX Screen buffer external format



## [](#_c_specification)C Specification

To create an image with an [QNX Screen external format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-external-screen-buffer-external-formats), add a `VkExternalFormatQNX` structure in the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html). `VkExternalFormatQNX` is defined as:

```c++
// Provided by VK_QNX_external_memory_screen_buffer
typedef struct VkExternalFormatQNX {
    VkStructureType    sType;
    void*              pNext;
    uint64_t           externalFormat;
} VkExternalFormatQNX;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `externalFormat` is an implementation-defined identifier for the external format

## [](#_description)Description

If `externalFormat` is zero, the effect is as if the `VkExternalFormatQNX` structure was not present. Otherwise, the `image` will have the specified external format.

Valid Usage

- [](#VUID-VkExternalFormatQNX-externalFormat-08956)VUID-VkExternalFormatQNX-externalFormat-08956  
  `externalFormat` **must** be `0` or a value returned in the `externalFormat` member of [VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenBufferFormatPropertiesQNX.html) by an earlier call to [vkGetScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetScreenBufferPropertiesQNX.html)

Valid Usage (Implicit)

- [](#VUID-VkExternalFormatQNX-sType-sType)VUID-VkExternalFormatQNX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_FORMAT_QNX`

## [](#_see_also)See Also

[VK\_QNX\_external\_memory\_screen\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QNX_external_memory_screen_buffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalFormatQNX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0