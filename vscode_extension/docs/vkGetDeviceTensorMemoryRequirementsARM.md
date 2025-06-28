# vkGetDeviceTensorMemoryRequirementsARM(3) Manual Page

## Name

vkGetDeviceTensorMemoryRequirementsARM - Returns the memory requirements for specified tensor creation infos



## [](#_c_specification)C Specification

To determine the memory requirements for a tensor resource without creating an object, call:

```c++
// Provided by VK_ARM_tensors
void vkGetDeviceTensorMemoryRequirementsARM(
    VkDevice                                    device,
    const VkDeviceTensorMemoryRequirementsARM*  pInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device intended to own the tensor.
- `pInfo` is a pointer to a [VkDeviceTensorMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceTensorMemoryRequirementsARM.html) structure containing parameters required for the memory requirements query.
- `pMemoryRequirements` is a pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure in which the memory requirements of the tensor object are returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetDeviceTensorMemoryRequirementsARM-tensors-09831)VUID-vkGetDeviceTensorMemoryRequirementsARM-tensors-09831  
  The [`tensors`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-tensors) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceTensorMemoryRequirementsARM-device-parameter)VUID-vkGetDeviceTensorMemoryRequirementsARM-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceTensorMemoryRequirementsARM-pInfo-parameter)VUID-vkGetDeviceTensorMemoryRequirementsARM-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkDeviceTensorMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceTensorMemoryRequirementsARM.html) structure
- [](#VUID-vkGetDeviceTensorMemoryRequirementsARM-pMemoryRequirements-parameter)VUID-vkGetDeviceTensorMemoryRequirementsARM-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceTensorMemoryRequirementsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceTensorMemoryRequirementsARM.html), [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceTensorMemoryRequirementsARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0