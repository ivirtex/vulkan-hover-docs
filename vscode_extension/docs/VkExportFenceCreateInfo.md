# VkExportFenceCreateInfo(3) Manual Page

## Name

VkExportFenceCreateInfo - Structure specifying handle types that can be exported from a fence



## [](#_c_specification)C Specification

To create a fence whose payload **can** be exported to external handles, add a [VkExportFenceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportFenceCreateInfo.html) structure to the `pNext` chain of the [VkFenceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceCreateInfo.html) structure. The `VkExportFenceCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkExportFenceCreateInfo {
    VkStructureType                   sType;
    const void*                       pNext;
    VkExternalFenceHandleTypeFlags    handleTypes;
} VkExportFenceCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_external_fence
typedef VkExportFenceCreateInfo VkExportFenceCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `handleTypes` is a bitmask of [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html) specifying one or more fence handle types the application **can** export from the resulting fence. The application **can** request multiple handle types for the same fence.

## [](#_description)Description

Valid Usage

- [](#VUID-VkExportFenceCreateInfo-handleTypes-01446)VUID-VkExportFenceCreateInfo-handleTypes-01446  
  The bits in `handleTypes` **must** be supported and compatible, as reported by [VkExternalFenceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceProperties.html)

Valid Usage (Implicit)

- [](#VUID-VkExportFenceCreateInfo-sType-sType)VUID-VkExportFenceCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_FENCE_CREATE_INFO`
- [](#VUID-VkExportFenceCreateInfo-handleTypes-parameter)VUID-VkExportFenceCreateInfo-handleTypes-parameter  
  `handleTypes` **must** be a valid combination of [VkExternalFenceHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlagBits.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkExternalFenceHandleTypeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFenceHandleTypeFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportFenceCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0