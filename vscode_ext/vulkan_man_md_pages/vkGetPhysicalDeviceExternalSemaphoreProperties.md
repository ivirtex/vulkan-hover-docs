# vkGetPhysicalDeviceExternalSemaphoreProperties(3) Manual Page

## Name

vkGetPhysicalDeviceExternalSemaphoreProperties - Function for querying
external semaphore handle capabilities.



## <a href="#_c_specification" class="anchor"></a>C Specification

Semaphores **may** support import and export of their <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-payloads"
target="_blank" rel="noopener">payload</a> to external handles. To query
the external handle types supported by semaphores, call:

``` c
// Provided by VK_VERSION_1_1
void vkGetPhysicalDeviceExternalSemaphoreProperties(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceExternalSemaphoreInfo* pExternalSemaphoreInfo,
    VkExternalSemaphoreProperties*              pExternalSemaphoreProperties);
```

or the equivalent command

``` c
// Provided by VK_KHR_external_semaphore_capabilities
void vkGetPhysicalDeviceExternalSemaphorePropertiesKHR(
    VkPhysicalDevice                            physicalDevice,
    const VkPhysicalDeviceExternalSemaphoreInfo* pExternalSemaphoreInfo,
    VkExternalSemaphoreProperties*              pExternalSemaphoreProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the
  semaphore capabilities.

- `pExternalSemaphoreInfo` is a pointer to a
  [VkPhysicalDeviceExternalSemaphoreInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html)
  structure describing the parameters that would be consumed by
  [vkCreateSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSemaphore.html).

- `pExternalSemaphoreProperties` is a pointer to a
  [VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreProperties.html)
  structure in which capabilities are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-pExternalSemaphoreInfo-parameter"
  id="VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-pExternalSemaphoreInfo-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-pExternalSemaphoreInfo-parameter  
  `pExternalSemaphoreInfo` **must** be a valid pointer to a valid
  [VkPhysicalDeviceExternalSemaphoreInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html)
  structure

- <a
  href="#VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-pExternalSemaphoreProperties-parameter"
  id="VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-pExternalSemaphoreProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalSemaphoreProperties-pExternalSemaphoreProperties-parameter  
  `pExternalSemaphoreProperties` **must** be a valid pointer to a
  [VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreProperties.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkExternalSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreProperties.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkPhysicalDeviceExternalSemaphoreInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalSemaphoreInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceExternalSemaphoreProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
