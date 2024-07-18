# vkGetPhysicalDeviceMultisamplePropertiesEXT(3) Manual Page

## Name

vkGetPhysicalDeviceMultisamplePropertiesEXT - Report sample count
specific multisampling capabilities of a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

To query additional multisampling capabilities which **may** be
supported for a specific sample count, beyond the minimum capabilities
described for <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits"
target="_blank" rel="noopener">Limits</a> above, call:

``` c
// Provided by VK_EXT_sample_locations
void vkGetPhysicalDeviceMultisamplePropertiesEXT(
    VkPhysicalDevice                            physicalDevice,
    VkSampleCountFlagBits                       samples,
    VkMultisamplePropertiesEXT*                 pMultisampleProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the
  additional multisampling capabilities.

- `samples` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html)
  value specifying the sample count to query capabilities for.

- `pMultisampleProperties` is a pointer to a
  [VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisamplePropertiesEXT.html)
  structure in which information about additional multisampling
  capabilities specific to the sample count is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-samples-parameter"
  id="VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-samples-parameter"></a>
  VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-samples-parameter  
  `samples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value

- <a
  href="#VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-pMultisampleProperties-parameter"
  id="VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-pMultisampleProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceMultisamplePropertiesEXT-pMultisampleProperties-parameter  
  `pMultisampleProperties` **must** be a valid pointer to a
  [VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisamplePropertiesEXT.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_sample_locations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sample_locations.html),
[VkMultisamplePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultisamplePropertiesEXT.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html),
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceMultisamplePropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
