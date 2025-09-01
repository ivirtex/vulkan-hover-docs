# vkGetDeviceImageMemoryRequirements(3) Manual Page

## Name

vkGetDeviceImageMemoryRequirements - Returns the memory requirements for specified Vulkan object



## [](#_c_specification)C Specification

To determine the memory requirements for an image resource without creating an object, call:

```c++
// Provided by VK_VERSION_1_3
void vkGetDeviceImageMemoryRequirements(
    VkDevice                                    device,
    const VkDeviceImageMemoryRequirements*      pInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

or the equivalent command

```c++
// Provided by VK_KHR_maintenance4
void vkGetDeviceImageMemoryRequirementsKHR(
    VkDevice                                    device,
    const VkDeviceImageMemoryRequirements*      pInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device intended to own the image.
- `pInfo` is a pointer to a [VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageMemoryRequirements.html) structure containing parameters required for the memory requirements query.
- `pMemoryRequirements` is a pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure in which the memory requirements of the image object are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceImageMemoryRequirements-device-parameter)VUID-vkGetDeviceImageMemoryRequirements-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceImageMemoryRequirements-pInfo-parameter)VUID-vkGetDeviceImageMemoryRequirements-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageMemoryRequirements.html) structure
- [](#VUID-vkGetDeviceImageMemoryRequirements-pMemoryRequirements-parameter)VUID-vkGetDeviceImageMemoryRequirements-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure

## [](#_see_also)See Also

[VK\_KHR\_maintenance4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance4.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageMemoryRequirements.html), [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceImageMemoryRequirements).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0