# VkExportSemaphoreCreateInfo(3) Manual Page

## Name

VkExportSemaphoreCreateInfo - Structure specifying handle types that can be exported from a semaphore



## [](#_c_specification)C Specification

To create a semaphore whose payload **can** be exported to external handles, add a [VkExportSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportSemaphoreCreateInfo.html) structure to the `pNext` chain of the [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreCreateInfo.html) structure. The `VkExportSemaphoreCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkExportSemaphoreCreateInfo {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExternalSemaphoreHandleTypeFlags    handleTypes;
} VkExportSemaphoreCreateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_external_semaphore
typedef VkExportSemaphoreCreateInfo VkExportSemaphoreCreateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `handleTypes` is a bitmask of [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html) specifying one or more semaphore handle types the application **can** export from the resulting semaphore. The application **can** request multiple handle types for the same semaphore.

## [](#_description)Description

Valid Usage

- [](#VUID-VkExportSemaphoreCreateInfo-handleTypes-01124)VUID-VkExportSemaphoreCreateInfo-handleTypes-01124  
  The bits in `handleTypes` **must** be supported and compatible, as reported by [VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreProperties.html)

Valid Usage (Implicit)

- [](#VUID-VkExportSemaphoreCreateInfo-sType-sType)VUID-VkExportSemaphoreCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO`
- [](#VUID-VkExportSemaphoreCreateInfo-handleTypes-parameter)VUID-VkExportSemaphoreCreateInfo-handleTypes-parameter  
  `handleTypes` **must** be a valid combination of [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html) values

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkExternalSemaphoreHandleTypeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportSemaphoreCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0