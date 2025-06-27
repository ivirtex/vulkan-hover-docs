# VkExternalFormatQNX(3) Manual Page

## Name

VkExternalFormatQNX - Structure containing a QNX Screen buffer external
format



## <a href="#_c_specification" class="anchor"></a>C Specification

To create an image with an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-screen-buffer-external-formats"
target="_blank" rel="noopener">QNX Screen external format</a>, add a
`VkExternalFormatQNX` structure in the `pNext` chain of
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html). `VkExternalFormatQNX` is
defined as:

``` c
// Provided by VK_QNX_external_memory_screen_buffer
typedef struct VkExternalFormatQNX {
    VkStructureType    sType;
    void*              pNext;
    uint64_t           externalFormat;
} VkExternalFormatQNX;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `externalFormat` is an implementation-defined identifier for the
  external format

## <a href="#_description" class="anchor"></a>Description

If `externalFormat` is zero, the effect is as if the
`VkExternalFormatQNX` structure was not present. Otherwise, the `image`
will have the specified external format.

Valid Usage

- <a href="#VUID-VkExternalFormatQNX-externalFormat-08956"
  id="VUID-VkExternalFormatQNX-externalFormat-08956"></a>
  VUID-VkExternalFormatQNX-externalFormat-08956  
  `externalFormat` **must** be `0` or a value returned in the
  `externalFormat` member of
  [VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferFormatPropertiesQNX.html)
  by an earlier call to
  [vkGetScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetScreenBufferPropertiesQNX.html)

Valid Usage (Implicit)

- <a href="#VUID-VkExternalFormatQNX-sType-sType"
  id="VUID-VkExternalFormatQNX-sType-sType"></a>
  VUID-VkExternalFormatQNX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_FORMAT_QNX`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QNX_external_memory_screen_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QNX_external_memory_screen_buffer.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExternalFormatQNX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
