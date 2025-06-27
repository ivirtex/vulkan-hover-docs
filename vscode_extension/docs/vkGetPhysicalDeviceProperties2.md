# vkGetPhysicalDeviceProperties2(3) Manual Page

## Name

vkGetPhysicalDeviceProperties2 - Returns properties of a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

To query general properties of physical devices once enumerated, call:

``` c
// Provided by VK_VERSION_1_1
void vkGetPhysicalDeviceProperties2(
    VkPhysicalDevice                            physicalDevice,
    VkPhysicalDeviceProperties2*                pProperties);
```

or the equivalent command

``` c
// Provided by VK_KHR_get_physical_device_properties2
void vkGetPhysicalDeviceProperties2KHR(
    VkPhysicalDevice                            physicalDevice,
    VkPhysicalDeviceProperties2*                pProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the handle to the physical device whose properties
  will be queried.

- `pProperties` is a pointer to a
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
  structure in which properties are returned.

## <a href="#_description" class="anchor"></a>Description

Each structure in `pProperties` and its `pNext` chain contains members
corresponding to implementation-dependent properties, behaviors, or
limits. `vkGetPhysicalDeviceProperties2` fills in each member to specify
the corresponding value for the implementation.

Valid Usage (Implicit)

- <a href="#VUID-vkGetPhysicalDeviceProperties2-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceProperties2-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceProperties2-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkGetPhysicalDeviceProperties2-pProperties-parameter"
  id="VUID-vkGetPhysicalDeviceProperties2-pProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceProperties2-pProperties-parameter  
  `pProperties` **must** be a valid pointer to a
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceProperties2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
