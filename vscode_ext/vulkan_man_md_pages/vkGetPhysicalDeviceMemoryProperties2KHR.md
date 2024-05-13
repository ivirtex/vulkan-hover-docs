# vkGetPhysicalDeviceMemoryProperties2(3) Manual Page

## Name

vkGetPhysicalDeviceMemoryProperties2 - Reports memory information for
the specified physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

To query memory properties, call:

``` c
// Provided by VK_VERSION_1_1
void vkGetPhysicalDeviceMemoryProperties2(
    VkPhysicalDevice                            physicalDevice,
    VkPhysicalDeviceMemoryProperties2*          pMemoryProperties);
```

or the equivalent command

``` c
// Provided by VK_KHR_get_physical_device_properties2
void vkGetPhysicalDeviceMemoryProperties2KHR(
    VkPhysicalDevice                            physicalDevice,
    VkPhysicalDeviceMemoryProperties2*          pMemoryProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the handle to the device to query.

- `pMemoryProperties` is a pointer to a
  [VkPhysicalDeviceMemoryProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties2.html)
  structure in which the properties are returned.

## <a href="#_description" class="anchor"></a>Description

`vkGetPhysicalDeviceMemoryProperties2` behaves similarly to
[vkGetPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceMemoryProperties.html),
with the ability to return extended information in a `pNext` chain of
output structures.

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceMemoryProperties2-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceMemoryProperties2-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceMemoryProperties2-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceMemoryProperties2-pMemoryProperties-parameter"
  id="VUID-vkGetPhysicalDeviceMemoryProperties2-pMemoryProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceMemoryProperties2-pMemoryProperties-parameter  
  `pMemoryProperties` **must** be a valid pointer to a
  [VkPhysicalDeviceMemoryProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties2.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkPhysicalDeviceMemoryProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceMemoryProperties2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
