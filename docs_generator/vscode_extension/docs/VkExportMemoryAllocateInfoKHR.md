# VkExportMemoryAllocateInfo(3) Manual Page

## Name

VkExportMemoryAllocateInfo - Specify exportable handle types for a device memory object



## [](#_c_specification)C Specification

When allocating memory whose payload **may** be exported to another process or Vulkan instance, add a [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfo.html) structure to the `pNext` chain of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) structure, specifying the handle types that **may** be exported.

The [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryAllocateInfo.html) structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkExportMemoryAllocateInfo {
    VkStructureType                    sType;
    const void*                        pNext;
    VkExternalMemoryHandleTypeFlags    handleTypes;
} VkExportMemoryAllocateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_external_memory
typedef VkExportMemoryAllocateInfo VkExportMemoryAllocateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `handleTypes` is zero or a bitmask of [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) specifying one or more memory handle types the application **can** export from the resulting allocation. The application **can** request multiple handle types for the same allocation.

## [](#_description)Description

Valid Usage

- [](#VUID-VkExportMemoryAllocateInfo-handleTypes-09860)VUID-VkExportMemoryAllocateInfo-handleTypes-09860  
  The bits in `handleTypes` **must** be supported and compatible, as reported by [VkExternalTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalTensorPropertiesARM.html), [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatProperties.html), or [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalBufferProperties.html)

Valid Usage (Implicit)

- [](#VUID-VkExportMemoryAllocateInfo-sType-sType)VUID-VkExportMemoryAllocateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_MEMORY_ALLOCATE_INFO`
- [](#VUID-VkExportMemoryAllocateInfo-handleTypes-parameter)VUID-VkExportMemoryAllocateInfo-handleTypes-parameter  
  `handleTypes` **must** be a valid combination of [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkExternalMemoryHandleTypeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportMemoryAllocateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0