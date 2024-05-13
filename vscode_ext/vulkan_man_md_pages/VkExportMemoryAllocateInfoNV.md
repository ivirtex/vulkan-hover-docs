# VkExportMemoryAllocateInfoNV(3) Manual Page

## Name

VkExportMemoryAllocateInfoNV - Specify memory handle types that may be
exported



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkExportMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfoNV.html)
structure is defined as:

``` c
// Provided by VK_NV_external_memory
typedef struct VkExportMemoryAllocateInfoNV {
    VkStructureType                      sType;
    const void*                          pNext;
    VkExternalMemoryHandleTypeFlagsNV    handleTypes;
} VkExportMemoryAllocateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `handleTypes` is a bitmask of
  [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html)
  specifying one or more memory handle types that **may** be exported.
  Multiple handle types **may** be requested for the same allocation as
  long as they are compatible, as reported by
  [vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html).

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkExportMemoryAllocateInfoNV-sType-sType"
  id="VUID-VkExportMemoryAllocateInfoNV-sType-sType"></a>
  VUID-VkExportMemoryAllocateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_NV`

- <a href="#VUID-VkExportMemoryAllocateInfoNV-handleTypes-parameter"
  id="VUID-VkExportMemoryAllocateInfoNV-handleTypes-parameter"></a>
  VUID-VkExportMemoryAllocateInfoNV-handleTypes-parameter  
  `handleTypes` **must** be a valid combination of
  [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_external_memory.html),
[VkExternalMemoryHandleTypeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagsNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportMemoryAllocateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
