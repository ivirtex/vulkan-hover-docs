# vkGetDeviceImageMemoryRequirements(3) Manual Page

## Name

vkGetDeviceImageMemoryRequirements - Returns the memory requirements for
specified Vulkan object



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the memory requirements for an image resource without
creating an object, call:

``` c
// Provided by VK_VERSION_1_3
void vkGetDeviceImageMemoryRequirements(
    VkDevice                                    device,
    const VkDeviceImageMemoryRequirements*      pInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

or the equivalent command

``` c
// Provided by VK_KHR_maintenance4
void vkGetDeviceImageMemoryRequirementsKHR(
    VkDevice                                    device,
    const VkDeviceImageMemoryRequirements*      pInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device intended to own the image.

- `pInfo` is a pointer to a
  [VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageMemoryRequirements.html)
  structure containing parameters required for the memory requirements
  query.

- `pMemoryRequirements` is a pointer to a
  [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html) structure in which
  the memory requirements of the image object are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkGetDeviceImageMemoryRequirements-device-parameter"
  id="VUID-vkGetDeviceImageMemoryRequirements-device-parameter"></a>
  VUID-vkGetDeviceImageMemoryRequirements-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetDeviceImageMemoryRequirements-pInfo-parameter"
  id="VUID-vkGetDeviceImageMemoryRequirements-pInfo-parameter"></a>
  VUID-vkGetDeviceImageMemoryRequirements-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageMemoryRequirements.html)
  structure

- <a
  href="#VUID-vkGetDeviceImageMemoryRequirements-pMemoryRequirements-parameter"
  id="VUID-vkGetDeviceImageMemoryRequirements-pMemoryRequirements-parameter"></a>
  VUID-vkGetDeviceImageMemoryRequirements-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a
  [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance4](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance4.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceImageMemoryRequirements.html),
[VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceImageMemoryRequirements"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
