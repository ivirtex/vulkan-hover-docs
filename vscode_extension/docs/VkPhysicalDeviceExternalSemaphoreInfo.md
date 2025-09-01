# VkPhysicalDeviceExternalSemaphoreInfo(3) Manual Page

## Name

VkPhysicalDeviceExternalSemaphoreInfo - Structure specifying semaphore creation parameters.



## [](#_c_specification)C Specification

The `VkPhysicalDeviceExternalSemaphoreInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceExternalSemaphoreInfo {
    VkStructureType                          sType;
    const void*                              pNext;
    VkExternalSemaphoreHandleTypeFlagBits    handleType;
} VkPhysicalDeviceExternalSemaphoreInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_external_semaphore_capabilities
typedef VkPhysicalDeviceExternalSemaphoreInfo VkPhysicalDeviceExternalSemaphoreInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `handleType` is a [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html) value specifying the external semaphore handle type for which capabilities will be returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceExternalSemaphoreInfo-sType-sType)VUID-VkPhysicalDeviceExternalSemaphoreInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_SEMAPHORE_INFO`
- [](#VUID-VkPhysicalDeviceExternalSemaphoreInfo-pNext-pNext)VUID-VkPhysicalDeviceExternalSemaphoreInfo-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of [VkSemaphoreTypeCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreTypeCreateInfo.html)
- [](#VUID-VkPhysicalDeviceExternalSemaphoreInfo-sType-unique)VUID-VkPhysicalDeviceExternalSemaphoreInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkPhysicalDeviceExternalSemaphoreInfo-handleType-parameter)VUID-VkPhysicalDeviceExternalSemaphoreInfo-handleType-parameter  
  `handleType` **must** be a valid [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkExternalSemaphoreHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreHandleTypeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetPhysicalDeviceExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalSemaphoreProperties.html), [vkGetPhysicalDeviceExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalSemaphoreProperties.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceExternalSemaphoreInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0