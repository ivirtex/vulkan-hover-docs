# vkGetPhysicalDeviceFeatures(3) Manual Page

## Name

vkGetPhysicalDeviceFeatures - Reports capabilities of a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

To query supported features, call:

``` c
// Provided by VK_VERSION_1_0
void vkGetPhysicalDeviceFeatures(
    VkPhysicalDevice                            physicalDevice,
    VkPhysicalDeviceFeatures*                   pFeatures);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the
  supported features.

- `pFeatures` is a pointer to a
  [VkPhysicalDeviceFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures.html) structure in
  which the physical device features are returned. For each feature, a
  value of `VK_TRUE` specifies that the feature is supported on this
  physical device, and `VK_FALSE` specifies that the feature is not
  supported.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkGetPhysicalDeviceFeatures-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceFeatures-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceFeatures-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkGetPhysicalDeviceFeatures-pFeatures-parameter"
  id="VUID-vkGetPhysicalDeviceFeatures-pFeatures-parameter"></a>
  VUID-vkGetPhysicalDeviceFeatures-pFeatures-parameter  
  `pFeatures` **must** be a valid pointer to a
  [VkPhysicalDeviceFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkPhysicalDeviceFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
