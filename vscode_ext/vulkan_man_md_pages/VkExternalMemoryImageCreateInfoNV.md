# VkExternalMemoryImageCreateInfoNV(3) Manual Page

## Name

VkExternalMemoryImageCreateInfoNV - Specify that an image may be backed
by external memory



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain includes a `VkExternalMemoryImageCreateInfoNV`
structure, then that structure defines a set of external memory handle
types that **may** be used as backing store for the image.

The `VkExternalMemoryImageCreateInfoNV` structure is defined as:

``` c
// Provided by VK_NV_external_memory
typedef struct VkExternalMemoryImageCreateInfoNV {
    VkStructureType                      sType;
    const void*                          pNext;
    VkExternalMemoryHandleTypeFlagsNV    handleTypes;
} VkExternalMemoryImageCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `handleTypes` is zero or a bitmask of
  [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html)
  specifying one or more external memory handle types.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkExternalMemoryImageCreateInfoNV-sType-sType"
  id="VUID-VkExternalMemoryImageCreateInfoNV-sType-sType"></a>
  VUID-VkExternalMemoryImageCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_IMAGE_CREATE_INFO_NV`

- <a href="#VUID-VkExternalMemoryImageCreateInfoNV-handleTypes-parameter"
  id="VUID-VkExternalMemoryImageCreateInfoNV-handleTypes-parameter"></a>
  VUID-VkExternalMemoryImageCreateInfoNV-handleTypes-parameter  
  `handleTypes` **must** be a valid combination of
  [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_external_memory.html),
[VkExternalMemoryHandleTypeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagsNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExternalMemoryImageCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
