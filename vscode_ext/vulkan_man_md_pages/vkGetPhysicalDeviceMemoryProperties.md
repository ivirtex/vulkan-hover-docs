# vkGetPhysicalDeviceMemoryProperties(3) Manual Page

## Name

vkGetPhysicalDeviceMemoryProperties - Reports memory information for the
specified physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

To query memory properties, call:

``` c
// Provided by VK_VERSION_1_0
void vkGetPhysicalDeviceMemoryProperties(
    VkPhysicalDevice                            physicalDevice,
    VkPhysicalDeviceMemoryProperties*           pMemoryProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the handle to the device to query.

- `pMemoryProperties` is a pointer to a
  [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html)
  structure in which the properties are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceMemoryProperties-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceMemoryProperties-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceMemoryProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceMemoryProperties-pMemoryProperties-parameter"
  id="VUID-vkGetPhysicalDeviceMemoryProperties-pMemoryProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceMemoryProperties-pMemoryProperties-parameter  
  `pMemoryProperties` **must** be a valid pointer to a
  [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceMemoryProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
