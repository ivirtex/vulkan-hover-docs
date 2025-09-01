# VkExportMemoryAllocateInfoNV(3) Manual Page

## Name

VkExportMemoryAllocateInfoNV - Specify memory handle types that may be exported



## [](#_c_specification)C Specification

The [VkExportMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_external_memory
typedef struct VkExportMemoryAllocateInfoNV {
    VkStructureType                      sType;
    const void*                          pNext;
    VkExternalMemoryHandleTypeFlagsNV    handleTypes;
} VkExportMemoryAllocateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `handleTypes` is a bitmask of [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html) specifying one or more memory handle types that **may** be exported. Multiple handle types **may** be requested for the same allocation as long as they are compatible, as reported by [vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html).

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkExportMemoryAllocateInfoNV-sType-sType)VUID-VkExportMemoryAllocateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO_NV`
- [](#VUID-VkExportMemoryAllocateInfoNV-handleTypes-parameter)VUID-VkExportMemoryAllocateInfoNV-handleTypes-parameter  
  `handleTypes` **must** be a valid combination of [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html) values

## [](#_see_also)See Also

[VK\_NV\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_memory.html), [VkExternalMemoryHandleTypeFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagsNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportMemoryAllocateInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0