# vkGetPhysicalDeviceQueueFamilyProperties2(3) Manual Page

## Name

vkGetPhysicalDeviceQueueFamilyProperties2 - Reports properties of the
queues of the specified physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

To query properties of queues available on a physical device, call:

``` c
// Provided by VK_VERSION_1_1
void vkGetPhysicalDeviceQueueFamilyProperties2(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pQueueFamilyPropertyCount,
    VkQueueFamilyProperties2*                   pQueueFamilyProperties);
```

or the equivalent command

``` c
// Provided by VK_KHR_get_physical_device_properties2
void vkGetPhysicalDeviceQueueFamilyProperties2KHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pQueueFamilyPropertyCount,
    VkQueueFamilyProperties2*                   pQueueFamilyProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the handle to the physical device whose properties
  will be queried.

- `pQueueFamilyPropertyCount` is a pointer to an integer related to the
  number of queue families available or queried, as described in
  [vkGetPhysicalDeviceQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html).

- `pQueueFamilyProperties` is either `NULL` or a pointer to an array of
  [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2.html) structures.

## <a href="#_description" class="anchor"></a>Description

`vkGetPhysicalDeviceQueueFamilyProperties2` behaves similarly to
[vkGetPhysicalDeviceQueueFamilyProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties.html),
with the ability to return extended information in a `pNext` chain of
output structures.

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceQueueFamilyProperties2-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceQueueFamilyProperties2-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceQueueFamilyProperties2-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceQueueFamilyProperties2-pQueueFamilyPropertyCount-parameter"
  id="VUID-vkGetPhysicalDeviceQueueFamilyProperties2-pQueueFamilyPropertyCount-parameter"></a>
  VUID-vkGetPhysicalDeviceQueueFamilyProperties2-pQueueFamilyPropertyCount-parameter  
  `pQueueFamilyPropertyCount` **must** be a valid pointer to a
  `uint32_t` value

- <a
  href="#VUID-vkGetPhysicalDeviceQueueFamilyProperties2-pQueueFamilyProperties-parameter"
  id="VUID-vkGetPhysicalDeviceQueueFamilyProperties2-pQueueFamilyProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceQueueFamilyProperties2-pQueueFamilyProperties-parameter  
  If the value referenced by `pQueueFamilyPropertyCount` is not `0`, and
  `pQueueFamilyProperties` is not `NULL`, `pQueueFamilyProperties`
  **must** be a valid pointer to an array of `pQueueFamilyPropertyCount`
  [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceQueueFamilyProperties2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
