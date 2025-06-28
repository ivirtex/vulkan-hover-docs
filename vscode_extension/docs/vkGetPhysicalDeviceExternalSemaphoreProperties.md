# vkGetPhysicalDeviceExternalSemaphoreProperties(3) Manual Page

## Name

vkGetPhysicalDeviceExternalSemaphoreProperties - Function for querying external semaphore handle capabilities.



## [](#_c_specification)C Specification

Semaphores **may** support import and export of their [payload](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-payloads) to external handles. To query the external handle types supported by semaphores, call:

```c++
// Provided by VK_VERSION_1_1
void vkGetPhysicalDeviceExternalSemaphoreProperties(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceExternalSemaphoreInfo* pExternalSemaphoreInfo,
    VkExternalSemaphoreProperties*              pExternalSemaphoreProperties);
```

or the equivalent command

```c++
// Provided by VK_KHR_external_semaphore_capabilities
void vkGetPhysicalDeviceExternalSemaphorePropertiesKHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceExternalSemaphoreInfo* pExternalSemaphoreInfo,
    VkExternalSemaphoreProperties*              pExternalSemaphoreProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the semaphore capabilities.
- `pExternalSemaphoreInfo` is a pointer to a [VkPhysicalDeviceExternalSemaphoreInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html) structure describing the parameters that would be consumed by [vkCreateSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateSemaphore.html).
- `pExternalSemaphoreProperties` is a pointer to a [VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreProperties.html) structure in which capabilities are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-physicalDevice-parameter)VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-pExternalSemaphoreInfo-parameter)VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-pExternalSemaphoreInfo-parameter  
  `pExternalSemaphoreInfo` **must** be a valid pointer to a valid [VkPhysicalDeviceExternalSemaphoreInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html) structure
- [](#VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-pExternalSemaphoreProperties-parameter)VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-pExternalSemaphoreProperties-parameter  
  `pExternalSemaphoreProperties` **must** be a valid pointer to a [VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreProperties.html) structure

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalSemaphoreProperties.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceExternalSemaphoreInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceExternalSemaphoreProperties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0